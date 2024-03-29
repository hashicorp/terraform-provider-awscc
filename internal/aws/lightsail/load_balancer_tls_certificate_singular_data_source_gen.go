// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lightsail_load_balancer_tls_certificate", loadBalancerTlsCertificateDataSource)
}

// loadBalancerTlsCertificateDataSource returns the Terraform awscc_lightsail_load_balancer_tls_certificate data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lightsail::LoadBalancerTlsCertificate resource.
func loadBalancerTlsCertificateDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CertificateAlternativeNames
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of strings listing alternative domains and subdomains for your SSL/TLS certificate.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"certificate_alternative_names": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "An array of strings listing alternative domains and subdomains for your SSL/TLS certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CertificateDomainName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The domain name (e.g., example.com ) for your SSL/TLS certificate.",
		//	  "type": "string"
		//	}
		"certificate_domain_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The domain name (e.g., example.com ) for your SSL/TLS certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CertificateName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The SSL/TLS certificate name.",
		//	  "type": "string"
		//	}
		"certificate_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The SSL/TLS certificate name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HttpsRedirectionEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A Boolean value that indicates whether HTTPS redirection is enabled for the load balancer.",
		//	  "type": "boolean"
		//	}
		"https_redirection_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A Boolean value that indicates whether HTTPS redirection is enabled for the load balancer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IsAttached
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "When true, the SSL/TLS certificate is attached to the Lightsail load balancer.",
		//	  "type": "boolean"
		//	}
		"is_attached": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "When true, the SSL/TLS certificate is attached to the Lightsail load balancer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of your load balancer.",
		//	  "pattern": "\\w[\\w\\-]*\\w",
		//	  "type": "string"
		//	}
		"load_balancer_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of your load balancer.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LoadBalancerTlsCertificateArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"load_balancer_tls_certificate_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The validation status of the SSL/TLS certificate.",
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The validation status of the SSL/TLS certificate.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Lightsail::LoadBalancerTlsCertificate",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::LoadBalancerTlsCertificate").WithTerraformTypeName("awscc_lightsail_load_balancer_tls_certificate")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate_alternative_names":     "CertificateAlternativeNames",
		"certificate_domain_name":           "CertificateDomainName",
		"certificate_name":                  "CertificateName",
		"https_redirection_enabled":         "HttpsRedirectionEnabled",
		"is_attached":                       "IsAttached",
		"load_balancer_name":                "LoadBalancerName",
		"load_balancer_tls_certificate_arn": "LoadBalancerTlsCertificateArn",
		"status":                            "Status",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
