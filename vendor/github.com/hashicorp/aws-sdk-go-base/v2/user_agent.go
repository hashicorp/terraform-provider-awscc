package awsbase

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func customUserAgentMiddleware(c *Config) middleware.BuildMiddleware {
	return middleware.BuildMiddlewareFunc("CustomUserAgent",
		func(ctx context.Context, in middleware.BuildInput, next middleware.BuildHandler) (middleware.BuildOutput, middleware.Metadata, error) {
			request, ok := in.Request.(*smithyhttp.Request)
			if !ok {
				return middleware.BuildOutput{}, middleware.Metadata{}, fmt.Errorf("unknown request type %T", in.Request)
			}

			builder := smithyhttp.NewUserAgentBuilder()
			for _, v := range c.UserAgentProducts {
				builder.AddKey(userAgentProduct(v))
			}
			prependUserAgentHeader(request, builder.Build())

			return next.HandleBuild(ctx, in)
		})
}

func userAgentProduct(product *UserAgentProduct) string {
	var b strings.Builder
	fmt.Fprint(&b, product.Name)
	if product.Version != "" {
		fmt.Fprintf(&b, "/%s", product.Version)
	}
	if len(product.Extra) > 0 {
		fmt.Fprintf(&b, " (%s)", strings.Join(product.Extra, "; "))
	}

	return b.String()
}

// Because the default User-Agent middleware prepends itself to the contents of the User-Agent header,
// we have to run after it and also prepend our custom User-Agent
func prependUserAgentHeader(request *smithyhttp.Request, value string) {
	current := request.Header.Get("User-Agent")
	if len(current) > 0 {
		current = value + " " + current
	} else {
		current = value
	}
	request.Header["User-Agent"] = append(request.Header["User-Agent"][:0], current)

}
