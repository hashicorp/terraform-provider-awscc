package generic

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

func DataSourceNotFoundDiag(err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"AWS Data Source Not Found",
		fmt.Sprintf("After attempting to read the data source, the API returned a resource not found error for the id provided. Original Error: %s", err.Error()),
	)
}

func DesiredStateErrorDiag(source string, err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"Creation Of Cloud Control API Desired State Unsuccessful",
		fmt.Sprintf("Unable to create Cloud Control API Desired State from Terraform %s. This is typically an error with the Terraform provider implementation. Original Error: %s", source, err.Error()),
	)
}

func ResourceAttributeNotSetInImportStateDiag(err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"Terraform Resource Attribute Not Set in Import State",
		fmt.Sprintf("Terraform resource attribute not set in Import State. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
	)
}

func ResourceIdentifierNotFoundDiag(err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"Terraform Resource Identifier Not Found",
		fmt.Sprintf("Terraform resource primary identifier not found in State. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
	)
}

func ResourceIdentifierNotSetDiag(err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"Terraform Resource Identifier Not Set",
		fmt.Sprintf("Terraform resource primary identifier not set in State. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
	)
}

func ResourceNotFoundAfterWriteDiag(err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"AWS Resource Not Found After Creation or Update",
		fmt.Sprintf("After creating or updating the AWS resource and attempting to read the resource, the API returned a resource not found error. This is typically an error with the Terraform resource implementation. Original Error: %s", err.Error()),
	)
}

func ResourceNotFoundWarningDiag(err error) diag.Diagnostic {
	return diag.NewWarningDiagnostic(
		"AWS Resource Not Found During Refresh",
		fmt.Sprintf("Automatically removing from Terraform State instead of returning the error, which may trigger resource recreation. Original Error: %s", err.Error()),
	)
}

func ServiceOperationEmptyResultDiag(service string, operation string) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"AWS SDK Go Service Operation Empty Response",
		fmt.Sprintf("Calling %s service %s operation returned missing contents in the response. This is typically an error with the API implementation.", service, operation),
	)
}

func ServiceOperationErrorDiag(service string, operation string, err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"AWS SDK Go Service Operation Unsuccessful",
		fmt.Sprintf("Calling %s service %s operation returned: %s", service, operation, err.Error()),
	)
}

func ServiceOperationWaiterErrorDiag(service string, operation string, err error) diag.Diagnostic {
	return diag.NewErrorDiagnostic(
		"AWS SDK Go Service Operation Incomplete",
		fmt.Sprintf("Waiting for %s service %s operation completion returned: %s", service, operation, err.Error()),
	)
}
