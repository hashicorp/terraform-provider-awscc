package generic

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func ClientNotFoundDiag(err error) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "AWS CloudFormation Client Not Found",
		Detail:   fmt.Sprintf("AWS CloudFormation client not available. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
	}
}

func DesiredStateErrorDiag(source string, err error) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "Creation Of CloudFormation Desired State Unsuccessful",
		Detail:   fmt.Sprintf("Unable to create CloudFormation Desired State from Terraform %s. This is typically an error with the Terraform provider implementation. Original Error: %s", source, err.Error()),
	}
}

func ResourceIdentifierNotFoundDiag(err error) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "Terraform Resource Identifier Not Found",
		Detail:   fmt.Sprintf("Terraform resource primary identifier not found in State. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
	}
}

func ResourceNotFoundAfterCreationDiag(err error) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "AWS Resource Not Found After Creation",
		Detail:   fmt.Sprintf("After creating the AWS resource and attempting to read the resource, the API returned a resource not found error. This is typically an error with the Terraform resource implementation. Original Error: %s", err.Error()),
	}
}

func ResourceNotFoundWarningDiag(err error) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityWarning,
		Summary:  "AWS Resource Not Found During Refresh",
		Detail:   fmt.Sprintf("Automatically removing from Terraform State instead of returning the error, which may trigger resource recreation. Original Error: %s", err.Error()),
	}
}

func ServiceOperationEmptyResultDiag(service string, operation string) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "AWS SDK Go Service Operation Empty Response",
		Detail:   fmt.Sprintf("Calling %s service %s operation returned missing contents in the response. This is typically an error with the API implementation.", service, operation),
	}
}

func ServiceOperationErrorDiag(service string, operation string, err error) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "AWS SDK Go Service Operation Unsuccessful",
		Detail:   fmt.Sprintf("Calling %s service %s operation returned: %s", service, operation, err.Error()),
	}
}

func ServiceOperationWaiterErrorDiag(service string, operation string, err error) *tfprotov6.Diagnostic {
	return &tfprotov6.Diagnostic{
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "AWS SDK Go Service Operation Incomplete",
		Detail:   fmt.Sprintf("Waiting for %s service %s operation completion returned: %s", service, operation, err.Error()),
	}
}
