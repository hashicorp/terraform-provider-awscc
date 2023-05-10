resource "awscc_opensearchserverless_security_config" "config" {
  name        = "awscc-security-config"
  description = "OpenSearch Serverless security policy template"
  type        = "saml"
  saml_options = {
    group_attribute = "ALLGroups"
    session_timeout = "120"
    user_attribute  = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier"
    metadata        = <<XML
        <md:EntityDescriptor entityID="http://www.okta.com/foobar" xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata">
          <md:IDPSSODescriptor WantAuthnRequestsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
          <md:KeyDescriptor use="signing">
            <ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
              <ds:X509Data>
                <ds:X509Certificate>Mfoobar</ds:X509Certificate>
              </ds:X509Data>
            </ds:KeyInfo>
          </md:KeyDescriptor>
            <md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress</md:NameIDFormat>
            <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://trial-1234567.okta.com/app/trial-1234567_saml2_1/foobar/sso/saml"/>
            <md:SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location="https://trial-1234567.okta.com/app/trial-1234567_saml2_1/foobar/sso/saml"/>
          </md:IDPSSODescriptor>
        </md:EntityDescriptor>
    XML
  }
}