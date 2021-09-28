package awsbase

import (
	"github.com/hashicorp/aws-sdk-go-base/v2/internal/config"
)

// Config, APNInfo, APNProduct, and AssumeRole are aliased to an internal package to break a dependency cycle
// in internal/httpclient.

type Config = config.Config

type APNInfo = config.APNInfo

type APNProduct = config.APNProduct

type AssumeRole = config.AssumeRole
