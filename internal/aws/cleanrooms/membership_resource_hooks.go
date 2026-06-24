// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanrooms

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	tfcloudcontrol "github.com/hashicorp/terraform-provider-awscc/internal/service/cloudcontrol"
)

const (
	analysisTemplateCfTypeName = "AWS::CleanRooms::AnalysisTemplate"
	templateDeleteTimeout      = 10 * time.Minute
)

// preDeleteMembership deletes all analysis templates on a membership before
// the membership itself is deleted via Cloud Control API. Without this, the
// DeleteMembership call returns a 409 Conflict if any analysis templates exist.
//
// This uses only the Cloud Control API client — no native service SDK needed.
// Analysis template identifiers have the format "MembershipId|AnalysisTemplateId",
// so we list all analysis templates and filter by the membership ID prefix.
func preDeleteMembership(ctx context.Context, provider tfcloudcontrol.Provider, membershipID string) error {
	conn := provider.CloudControlAPIClient(ctx)
	roleARN := provider.RoleARN(ctx)

	resources, err := tfcloudcontrol.ListResourcesByTypeName(ctx, conn, roleARN, analysisTemplateCfTypeName)
	if err != nil {
		// If no analysis templates exist at all, that's fine — nothing to clean up.
		if isNotFoundError(err) {
			return nil
		}
		return fmt.Errorf("listing analysis templates: %w", err)
	}

	for _, resource := range resources {
		if resource.Identifier == nil {
			continue
		}

		identifier := *resource.Identifier

		// Analysis template identifiers are "MembershipId|AnalysisTemplateId".
		// Only delete templates belonging to this membership.
		if !strings.HasPrefix(identifier, membershipID+"|") {
			continue
		}

		tflog.Info(ctx, "Deleting analysis template before membership deletion", map[string]any{
			"membership_id":              membershipID,
			"analysis_template_identifier": identifier,
		})

		err := tfcloudcontrol.DeleteResource(
			ctx, conn, roleARN,
			analysisTemplateCfTypeName, identifier,
			templateDeleteTimeout,
		)
		if err != nil {
			return fmt.Errorf("deleting analysis template %s: %w", identifier, err)
		}
	}

	return nil
}

// isNotFoundError checks if the error is a "not found" error from the
// Cloud Control list operation (empty results).
func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "Empty result")
}
