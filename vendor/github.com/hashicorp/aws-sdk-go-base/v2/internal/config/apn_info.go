package config

import (
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Builds the user-agent string for APN
func (apn APNInfo) BuildUserAgentString() string {
	builder := smithyhttp.NewUserAgentBuilder()
	builder.AddKeyValue("APN", "1.0")
	builder.AddKeyValue(apn.PartnerName, "1.0")
	for _, p := range apn.Products {
		p.buildUserAgentPart(builder)
	}
	return builder.Build()
}

func (p APNProduct) buildUserAgentPart(b *smithyhttp.UserAgentBuilder) {
	if p.Name != "" {
		if p.Version != "" {
			b.AddKeyValue(p.Name, p.Version)
		} else {
			b.AddKey(p.Name)
		}
	}
	if p.Comment != "" {
		b.AddKey("(" + p.Comment + ")")
	}
}
