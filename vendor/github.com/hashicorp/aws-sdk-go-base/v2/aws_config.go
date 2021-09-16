package awsbase

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/ec2/imds"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/smithy-go/middleware"
	"github.com/hashicorp/aws-sdk-go-base/v2/internal/constants"
	"github.com/hashicorp/aws-sdk-go-base/v2/internal/endpoints"
	"github.com/hashicorp/go-cleanhttp"
)

func GetAwsConfig(ctx context.Context, c *Config) (aws.Config, error) {
	credentialsProvider, err := getCredentialsProvider(ctx, c)
	if err != nil {
		return aws.Config{}, err
	}

	var retryer aws.Retryer
	retryer = retry.NewStandard()
	if c.MaxRetries != 0 {
		retryer = retry.AddWithMaxAttempts(retryer, c.MaxRetries)
	}
	retryer = &networkErrorShortcutter{
		Retryer: retryer,
	}

	loadOptions := append(
		commonLoadOptions(c),
		config.WithCredentialsProvider(credentialsProvider),
		config.WithRetryer(func() aws.Retryer {
			return retryer
		}),
	)
	cfg, err := config.LoadDefaultConfig(ctx, loadOptions...)
	if err != nil {
		return cfg, fmt.Errorf("loading configuration: %w", err)
	}

	if !c.SkipCredsValidation {
		if _, _, err := getAccountIDAndPartitionFromSTSGetCallerIdentity(ctx, sts.NewFromConfig(cfg)); err != nil {
			return cfg, fmt.Errorf("error validating provider credentials: %w", err)
		}
	}

	return cfg, nil
}

// networkErrorShortcutter is used to enable networking error shortcutting
type networkErrorShortcutter struct {
	aws.Retryer
}

// We're misusing RetryDelay here, since this is the only function that takes the attempt count
func (r *networkErrorShortcutter) RetryDelay(attempt int, err error) (time.Duration, error) {
	if attempt >= constants.MaxNetworkRetryCount {
		var netOpErr *net.OpError
		if errors.As(err, &netOpErr) {
			// It's disappointing that we have to do string matching here, rather than being able to using `errors.Is()` or even strings exported by the Go `net` package
			if strings.Contains(netOpErr.Error(), "no such host") || strings.Contains(netOpErr.Error(), "connection refused") {
				log.Printf("[WARN] Disabling retries after next request due to networking issue: %s", err)
				return 0, &retry.MaxAttemptsError{
					Attempt: attempt,
					Err:     err,
				}
			}
		}
	}

	return r.Retryer.RetryDelay(attempt, err)
}

func GetAwsAccountIDAndPartition(ctx context.Context, awsConfig aws.Config, skipCredsValidation, skipRequestingAccountId bool) (string, string, error) {
	if !skipCredsValidation {
		stsClient := sts.NewFromConfig(awsConfig)
		accountID, partition, err := getAccountIDAndPartitionFromSTSGetCallerIdentity(ctx, stsClient)
		if err != nil {
			return "", "", fmt.Errorf("error validating provider credentials: %w", err)
		}

		return accountID, partition, nil
	}

	if !skipRequestingAccountId {
		credentialsProviderName := ""
		if credentialsValue, err := awsConfig.Credentials.Retrieve(context.Background()); err == nil {
			credentialsProviderName = credentialsValue.Source
		}

		iamClient := iam.NewFromConfig(awsConfig)
		stsClient := sts.NewFromConfig(awsConfig)
		accountID, partition, err := getAccountIDAndPartition(ctx, iamClient, stsClient, credentialsProviderName)

		if err == nil {
			return accountID, partition, nil
		}

		return "", "", fmt.Errorf(
			"AWS account ID not previously found and failed retrieving via all available methods. "+
				"See https://www.terraform.io/docs/providers/aws/index.html#skip_requesting_account_id for workaround and implications. "+
				"Errors: %w", err)
	}

	return "", endpoints.PartitionForRegion(awsConfig.Region), nil
}

func commonLoadOptions(c *Config) []func(*config.LoadOptions) error {
	httpClient := cleanhttp.DefaultClient()
	if c.Insecure {
		transport := httpClient.Transport.(*http.Transport)
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	apiOptions := make([]func(*middleware.Stack) error, 0)
	if len(c.UserAgentProducts) > 0 {
		apiOptions = append(apiOptions, func(stack *middleware.Stack) error {
			// Because the default User-Agent middleware prepends itself to the contents of the User-Agent header,
			// we have to run after it and also prepend our custom User-Agent
			return stack.Build.Add(customUserAgentMiddleware(c), middleware.After)
		})
	}
	if v := os.Getenv(constants.AppendUserAgentEnvVar); v != "" {
		log.Printf("[DEBUG] Using additional User-Agent Info: %s", v)
		apiOptions = append(apiOptions, awsmiddleware.AddUserAgentKey(v))
	}

	loadOptions := []func(*config.LoadOptions) error{
		config.WithRegion(c.Region),
		config.WithEndpointResolver(endpointResolver(c)),
		config.WithHTTPClient(httpClient),
		config.WithAPIOptions(apiOptions),
	}

	if len(c.SharedConfigFiles) > 0 {
		loadOptions = append(
			loadOptions,
			config.WithSharedConfigFiles(c.SharedConfigFiles),
		)
	}

	if c.DebugLogging {
		loadOptions = append(loadOptions,
			config.WithClientLogMode(aws.LogRequestWithBody|aws.LogResponseWithBody|aws.LogRetries),
			config.WithLogger(debugLogger{}),
		)
	}

	if c.SkipMetadataApiCheck {
		loadOptions = append(loadOptions,
			config.WithEC2IMDSClientEnableState(imds.ClientDisabled),
		)

		// This should not be needed, but https://github.com/aws/aws-sdk-go-v2/issues/1398
		os.Setenv("AWS_EC2_METADATA_DISABLED", "true")
	}

	return loadOptions
}
