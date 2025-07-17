
defaults {
  schema_cache_directory     = "../service/cloudformation/schemas"
  terraform_type_name_prefix = "awscc"
}

meta_schema {
  path = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}

resource_schema "aws_accessanalyzer_analyzer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AccessAnalyzer::Analyzer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_acmpca_certificate" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ACMPCA::Certificate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_acmpca_certificate_authority" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ACMPCA::CertificateAuthority"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_acmpca_certificate_authority_activation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ACMPCA::CertificateAuthorityActivation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_acmpca_permission" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ACMPCA::Permission"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_aiops_investigation_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AIOps::InvestigationGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_amazonmq_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AmazonMQ::Configuration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_amplify_app" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Amplify::App"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_amplify_branch" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Amplify::Branch"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_amplify_domain" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Amplify::Domain"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_amplifyuibuilder_component" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AmplifyUIBuilder::Component"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_amplifyuibuilder_form" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AmplifyUIBuilder::Form"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_amplifyuibuilder_theme" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AmplifyUIBuilder::Theme"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_apigateway_account" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::Account"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_api_key" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::ApiKey"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_authorizer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::Authorizer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_base_path_mapping" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::BasePathMapping"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_base_path_mapping_v2" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::BasePathMappingV2"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_client_certificate" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::ClientCertificate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_deployment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::Deployment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_documentation_part" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::DocumentationPart"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_documentation_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::DocumentationVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_domain_name" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::DomainName"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_domain_name_access_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::DomainNameAccessAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_domain_name_v2" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::DomainNameV2"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_gateway_response" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::GatewayResponse"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_method" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::Method"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_model" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::Model"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_request_validator" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::RequestValidator"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_resource" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::Resource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_rest_api" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::RestApi"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_stage" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::Stage"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_usage_plan" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::UsagePlan"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_usage_plan_key" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::UsagePlanKey"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigateway_vpc_link" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGateway::VpcLink"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigatewayv2_api" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::Api"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigatewayv2_api_mapping" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::ApiMapping"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigatewayv2_authorizer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::Authorizer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigatewayv2_deployment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::Deployment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigatewayv2_domain_name" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::DomainName"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigatewayv2_integration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::Integration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_apigatewayv2_integration_response" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::IntegrationResponse"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigatewayv2_model" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::Model"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigatewayv2_route" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::Route"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigatewayv2_route_response" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::RouteResponse"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigatewayv2_routing_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::RoutingRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apigatewayv2_vpc_link" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApiGatewayV2::VpcLink"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appconfig_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppConfig::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appconfig_configuration_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppConfig::ConfigurationProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appconfig_deployment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppConfig::Deployment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appconfig_deployment_strategy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppConfig::DeploymentStrategy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appconfig_environment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppConfig::Environment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appconfig_extension" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppConfig::Extension"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_appconfig_extension_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppConfig::ExtensionAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appconfig_hosted_configuration_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppConfig::HostedConfigurationVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appflow_connector" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppFlow::Connector"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appflow_connector_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppFlow::ConnectorProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appflow_flow" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppFlow::Flow"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appintegrations_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppIntegrations::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appintegrations_data_integration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppIntegrations::DataIntegration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_appintegrations_event_integration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppIntegrations::EventIntegration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_applicationautoscaling_scalable_target" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApplicationAutoScaling::ScalableTarget"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_applicationautoscaling_scaling_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApplicationAutoScaling::ScalingPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_applicationinsights_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApplicationInsights::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_applicationsignals_discovery" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApplicationSignals::Discovery"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_applicationsignals_service_level_objective" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ApplicationSignals::ServiceLevelObjective"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apprunner_auto_scaling_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppRunner::AutoScalingConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apprunner_observability_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppRunner::ObservabilityConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apprunner_service" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppRunner::Service"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apprunner_vpc_connector" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppRunner::VpcConnector"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apprunner_vpc_ingress_connection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppRunner::VpcIngressConnection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appstream_app_block" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppStream::AppBlock"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appstream_app_block_builder" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppStream::AppBlockBuilder"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appstream_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppStream::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appstream_application_entitlement_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppStream::ApplicationEntitlementAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appstream_application_fleet_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppStream::ApplicationFleetAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appstream_directory_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppStream::DirectoryConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appstream_entitlement" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppStream::Entitlement"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appstream_image_builder" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppStream::ImageBuilder"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appsync_api" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppSync::Api"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appsync_channel_namespace" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppSync::ChannelNamespace"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appsync_data_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppSync::DataSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appsync_domain_name" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppSync::DomainName"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appsync_domain_name_api_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppSync::DomainNameApiAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appsync_function_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppSync::FunctionConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appsync_graph_ql_api" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppSync::GraphQLApi"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appsync_resolver" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppSync::Resolver"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_appsync_source_api_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppSync::SourceApiAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_apptest_test_case" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AppTest::TestCase"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_aps_rule_groups_namespace" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::APS::RuleGroupsNamespace"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_aps_scraper" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::APS::Scraper"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_aps_workspace" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::APS::Workspace"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_arczonalshift_autoshift_observer_notification_status" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ARCZonalShift::AutoshiftObserverNotificationStatus"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_arczonalshift_zonal_autoshift_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ARCZonalShift::ZonalAutoshiftConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_athena_capacity_reservation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Athena::CapacityReservation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_athena_data_catalog" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Athena::DataCatalog"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_athena_named_query" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Athena::NamedQuery"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_athena_prepared_statement" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Athena::PreparedStatement"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_athena_work_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Athena::WorkGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_auditmanager_assessment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AuditManager::Assessment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_autoscaling_auto_scaling_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AutoScaling::AutoScalingGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_autoscaling_launch_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AutoScaling::LaunchConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_autoscaling_lifecycle_hook" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AutoScaling::LifecycleHook"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_autoscaling_scaling_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AutoScaling::ScalingPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_autoscaling_scheduled_action" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AutoScaling::ScheduledAction"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_autoscaling_warm_pool" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::AutoScaling::WarmPool"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_b2bi_capability" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::B2BI::Capability"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_b2bi_partnership" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::B2BI::Partnership"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_b2bi_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::B2BI::Profile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_b2bi_transformer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::B2BI::Transformer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_backup_backup_plan" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Backup::BackupPlan"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_backup_backup_selection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Backup::BackupSelection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_backup_backup_vault" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Backup::BackupVault"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_backup_framework" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Backup::Framework"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_backup_logically_air_gapped_backup_vault" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Backup::LogicallyAirGappedBackupVault"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_backup_report_plan" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Backup::ReportPlan"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_backup_restore_testing_plan" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Backup::RestoreTestingPlan"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_backup_restore_testing_selection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Backup::RestoreTestingSelection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_backupgateway_hypervisor" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::BackupGateway::Hypervisor"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_batch_compute_environment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Batch::ComputeEnvironment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_batch_consumable_resource" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Batch::ConsumableResource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_batch_job_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Batch::JobDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_batch_job_queue" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Batch::JobQueue"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_batch_scheduling_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Batch::SchedulingPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bcmdataexports_export" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::BCMDataExports::Export"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_bedrock_agent" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::Agent"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_agent_alias" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::AgentAlias"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_application_inference_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::ApplicationInferenceProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_blueprint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::Blueprint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_data_automation_project" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::DataAutomationProject"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_data_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::DataSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_flow" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::Flow"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_bedrock_flow_alias" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::FlowAlias"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_flow_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::FlowVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_guardrail" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::Guardrail"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_guardrail_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::GuardrailVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_intelligent_prompt_router" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::IntelligentPromptRouter"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_knowledge_base" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::KnowledgeBase"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_prompt" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::Prompt"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_bedrock_prompt_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Bedrock::PromptVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_billing_billing_view" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Billing::BillingView"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_billingconductor_billing_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::BillingConductor::BillingGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_billingconductor_custom_line_item" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::BillingConductor::CustomLineItem"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_billingconductor_pricing_plan" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::BillingConductor::PricingPlan"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_billingconductor_pricing_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::BillingConductor::PricingRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_budgets_budgets_action" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Budgets::BudgetsAction"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cassandra_keyspace" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cassandra::Keyspace"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cassandra_table" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cassandra::Table"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cassandra_type" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cassandra::Type"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ce_anomaly_monitor" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CE::AnomalyMonitor"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ce_anomaly_subscription" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CE::AnomalySubscription"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ce_cost_category" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CE::CostCategory"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_certificatemanager_account" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CertificateManager::Account"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_chatbot_custom_action" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Chatbot::CustomAction"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_chatbot_microsoft_teams_channel_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Chatbot::MicrosoftTeamsChannelConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_chatbot_slack_channel_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Chatbot::SlackChannelConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cleanrooms_analysis_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CleanRooms::AnalysisTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cleanrooms_collaboration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CleanRooms::Collaboration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cleanrooms_configured_table" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CleanRooms::ConfiguredTable"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cleanrooms_configured_table_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CleanRooms::ConfiguredTableAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cleanrooms_id_mapping_table" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CleanRooms::IdMappingTable"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cleanrooms_id_namespace_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CleanRooms::IdNamespaceAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_cleanrooms_membership" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CleanRooms::Membership"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cleanrooms_privacy_budget_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CleanRooms::PrivacyBudgetTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cleanroomsml_training_dataset" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CleanRoomsML::TrainingDataset"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_guard_hook" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::GuardHook"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_hook_default_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::HookDefaultVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_hook_type_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::HookTypeConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_hook_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::HookVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_lambda_hook" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::LambdaHook"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_module_default_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::ModuleDefaultVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_module_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::ModuleVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_public_type_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::PublicTypeVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_publisher" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::Publisher"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_resource_default_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::ResourceDefaultVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_resource_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::ResourceVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_stack" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::Stack"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_stack_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::StackSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudformation_type_activation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFormation::TypeActivation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_anycast_ip_list" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::AnycastIpList"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_cache_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::CachePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_cloudfront_origin_access_identity" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::CloudFrontOriginAccessIdentity"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_connection_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::ConnectionGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_continuous_deployment_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::ContinuousDeploymentPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_distribution" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::Distribution"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_distribution_tenant" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::DistributionTenant"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_function" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::Function"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_key_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::KeyGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_key_value_store" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::KeyValueStore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_monitoring_subscription" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::MonitoringSubscription"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_origin_access_control" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::OriginAccessControl"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_origin_request_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::OriginRequestPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_public_key" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::PublicKey"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_realtime_log_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::RealtimeLogConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_response_headers_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::ResponseHeadersPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudfront_vpc_origin" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudFront::VpcOrigin"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudtrail_channel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudTrail::Channel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudtrail_dashboard" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudTrail::Dashboard"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudtrail_event_data_store" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudTrail::EventDataStore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudtrail_resource_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudTrail::ResourcePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudtrail_trail" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudTrail::Trail"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudwatch_alarm" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudWatch::Alarm"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudwatch_composite_alarm" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudWatch::CompositeAlarm"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudwatch_dashboard" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudWatch::Dashboard"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cloudwatch_metric_stream" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CloudWatch::MetricStream"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codeartifact_domain" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeArtifact::Domain"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codeartifact_package_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeArtifact::PackageGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codeartifact_repository" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeArtifact::Repository"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codebuild_fleet" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeBuild::Fleet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codeconnections_connection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeConnections::Connection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codedeploy_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeDeploy::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codedeploy_deployment_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeDeploy::DeploymentConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codeguruprofiler_profiling_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeGuruProfiler::ProfilingGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codegurureviewer_repository_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeGuruReviewer::RepositoryAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codepipeline_custom_action_type" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodePipeline::CustomActionType"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codepipeline_pipeline" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodePipeline::Pipeline"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codestarconnections_connection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeStarConnections::Connection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codestarconnections_repository_link" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeStarConnections::RepositoryLink"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codestarconnections_sync_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeStarConnections::SyncConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_codestarnotifications_notification_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CodeStarNotifications::NotificationRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_identity_pool" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::IdentityPool"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_identity_pool_principal_tag" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::IdentityPoolPrincipalTag"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_identity_pool_role_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::IdentityPoolRoleAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_log_delivery_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::LogDeliveryConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_managed_login_branding" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::ManagedLoginBranding"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_user_pool" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::UserPool"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_cognito_user_pool_client" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::UserPoolClient"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_user_pool_domain" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::UserPoolDomain"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_user_pool_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::UserPoolGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_user_pool_identity_provider" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::UserPoolIdentityProvider"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_user_pool_resource_server" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::UserPoolResourceServer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_user_pool_risk_configuration_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::UserPoolRiskConfigurationAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_user_pool_ui_customization_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::UserPoolUICustomizationAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_user_pool_user" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::UserPoolUser"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cognito_user_pool_user_to_group_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Cognito::UserPoolUserToGroupAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_comprehend_document_classifier" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Comprehend::DocumentClassifier"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_comprehend_flywheel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Comprehend::Flywheel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_config_aggregation_authorization" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Config::AggregationAuthorization"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_config_config_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Config::ConfigRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_config_configuration_aggregator" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Config::ConfigurationAggregator"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_config_conformance_pack" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Config::ConformancePack"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_config_organization_conformance_pack" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Config::OrganizationConformancePack"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_config_stored_query" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Config::StoredQuery"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_agent_status" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::AgentStatus"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_approved_origin" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::ApprovedOrigin"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_contact_flow" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::ContactFlow"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_contact_flow_module" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::ContactFlowModule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_contact_flow_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::ContactFlowVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_email_address" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::EmailAddress"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_evaluation_form" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::EvaluationForm"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_connect_hours_of_operation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::HoursOfOperation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::Instance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_instance_storage_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::InstanceStorageConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_integration_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::IntegrationAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_phone_number" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::PhoneNumber"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_predefined_attribute" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::PredefinedAttribute"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_prompt" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::Prompt"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_queue" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::Queue"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_quick_connect" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::QuickConnect"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_routing_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::RoutingProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::Rule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_connect_security_key" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::SecurityKey"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_security_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::SecurityProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_task_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::TaskTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_traffic_distribution_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::TrafficDistributionGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_user" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::User"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_user_hierarchy_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::UserHierarchyGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_user_hierarchy_structure" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::UserHierarchyStructure"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_view" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::View"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connect_view_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Connect::ViewVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connectcampaigns_campaign" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ConnectCampaigns::Campaign"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_connectcampaignsv2_campaign" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ConnectCampaignsV2::Campaign"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_controltower_enabled_baseline" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ControlTower::EnabledBaseline"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_controltower_enabled_control" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ControlTower::EnabledControl"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_controltower_landing_zone" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ControlTower::LandingZone"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_cur_report_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CUR::ReportDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_customerprofiles_calculated_attribute_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CustomerProfiles::CalculatedAttributeDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_customerprofiles_domain" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CustomerProfiles::Domain"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_customerprofiles_event_stream" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CustomerProfiles::EventStream"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_customerprofiles_event_trigger" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CustomerProfiles::EventTrigger"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_customerprofiles_integration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CustomerProfiles::Integration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_customerprofiles_object_type" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CustomerProfiles::ObjectType"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_customerprofiles_segment_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::CustomerProfiles::SegmentDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_databrew_dataset" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataBrew::Dataset"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_databrew_job" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataBrew::Job"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_databrew_project" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataBrew::Project"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_databrew_recipe" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataBrew::Recipe"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_databrew_ruleset" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataBrew::Ruleset"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_databrew_schedule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataBrew::Schedule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datapipeline_pipeline" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataPipeline::Pipeline"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_agent" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::Agent"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_location_azure_blob" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::LocationAzureBlob"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_location_efs" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::LocationEFS"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_location_fsx_lustre" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::LocationFSxLustre"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_location_fsx_ontap" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::LocationFSxONTAP"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_location_fsx_open_zfs" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::LocationFSxOpenZFS"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_location_fsx_windows" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::LocationFSxWindows"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_location_hdfs" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::LocationHDFS"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_location_nfs" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::LocationNFS"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_location_object_storage" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::LocationObjectStorage"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_location_s3" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::LocationS3"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_location_smb" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::LocationSMB"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_storage_system" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::StorageSystem"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datasync_task" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataSync::Task"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_connection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::Connection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_data_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::DataSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_domain" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::Domain"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_domain_unit" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::DomainUnit"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_environment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::Environment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_environment_actions" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::EnvironmentActions"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_environment_blueprint_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::EnvironmentBlueprintConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_environment_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::EnvironmentProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_group_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::GroupProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_owner" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::Owner"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_project" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::Project"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_project_membership" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::ProjectMembership"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_project_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::ProjectProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_subscription_target" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::SubscriptionTarget"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_datazone_user_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DataZone::UserProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_deadline_farm" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Deadline::Farm"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_deadline_fleet" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Deadline::Fleet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_deadline_license_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Deadline::LicenseEndpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_deadline_limit" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Deadline::Limit"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_deadline_metered_product" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Deadline::MeteredProduct"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_deadline_monitor" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Deadline::Monitor"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_deadline_queue" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Deadline::Queue"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_deadline_queue_environment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Deadline::QueueEnvironment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_deadline_queue_fleet_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Deadline::QueueFleetAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_deadline_queue_limit_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Deadline::QueueLimitAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_deadline_storage_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Deadline::StorageProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_detective_graph" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Detective::Graph"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_detective_member_invitation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Detective::MemberInvitation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_detective_organization_admin" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Detective::OrganizationAdmin"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_devopsguru_log_anomaly_detection_integration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DevOpsGuru::LogAnomalyDetectionIntegration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_devopsguru_notification_channel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DevOpsGuru::NotificationChannel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_devopsguru_resource_collection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DevOpsGuru::ResourceCollection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_directoryservice_simple_ad" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DirectoryService::SimpleAD"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_dms_data_migration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DMS::DataMigration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_dms_data_provider" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DMS::DataProvider"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_dms_instance_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DMS::InstanceProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_dms_migration_project" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DMS::MigrationProject"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_dms_replication_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DMS::ReplicationConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_docdbelastic_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DocDBElastic::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_dsql_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DSQL::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_dynamodb_global_table" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DynamoDB::GlobalTable"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_dynamodb_table" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::DynamoDB::Table"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_capacity_reservation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::CapacityReservation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_capacity_reservation_fleet" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::CapacityReservationFleet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_carrier_gateway" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::CarrierGateway"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_customer_gateway" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::CustomerGateway"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_dhcp_options" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::DHCPOptions"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_ec2_fleet" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::EC2Fleet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_egress_only_internet_gateway" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::EgressOnlyInternetGateway"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_eip" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::EIP"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_eip_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::EIPAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_enclave_certificate_iam_role_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::EnclaveCertificateIamRoleAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_flow_log" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::FlowLog"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_gateway_route_table_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::GatewayRouteTableAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_host" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::Host"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::Instance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_instance_connect_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::InstanceConnectEndpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_internet_gateway" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::InternetGateway"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_ipam" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::IPAM"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_ipam_allocation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::IPAMAllocation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_ipam_pool" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::IPAMPool"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_ipam_pool_cidr" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::IPAMPoolCidr"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_ipam_resource_discovery" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::IPAMResourceDiscovery"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_ipam_resource_discovery_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::IPAMResourceDiscoveryAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_ipam_scope" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::IPAMScope"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_key_pair" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::KeyPair"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_launch_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::LaunchTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_local_gateway_route" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::LocalGatewayRoute"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_local_gateway_route_table" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::LocalGatewayRouteTable"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_local_gateway_route_table_virtual_interface_group_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_local_gateway_route_table_vpc_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::LocalGatewayRouteTableVPCAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_nat_gateway" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::NatGateway"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_network_acl" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::NetworkAcl"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_network_insights_access_scope" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::NetworkInsightsAccessScope"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_network_insights_access_scope_analysis" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::NetworkInsightsAccessScopeAnalysis"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_network_insights_analysis" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::NetworkInsightsAnalysis"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_network_insights_path" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::NetworkInsightsPath"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_network_interface" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::NetworkInterface"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_network_interface_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::NetworkInterfaceAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_network_performance_metric_subscription" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::NetworkPerformanceMetricSubscription"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_placement_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::PlacementGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_prefix_list" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::PrefixList"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_route" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::Route"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_route_server" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::RouteServer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_route_server_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::RouteServerAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_route_server_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::RouteServerEndpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_route_server_peer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::RouteServerPeer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_route_server_propagation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::RouteServerPropagation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_route_table" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::RouteTable"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_security_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::SecurityGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_security_group_egress" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::SecurityGroupEgress"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_security_group_ingress" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::SecurityGroupIngress"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_security_group_vpc_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::SecurityGroupVpcAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_snapshot_block_public_access" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::SnapshotBlockPublicAccess"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_spot_fleet" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::SpotFleet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_subnet" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::Subnet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_subnet_cidr_block" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::SubnetCidrBlock"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_subnet_network_acl_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::SubnetNetworkAclAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_subnet_route_table_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::SubnetRouteTableAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_traffic_mirror_filter" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TrafficMirrorFilter"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_traffic_mirror_filter_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TrafficMirrorFilterRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_traffic_mirror_target" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TrafficMirrorTarget"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGateway"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_connect" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayConnect"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_multicast_domain" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayMulticastDomain"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_multicast_domain_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayMulticastDomainAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_multicast_group_member" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayMulticastGroupMember"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_multicast_group_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayMulticastGroupSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_peering_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayPeeringAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_route" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayRoute"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_route_table" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayRouteTable"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_route_table_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayRouteTableAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_route_table_propagation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayRouteTablePropagation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_transit_gateway_vpc_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::TransitGatewayVpcAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_verified_access_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VerifiedAccessEndpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_verified_access_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VerifiedAccessGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_verified_access_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VerifiedAccessInstance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_verified_access_trust_provider" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VerifiedAccessTrustProvider"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_volume" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::Volume"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_volume_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VolumeAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpc" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPC"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpc_block_public_access_exclusion" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPCBlockPublicAccessExclusion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpc_block_public_access_options" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPCBlockPublicAccessOptions"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpc_cidr_block" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPCCidrBlock"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpc_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPCEndpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpc_endpoint_connection_notification" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPCEndpointConnectionNotification"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpc_endpoint_service" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPCEndpointService"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpc_endpoint_service_permissions" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPCEndpointServicePermissions"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpc_gateway_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPCGatewayAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpc_peering_connection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPCPeeringConnection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpcdhcp_options_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPCDHCPOptionsAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpn_connection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPNConnection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpn_connection_route" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPNConnectionRoute"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ec2_vpn_gateway" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EC2::VPNGateway"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecr_public_repository" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECR::PublicRepository"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecr_pull_through_cache_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECR::PullThroughCacheRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecr_registry_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECR::RegistryPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecr_registry_scanning_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECR::RegistryScanningConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecr_replication_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECR::ReplicationConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecr_repository" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECR::Repository"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecr_repository_creation_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECR::RepositoryCreationTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecs_capacity_provider" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECS::CapacityProvider"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecs_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECS::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecs_cluster_capacity_provider_associations" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECS::ClusterCapacityProviderAssociations"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecs_primary_task_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECS::PrimaryTaskSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecs_service" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECS::Service"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecs_task_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECS::TaskDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ecs_task_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ECS::TaskSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_efs_access_point" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EFS::AccessPoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_efs_file_system" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EFS::FileSystem"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_efs_mount_target" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EFS::MountTarget"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_eks_access_entry" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EKS::AccessEntry"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_eks_addon" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EKS::Addon"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_eks_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EKS::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_eks_fargate_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EKS::FargateProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_eks_identity_provider_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EKS::IdentityProviderConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_eks_nodegroup" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EKS::Nodegroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_eks_pod_identity_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EKS::PodIdentityAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticache_global_replication_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElastiCache::GlobalReplicationGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticache_parameter_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElastiCache::ParameterGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticache_serverless_cache" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElastiCache::ServerlessCache"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticache_subnet_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElastiCache::SubnetGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticache_user" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElastiCache::User"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticache_user_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElastiCache::UserGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticbeanstalk_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElasticBeanstalk::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticbeanstalk_application_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElasticBeanstalk::ApplicationVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticbeanstalk_configuration_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElasticBeanstalk::ConfigurationTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticbeanstalk_environment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElasticBeanstalk::Environment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticloadbalancingv2_listener" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElasticLoadBalancingV2::Listener"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_elasticloadbalancingv2_listener_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElasticLoadBalancingV2::ListenerRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_elasticloadbalancingv2_load_balancer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElasticLoadBalancingV2::LoadBalancer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticloadbalancingv2_target_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElasticLoadBalancingV2::TargetGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticloadbalancingv2_trust_store" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElasticLoadBalancingV2::TrustStore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_elasticloadbalancingv2_trust_store_revocation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ElasticLoadBalancingV2::TrustStoreRevocation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_emr_security_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EMR::SecurityConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_emr_step" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EMR::Step"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_emr_studio" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EMR::Studio"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_emr_studio_session_mapping" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EMR::StudioSessionMapping"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_emr_wal_workspace" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EMR::WALWorkspace"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_emrcontainers_virtual_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EMRContainers::VirtualCluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_emrserverless_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EMRServerless::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_entityresolution_id_mapping_workflow" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EntityResolution::IdMappingWorkflow"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_entityresolution_id_namespace" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EntityResolution::IdNamespace"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_entityresolution_matching_workflow" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EntityResolution::MatchingWorkflow"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_entityresolution_policy_statement" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EntityResolution::PolicyStatement"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_entityresolution_schema_mapping" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EntityResolution::SchemaMapping"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_events_api_destination" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Events::ApiDestination"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_events_archive" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Events::Archive"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_events_connection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Events::Connection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_events_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Events::Endpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_events_event_bus" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Events::EventBus"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_events_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Events::Rule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_eventschemas_discoverer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EventSchemas::Discoverer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_eventschemas_registry" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EventSchemas::Registry"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_eventschemas_registry_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EventSchemas::RegistryPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_eventschemas_schema" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EventSchemas::Schema"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_evidently_experiment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Evidently::Experiment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_evidently_feature" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Evidently::Feature"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_evidently_launch" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Evidently::Launch"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_evidently_project" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Evidently::Project"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_evidently_segment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Evidently::Segment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_evs_environment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::EVS::Environment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_finspace_environment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FinSpace::Environment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_fis_experiment_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FIS::ExperimentTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_fis_target_account_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FIS::TargetAccountConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_fms_notification_channel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FMS::NotificationChannel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_fms_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FMS::Policy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_fms_resource_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FMS::ResourceSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_forecast_dataset" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Forecast::Dataset"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_forecast_dataset_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Forecast::DatasetGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_frauddetector_detector" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FraudDetector::Detector"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_frauddetector_entity_type" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FraudDetector::EntityType"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_frauddetector_event_type" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FraudDetector::EventType"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_frauddetector_label" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FraudDetector::Label"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_frauddetector_list" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FraudDetector::List"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_frauddetector_outcome" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FraudDetector::Outcome"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_frauddetector_variable" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FraudDetector::Variable"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_fsx_data_repository_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FSx::DataRepositoryAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_fsx_s3_access_point_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::FSx::S3AccessPointAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_gamelift_alias" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GameLift::Alias"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_gamelift_build" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GameLift::Build"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_gamelift_container_fleet" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GameLift::ContainerFleet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_gamelift_container_group_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GameLift::ContainerGroupDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_gamelift_fleet" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GameLift::Fleet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_gamelift_game_server_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GameLift::GameServerGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_gamelift_game_session_queue" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GameLift::GameSessionQueue"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_gamelift_location" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GameLift::Location"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_gamelift_matchmaking_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GameLift::MatchmakingConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_gamelift_matchmaking_rule_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GameLift::MatchmakingRuleSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_gamelift_script" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GameLift::Script"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_globalaccelerator_accelerator" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GlobalAccelerator::Accelerator"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_globalaccelerator_cross_account_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GlobalAccelerator::CrossAccountAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_globalaccelerator_endpoint_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GlobalAccelerator::EndpointGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_globalaccelerator_listener" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GlobalAccelerator::Listener"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_glue_crawler" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Glue::Crawler"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_glue_database" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Glue::Database"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_glue_job" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Glue::Job"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_glue_registry" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Glue::Registry"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_glue_schema" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Glue::Schema"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_glue_schema_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Glue::SchemaVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_glue_schema_version_metadata" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Glue::SchemaVersionMetadata"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_glue_trigger" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Glue::Trigger"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_glue_usage_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Glue::UsageProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_grafana_workspace" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Grafana::Workspace"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_greengrassv2_component_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GreengrassV2::ComponentVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_greengrassv2_deployment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GreengrassV2::Deployment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_groundstation_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GroundStation::Config"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_groundstation_dataflow_endpoint_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GroundStation::DataflowEndpointGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_groundstation_mission_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GroundStation::MissionProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_guardduty_detector" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GuardDuty::Detector"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_guardduty_filter" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GuardDuty::Filter"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_guardduty_ip_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GuardDuty::IPSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_guardduty_malware_protection_plan" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GuardDuty::MalwareProtectionPlan"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_guardduty_master" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GuardDuty::Master"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_guardduty_member" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GuardDuty::Member"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_guardduty_publishing_destination" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GuardDuty::PublishingDestination"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_guardduty_threat_intel_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::GuardDuty::ThreatIntelSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_healthimaging_datastore" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::HealthImaging::Datastore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_healthlake_fhir_datastore" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::HealthLake::FHIRDatastore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::Group"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_group_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::GroupPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_instance_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::InstanceProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_managed_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::ManagedPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_oidc_provider" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::OIDCProvider"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_role" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::Role"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_role_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::RolePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_saml_provider" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::SAMLProvider"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_server_certificate" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::ServerCertificate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_service_linked_role" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::ServiceLinkedRole"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_user" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::User"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_user_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::UserPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iam_virtual_mfa_device" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IAM::VirtualMFADevice"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_identitystore_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IdentityStore::Group"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_identitystore_group_membership" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IdentityStore::GroupMembership"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_imagebuilder_component" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ImageBuilder::Component"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_imagebuilder_container_recipe" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ImageBuilder::ContainerRecipe"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_imagebuilder_distribution_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ImageBuilder::DistributionConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_imagebuilder_image" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ImageBuilder::Image"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_imagebuilder_image_pipeline" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ImageBuilder::ImagePipeline"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_imagebuilder_image_recipe" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ImageBuilder::ImageRecipe"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_imagebuilder_infrastructure_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ImageBuilder::InfrastructureConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_imagebuilder_lifecycle_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ImageBuilder::LifecyclePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_imagebuilder_workflow" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ImageBuilder::Workflow"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_inspector_assessment_target" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Inspector::AssessmentTarget"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_inspector_assessment_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Inspector::AssessmentTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_inspector_resource_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Inspector::ResourceGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_inspectorv2_cis_scan_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::InspectorV2::CisScanConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_inspectorv2_filter" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::InspectorV2::Filter"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_internetmonitor_monitor" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::InternetMonitor::Monitor"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_invoicing_invoice_unit" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Invoicing::InvoiceUnit"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_account_audit_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::AccountAuditConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_authorizer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::Authorizer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_billing_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::BillingGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_ca_certificate" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::CACertificate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_certificate" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::Certificate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_certificate_provider" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::CertificateProvider"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_command" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::Command"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_custom_metric" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::CustomMetric"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_dimension" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::Dimension"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_domain_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::DomainConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_fleet_metric" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::FleetMetric"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_job_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::JobTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_logging" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::Logging"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_mitigation_action" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::MitigationAction"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::Policy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_provisioning_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::ProvisioningTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_resource_specific_logging" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::ResourceSpecificLogging"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_role_alias" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::RoleAlias"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_scheduled_audit" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::ScheduledAudit"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_security_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::SecurityProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_software_package" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::SoftwarePackage"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_software_package_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::SoftwarePackageVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_thing" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::Thing"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_thing_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::ThingGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_thing_type" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::ThingType"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_topic_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::TopicRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iot_topic_rule_destination" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoT::TopicRuleDestination"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotanalytics_channel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTAnalytics::Channel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotanalytics_dataset" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTAnalytics::Dataset"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotanalytics_datastore" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTAnalytics::Datastore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotanalytics_pipeline" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTAnalytics::Pipeline"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotcoredeviceadvisor_suite_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTCoreDeviceAdvisor::SuiteDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotevents_alarm_model" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTEvents::AlarmModel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotevents_detector_model" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTEvents::DetectorModel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotevents_input" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTEvents::Input"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotfleethub_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTFleetHub::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotfleetwise_campaign" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTFleetWise::Campaign"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotfleetwise_decoder_manifest" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTFleetWise::DecoderManifest"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_iotfleetwise_fleet" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTFleetWise::Fleet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotfleetwise_model_manifest" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTFleetWise::ModelManifest"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotfleetwise_signal_catalog" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTFleetWise::SignalCatalog"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotfleetwise_state_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTFleetWise::StateTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotfleetwise_vehicle" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTFleetWise::Vehicle"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotsitewise_access_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTSiteWise::AccessPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotsitewise_asset" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTSiteWise::Asset"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotsitewise_asset_model" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTSiteWise::AssetModel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotsitewise_dashboard" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTSiteWise::Dashboard"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotsitewise_dataset" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTSiteWise::Dataset"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotsitewise_gateway" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTSiteWise::Gateway"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotsitewise_portal" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTSiteWise::Portal"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotsitewise_project" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTSiteWise::Project"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iottwinmaker_component_type" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTTwinMaker::ComponentType"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_iottwinmaker_entity" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTTwinMaker::Entity"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_iottwinmaker_scene" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTTwinMaker::Scene"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iottwinmaker_sync_job" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTTwinMaker::SyncJob"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iottwinmaker_workspace" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTTwinMaker::Workspace"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotwireless_destination" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTWireless::Destination"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotwireless_device_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTWireless::DeviceProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotwireless_fuota_task" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTWireless::FuotaTask"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotwireless_multicast_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTWireless::MulticastGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotwireless_network_analyzer_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTWireless::NetworkAnalyzerConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotwireless_partner_account" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTWireless::PartnerAccount"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotwireless_service_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTWireless::ServiceProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotwireless_task_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTWireless::TaskDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotwireless_wireless_device" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTWireless::WirelessDevice"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotwireless_wireless_device_import_task" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTWireless::WirelessDeviceImportTask"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_iotwireless_wireless_gateway" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IoTWireless::WirelessGateway"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivs_channel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVS::Channel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivs_encoder_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVS::EncoderConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivs_ingest_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVS::IngestConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivs_playback_key_pair" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVS::PlaybackKeyPair"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivs_playback_restriction_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVS::PlaybackRestrictionPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivs_public_key" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVS::PublicKey"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivs_recording_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVS::RecordingConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivs_stage" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVS::Stage"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivs_storage_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVS::StorageConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivs_stream_key" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVS::StreamKey"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivschat_logging_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVSChat::LoggingConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ivschat_room" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::IVSChat::Room"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kafkaconnect_connector" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::KafkaConnect::Connector"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kafkaconnect_custom_plugin" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::KafkaConnect::CustomPlugin"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kafkaconnect_worker_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::KafkaConnect::WorkerConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kendra_data_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Kendra::DataSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kendra_faq" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Kendra::Faq"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kendra_index" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Kendra::Index"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kendraranking_execution_plan" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::KendraRanking::ExecutionPlan"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kinesis_resource_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Kinesis::ResourcePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kinesis_stream" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Kinesis::Stream"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kinesis_stream_consumer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Kinesis::StreamConsumer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kinesisanalyticsv2_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::KinesisAnalyticsV2::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kinesisfirehose_delivery_stream" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::KinesisFirehose::DeliveryStream"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kinesisvideo_signaling_channel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::KinesisVideo::SignalingChannel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kinesisvideo_stream" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::KinesisVideo::Stream"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kms_alias" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::KMS::Alias"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kms_key" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::KMS::Key"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_kms_replica_key" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::KMS::ReplicaKey"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lakeformation_data_cells_filter" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::LakeFormation::DataCellsFilter"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lakeformation_principal_permissions" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::LakeFormation::PrincipalPermissions"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lakeformation_tag" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::LakeFormation::Tag"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lakeformation_tag_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::LakeFormation::TagAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lambda_alias" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lambda::Alias"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lambda_code_signing_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lambda::CodeSigningConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lambda_event_invoke_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lambda::EventInvokeConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lambda_event_source_mapping" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lambda::EventSourceMapping"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lambda_function" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lambda::Function"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lambda_layer_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lambda::LayerVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lambda_layer_version_permission" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lambda::LayerVersionPermission"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lambda_permission" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lambda::Permission"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lambda_url" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lambda::Url"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lambda_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lambda::Version"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_launchwizard_deployment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::LaunchWizard::Deployment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lex_bot" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lex::Bot"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lex_bot_alias" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lex::BotAlias"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lex_bot_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lex::BotVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lex_resource_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lex::ResourcePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_licensemanager_grant" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::LicenseManager::Grant"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_licensemanager_license" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::LicenseManager::License"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_alarm" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::Alarm"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_bucket" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::Bucket"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_certificate" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::Certificate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_container" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::Container"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_database" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::Database"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_disk" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::Disk"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_distribution" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::Distribution"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::Instance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_instance_snapshot" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::InstanceSnapshot"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_load_balancer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::LoadBalancer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_load_balancer_tls_certificate" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::LoadBalancerTlsCertificate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lightsail_static_ip" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Lightsail::StaticIp"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_location_api_key" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Location::APIKey"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_location_geofence_collection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Location::GeofenceCollection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_location_map" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Location::Map"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_location_place_index" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Location::PlaceIndex"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_location_route_calculator" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Location::RouteCalculator"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_location_tracker" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Location::Tracker"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_location_tracker_consumer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Location::TrackerConsumer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_account_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::AccountPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_delivery" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::Delivery"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_delivery_destination" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::DeliveryDestination"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_delivery_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::DeliverySource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_destination" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::Destination"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_integration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::Integration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_log_anomaly_detector" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::LogAnomalyDetector"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_log_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::LogGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_log_stream" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::LogStream"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_metric_filter" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::MetricFilter"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_query_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::QueryDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_resource_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::ResourcePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_subscription_filter" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::SubscriptionFilter"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_logs_transformer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Logs::Transformer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_lookoutequipment_inference_scheduler" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::LookoutEquipment::InferenceScheduler"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lookoutmetrics_alert" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::LookoutMetrics::Alert"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lookoutmetrics_anomaly_detector" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::LookoutMetrics::AnomalyDetector"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_lookoutvision_project" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::LookoutVision::Project"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_m2_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::M2::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_m2_deployment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::M2::Deployment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_m2_environment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::M2::Environment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_macie_allow_list" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Macie::AllowList"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_macie_custom_data_identifier" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Macie::CustomDataIdentifier"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_macie_findings_filter" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Macie::FindingsFilter"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_macie_session" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Macie::Session"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_managedblockchain_accessor" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ManagedBlockchain::Accessor"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediaconnect_bridge" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaConnect::Bridge"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediaconnect_bridge_output" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaConnect::BridgeOutput"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediaconnect_bridge_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaConnect::BridgeSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediaconnect_flow" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaConnect::Flow"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediaconnect_flow_entitlement" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaConnect::FlowEntitlement"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediaconnect_flow_output" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaConnect::FlowOutput"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediaconnect_flow_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaConnect::FlowSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediaconnect_flow_vpc_interface" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaConnect::FlowVpcInterface"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediaconnect_gateway" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaConnect::Gateway"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_medialive_channel_placement_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaLive::ChannelPlacementGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_medialive_cloudwatch_alarm_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaLive::CloudWatchAlarmTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_medialive_cloudwatch_alarm_template_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaLive::CloudWatchAlarmTemplateGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_medialive_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaLive::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_medialive_event_bridge_rule_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaLive::EventBridgeRuleTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_medialive_event_bridge_rule_template_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaLive::EventBridgeRuleTemplateGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_medialive_multiplex" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaLive::Multiplex"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_medialive_multiplexprogram" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaLive::Multiplexprogram"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_medialive_network" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaLive::Network"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_medialive_sdi_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaLive::SdiSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_medialive_signal_map" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaLive::SignalMap"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediapackage_asset" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaPackage::Asset"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediapackage_channel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaPackage::Channel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediapackage_origin_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaPackage::OriginEndpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediapackage_packaging_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaPackage::PackagingConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediapackage_packaging_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaPackage::PackagingGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediapackagev2_channel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaPackageV2::Channel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediapackagev2_channel_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaPackageV2::ChannelGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediapackagev2_channel_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaPackageV2::ChannelPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediapackagev2_origin_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaPackageV2::OriginEndpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediapackagev2_origin_endpoint_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaPackageV2::OriginEndpointPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediatailor_channel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaTailor::Channel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediatailor_channel_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaTailor::ChannelPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediatailor_live_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaTailor::LiveSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediatailor_playback_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaTailor::PlaybackConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_mediatailor_source_location" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaTailor::SourceLocation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mediatailor_vod_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MediaTailor::VodSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_memorydb_acl" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MemoryDB::ACL"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_memorydb_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MemoryDB::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_memorydb_multi_region_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MemoryDB::MultiRegionCluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_memorydb_parameter_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MemoryDB::ParameterGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_memorydb_subnet_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MemoryDB::SubnetGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_memorydb_user" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MemoryDB::User"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mpa_approval_team" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MPA::ApprovalTeam"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mpa_identity_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MPA::IdentitySource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_msk_batch_scram_secret" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MSK::BatchScramSecret"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_msk_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MSK::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_msk_cluster_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MSK::ClusterPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_msk_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MSK::Configuration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_msk_replicator" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MSK::Replicator"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_msk_serverless_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MSK::ServerlessCluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_msk_vpc_connection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MSK::VpcConnection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_mwaa_environment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::MWAA::Environment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_neptune_db_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Neptune::DBCluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_neptune_db_cluster_parameter_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Neptune::DBClusterParameterGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_neptune_db_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Neptune::DBInstance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_neptune_db_parameter_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Neptune::DBParameterGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_neptune_db_subnet_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Neptune::DBSubnetGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_neptunegraph_graph" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NeptuneGraph::Graph"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_neptunegraph_private_graph_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NeptuneGraph::PrivateGraphEndpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkfirewall_firewall" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkFirewall::Firewall"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkfirewall_firewall_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkFirewall::FirewallPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkfirewall_logging_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkFirewall::LoggingConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkfirewall_rule_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkFirewall::RuleGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkfirewall_tls_inspection_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkFirewall::TLSInspectionConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkfirewall_vpc_endpoint_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkFirewall::VpcEndpointAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_connect_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::ConnectAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_connect_peer" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::ConnectPeer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_core_network" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::CoreNetwork"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_customer_gateway_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::CustomerGatewayAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_device" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::Device"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_direct_connect_gateway_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::DirectConnectGatewayAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_global_network" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::GlobalNetwork"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_link" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::Link"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_link_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::LinkAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_site" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::Site"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_site_to_site_vpn_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::SiteToSiteVpnAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_transit_gateway_peering" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::TransitGatewayPeering"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_transit_gateway_registration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::TransitGatewayRegistration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_transit_gateway_route_table_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::TransitGatewayRouteTableAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_networkmanager_vpc_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NetworkManager::VpcAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_nimblestudio_launch_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NimbleStudio::LaunchProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_nimblestudio_streaming_image" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NimbleStudio::StreamingImage"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_nimblestudio_studio" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NimbleStudio::Studio"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_nimblestudio_studio_component" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NimbleStudio::StudioComponent"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_notifications_channel_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Notifications::ChannelAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_notifications_event_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Notifications::EventRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_notifications_managed_notification_account_contact_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Notifications::ManagedNotificationAccountContactAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_notifications_managed_notification_additional_channel_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Notifications::ManagedNotificationAdditionalChannelAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_notifications_notification_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Notifications::NotificationConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_notifications_notification_hub" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Notifications::NotificationHub"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_notificationscontacts_email_contact" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::NotificationsContacts::EmailContact"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_oam_link" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Oam::Link"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_oam_sink" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Oam::Sink"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_odb_cloud_autonomous_vm_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ODB::CloudAutonomousVmCluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_odb_cloud_exadata_infrastructure" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ODB::CloudExadataInfrastructure"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_odb_cloud_vm_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ODB::CloudVmCluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_odb_odb_network" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ODB::OdbNetwork"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_omics_annotation_store" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Omics::AnnotationStore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_omics_reference_store" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Omics::ReferenceStore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_omics_run_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Omics::RunGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_omics_sequence_store" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Omics::SequenceStore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_omics_variant_store" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Omics::VariantStore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_omics_workflow" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Omics::Workflow"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_omics_workflow_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Omics::WorkflowVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_opensearchserverless_access_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::OpenSearchServerless::AccessPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_opensearchserverless_collection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::OpenSearchServerless::Collection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_opensearchserverless_index" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::OpenSearchServerless::Index"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_opensearchserverless_lifecycle_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::OpenSearchServerless::LifecyclePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_opensearchserverless_security_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::OpenSearchServerless::SecurityConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_opensearchserverless_security_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::OpenSearchServerless::SecurityPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_opensearchserverless_vpc_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::OpenSearchServerless::VpcEndpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_opensearchservice_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::OpenSearchService::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_opensearchservice_domain" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::OpenSearchService::Domain"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_opsworkscm_server" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::OpsWorksCM::Server"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = false
}
resource_schema "aws_organizations_account" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Organizations::Account"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_organizations_organization" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Organizations::Organization"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_organizations_organizational_unit" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Organizations::OrganizationalUnit"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_organizations_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Organizations::Policy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_organizations_resource_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Organizations::ResourcePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_osis_pipeline" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::OSIS::Pipeline"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_panorama_application_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Panorama::ApplicationInstance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_panorama_package" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Panorama::Package"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_panorama_package_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Panorama::PackageVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_paymentcryptography_alias" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PaymentCryptography::Alias"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_paymentcryptography_key" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PaymentCryptography::Key"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pcaconnectorad_connector" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PCAConnectorAD::Connector"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pcaconnectorad_directory_registration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PCAConnectorAD::DirectoryRegistration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pcaconnectorad_service_principal_name" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PCAConnectorAD::ServicePrincipalName"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pcaconnectorad_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PCAConnectorAD::Template"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pcaconnectorad_template_group_access_control_entry" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PCAConnectorAD::TemplateGroupAccessControlEntry"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pcaconnectorscep_challenge" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PCAConnectorSCEP::Challenge"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pcaconnectorscep_connector" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PCAConnectorSCEP::Connector"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pcs_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PCS::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pcs_compute_node_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PCS::ComputeNodeGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pcs_queue" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::PCS::Queue"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_personalize_dataset" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Personalize::Dataset"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_personalize_dataset_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Personalize::DatasetGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_personalize_schema" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Personalize::Schema"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_personalize_solution" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Personalize::Solution"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pinpoint_in_app_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Pinpoint::InAppTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_pipes_pipe" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Pipes::Pipe"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_proton_environment_account_connection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Proton::EnvironmentAccountConnection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_proton_environment_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Proton::EnvironmentTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_proton_service_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Proton::ServiceTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_qbusiness_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QBusiness::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_qbusiness_data_accessor" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QBusiness::DataAccessor"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_qbusiness_data_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QBusiness::DataSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_qbusiness_index" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QBusiness::Index"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_qbusiness_permission" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QBusiness::Permission"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_qbusiness_plugin" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QBusiness::Plugin"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_qbusiness_retriever" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QBusiness::Retriever"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_qbusiness_web_experience" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QBusiness::WebExperience"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_qldb_stream" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QLDB::Stream"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_quicksight_analysis" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QuickSight::Analysis"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_quicksight_custom_permissions" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QuickSight::CustomPermissions"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_quicksight_dashboard" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QuickSight::Dashboard"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_quicksight_data_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QuickSight::DataSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_quicksight_data_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QuickSight::DataSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_quicksight_folder" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QuickSight::Folder"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_quicksight_refresh_schedule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QuickSight::RefreshSchedule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_quicksight_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QuickSight::Template"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_quicksight_theme" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QuickSight::Theme"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_quicksight_topic" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QuickSight::Topic"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_quicksight_vpc_connection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::QuickSight::VPCConnection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ram_permission" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RAM::Permission"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ram_resource_share" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RAM::ResourceShare"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rbin_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Rbin::Rule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_custom_db_engine_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::CustomDBEngineVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_db_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::DBCluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_db_cluster_parameter_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::DBClusterParameterGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_db_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::DBInstance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_db_parameter_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::DBParameterGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_db_proxy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::DBProxy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_db_proxy_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::DBProxyEndpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_db_proxy_target_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::DBProxyTargetGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_db_shard_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::DBShardGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_db_subnet_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::DBSubnetGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_event_subscription" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::EventSubscription"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_global_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::GlobalCluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_integration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::Integration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rds_option_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RDS::OptionGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_redshift_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Redshift::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_redshift_cluster_parameter_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Redshift::ClusterParameterGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_redshift_cluster_subnet_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Redshift::ClusterSubnetGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_redshift_endpoint_access" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Redshift::EndpointAccess"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_redshift_endpoint_authorization" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Redshift::EndpointAuthorization"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_redshift_event_subscription" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Redshift::EventSubscription"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_redshift_integration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Redshift::Integration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_redshift_scheduled_action" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Redshift::ScheduledAction"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_redshiftserverless_namespace" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RedshiftServerless::Namespace"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_redshiftserverless_snapshot" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RedshiftServerless::Snapshot"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_redshiftserverless_workgroup" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RedshiftServerless::Workgroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_refactorspaces_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RefactorSpaces::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_refactorspaces_environment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RefactorSpaces::Environment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_refactorspaces_route" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RefactorSpaces::Route"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_refactorspaces_service" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RefactorSpaces::Service"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rekognition_collection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Rekognition::Collection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rekognition_project" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Rekognition::Project"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rekognition_stream_processor" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Rekognition::StreamProcessor"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_resiliencehub_app" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ResilienceHub::App"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_resiliencehub_resiliency_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ResilienceHub::ResiliencyPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_resourceexplorer2_default_view_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ResourceExplorer2::DefaultViewAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_resourceexplorer2_index" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ResourceExplorer2::Index"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_resourceexplorer2_view" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ResourceExplorer2::View"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_resourcegroups_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ResourceGroups::Group"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_resourcegroups_tag_sync_task" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ResourceGroups::TagSyncTask"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_robomaker_fleet" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RoboMaker::Fleet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_robomaker_robot" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RoboMaker::Robot"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_robomaker_robot_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RoboMaker::RobotApplication"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_robomaker_robot_application_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RoboMaker::RobotApplicationVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_robomaker_simulation_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RoboMaker::SimulationApplication"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_robomaker_simulation_application_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RoboMaker::SimulationApplicationVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rolesanywhere_crl" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RolesAnywhere::CRL"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rolesanywhere_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RolesAnywhere::Profile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rolesanywhere_trust_anchor" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RolesAnywhere::TrustAnchor"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53_cidr_collection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53::CidrCollection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53_dnssec" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53::DNSSEC"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53_health_check" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53::HealthCheck"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53_hosted_zone" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53::HostedZone"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53_key_signing_key" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53::KeySigningKey"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53_record_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53::RecordSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53profiles_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Profiles::Profile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53profiles_profile_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Profiles::ProfileAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53profiles_profile_resource_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Profiles::ProfileResourceAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53recoverycontrol_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53RecoveryControl::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53recoverycontrol_control_panel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53RecoveryControl::ControlPanel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53recoverycontrol_routing_control" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53RecoveryControl::RoutingControl"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53recoverycontrol_safety_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53RecoveryControl::SafetyRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53recoveryreadiness_cell" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53RecoveryReadiness::Cell"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53recoveryreadiness_readiness_check" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53RecoveryReadiness::ReadinessCheck"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53recoveryreadiness_recovery_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53RecoveryReadiness::RecoveryGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53recoveryreadiness_resource_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53RecoveryReadiness::ResourceSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53resolver_firewall_domain_list" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Resolver::FirewallDomainList"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53resolver_firewall_rule_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Resolver::FirewallRuleGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53resolver_firewall_rule_group_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Resolver::FirewallRuleGroupAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53resolver_outpost_resolver" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Resolver::OutpostResolver"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53resolver_resolver_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Resolver::ResolverConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53resolver_resolver_dnssec_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Resolver::ResolverDNSSECConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53resolver_resolver_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Resolver::ResolverEndpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53resolver_resolver_query_logging_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Resolver::ResolverQueryLoggingConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53resolver_resolver_query_logging_config_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53resolver_resolver_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Resolver::ResolverRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_route53resolver_resolver_rule_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Route53Resolver::ResolverRuleAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_rum_app_monitor" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::RUM::AppMonitor"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3_access_grant" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3::AccessGrant"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3_access_grants_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3::AccessGrantsInstance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3_access_grants_location" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3::AccessGrantsLocation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3_access_point" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3::AccessPoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3_bucket" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3::Bucket"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3_bucket_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3::BucketPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3_multi_region_access_point" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3::MultiRegionAccessPoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3_multi_region_access_point_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3::MultiRegionAccessPointPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3_storage_lens" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3::StorageLens"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3_storage_lens_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3::StorageLensGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3express_access_point" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Express::AccessPoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3express_bucket_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Express::BucketPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3express_directory_bucket" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Express::DirectoryBucket"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3objectlambda_access_point" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3ObjectLambda::AccessPoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3objectlambda_access_point_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3ObjectLambda::AccessPointPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3outposts_access_point" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Outposts::AccessPoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3outposts_bucket" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Outposts::Bucket"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3outposts_bucket_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Outposts::BucketPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3outposts_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Outposts::Endpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3tables_namespace" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Tables::Namespace"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3tables_table" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Tables::Table"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3tables_table_bucket" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Tables::TableBucket"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3tables_table_bucket_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Tables::TableBucketPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_s3tables_table_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::S3Tables::TablePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_app" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::App"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_app_image_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::AppImageConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_cluster" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::Cluster"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_data_quality_job_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::DataQualityJobDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_device" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::Device"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_device_fleet" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::DeviceFleet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_domain" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::Domain"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_endpoint" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::Endpoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_feature_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::FeatureGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_image" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::Image"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_image_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::ImageVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_inference_component" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::InferenceComponent"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_inference_experiment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::InferenceExperiment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_mlflow_tracking_server" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::MlflowTrackingServer"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_model_bias_job_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::ModelBiasJobDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_model_card" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::ModelCard"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_sagemaker_model_explainability_job_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::ModelExplainabilityJobDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_model_package" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::ModelPackage"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_model_package_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::ModelPackageGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_model_quality_job_definition" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::ModelQualityJobDefinition"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_monitoring_schedule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::MonitoringSchedule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_partner_app" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::PartnerApp"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_pipeline" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::Pipeline"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_project" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::Project"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_space" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::Space"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_studio_lifecycle_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::StudioLifecycleConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sagemaker_user_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SageMaker::UserProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_scheduler_schedule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Scheduler::Schedule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_scheduler_schedule_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Scheduler::ScheduleGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_secretsmanager_resource_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecretsManager::ResourcePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_secretsmanager_rotation_schedule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecretsManager::RotationSchedule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_secretsmanager_secret" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecretsManager::Secret"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_secretsmanager_secret_target_attachment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecretsManager::SecretTargetAttachment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_aggregator_v2" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::AggregatorV2"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_automation_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::AutomationRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_securityhub_automation_rule_v2" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::AutomationRuleV2"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_configuration_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::ConfigurationPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_delegated_admin" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::DelegatedAdmin"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_finding_aggregator" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::FindingAggregator"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_hub" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::Hub"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_hub_v2" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::HubV2"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_insight" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::Insight"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_organization_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::OrganizationConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_policy_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::PolicyAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_product_subscription" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::ProductSubscription"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_security_control" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::SecurityControl"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securityhub_standard" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityHub::Standard"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securitylake_aws_log_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityLake::AwsLogSource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securitylake_data_lake" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityLake::DataLake"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_securitylake_subscriber" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityLake::Subscriber"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_securitylake_subscriber_notification" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SecurityLake::SubscriberNotification"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_servicecatalog_cloudformation_provisioned_product" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ServiceCatalog::CloudFormationProvisionedProduct"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_servicecatalog_service_action" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ServiceCatalog::ServiceAction"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_servicecatalog_service_action_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ServiceCatalog::ServiceActionAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_servicecatalogappregistry_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ServiceCatalogAppRegistry::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_servicecatalogappregistry_attribute_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ServiceCatalogAppRegistry::AttributeGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_servicecatalogappregistry_attribute_group_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_servicecatalogappregistry_resource_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::ServiceCatalogAppRegistry::ResourceAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_configuration_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::ConfigurationSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_configuration_set_event_destination" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::ConfigurationSetEventDestination"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_contact_list" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::ContactList"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_dedicated_ip_pool" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::DedicatedIpPool"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_email_identity" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::EmailIdentity"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_mail_manager_addon_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::MailManagerAddonInstance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_mail_manager_addon_subscription" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::MailManagerAddonSubscription"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_mail_manager_address_list" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::MailManagerAddressList"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_mail_manager_archive" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::MailManagerArchive"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_mail_manager_ingress_point" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::MailManagerIngressPoint"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_mail_manager_relay" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::MailManagerRelay"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_mail_manager_rule_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::MailManagerRuleSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_mail_manager_traffic_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::MailManagerTrafficPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::Template"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ses_vdm_attributes" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SES::VdmAttributes"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_shield_drt_access" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Shield::DRTAccess"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_shield_proactive_engagement" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Shield::ProactiveEngagement"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_shield_protection" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Shield::Protection"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_shield_protection_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Shield::ProtectionGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_signer_profile_permission" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Signer::ProfilePermission"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_signer_signing_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Signer::SigningProfile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_simspaceweaver_simulation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SimSpaceWeaver::Simulation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sns_subscription" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SNS::Subscription"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sns_topic" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SNS::Topic"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sns_topic_inline_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SNS::TopicInlinePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sqs_queue" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SQS::Queue"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sqs_queue_inline_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SQS::QueueInlinePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssm_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSM::Association"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssm_document" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSM::Document"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssm_parameter" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSM::Parameter"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssm_patch_baseline" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSM::PatchBaseline"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssm_resource_data_sync" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSM::ResourceDataSync"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssm_resource_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSM::ResourcePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssmcontacts_contact" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSMContacts::Contact"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssmcontacts_contact_channel" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSMContacts::ContactChannel"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssmcontacts_plan" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSMContacts::Plan"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssmcontacts_rotation" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSMContacts::Rotation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssmguiconnect_preferences" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSMGuiConnect::Preferences"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssmincidents_replication_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSMIncidents::ReplicationSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssmincidents_response_plan" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSMIncidents::ResponsePlan"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_ssmquicksetup_configuration_manager" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSMQuickSetup::ConfigurationManager"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sso_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSO::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sso_application_assignment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSO::ApplicationAssignment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sso_assignment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSO::Assignment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sso_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSO::Instance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sso_instance_access_control_attribute_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSO::InstanceAccessControlAttributeConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_sso_permission_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SSO::PermissionSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_stepfunctions_activity" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::StepFunctions::Activity"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_stepfunctions_state_machine" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::StepFunctions::StateMachine"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_stepfunctions_state_machine_alias" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::StepFunctions::StateMachineAlias"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_stepfunctions_state_machine_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::StepFunctions::StateMachineVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_supportapp_account_alias" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SupportApp::AccountAlias"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_supportapp_slack_channel_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SupportApp::SlackChannelConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_supportapp_slack_workspace_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SupportApp::SlackWorkspaceConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_synthetics_canary" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Synthetics::Canary"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_synthetics_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Synthetics::Group"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_systemsmanagersap_application" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::SystemsManagerSAP::Application"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_timestream_database" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Timestream::Database"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_timestream_influx_db_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Timestream::InfluxDBInstance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_timestream_scheduled_query" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Timestream::ScheduledQuery"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_timestream_table" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Timestream::Table"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_transfer_agreement" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Transfer::Agreement"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_transfer_certificate" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Transfer::Certificate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_transfer_connector" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Transfer::Connector"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_transfer_profile" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Transfer::Profile"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_transfer_server" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Transfer::Server"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_transfer_user" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Transfer::User"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_transfer_web_app" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Transfer::WebApp"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_transfer_workflow" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Transfer::Workflow"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_verifiedpermissions_identity_source" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VerifiedPermissions::IdentitySource"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_verifiedpermissions_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VerifiedPermissions::Policy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_verifiedpermissions_policy_store" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VerifiedPermissions::PolicyStore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_verifiedpermissions_policy_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VerifiedPermissions::PolicyTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_voiceid_domain" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VoiceID::Domain"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_access_log_subscription" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::AccessLogSubscription"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_auth_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::AuthPolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_listener" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::Listener"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_resource_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::ResourceConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_resource_gateway" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::ResourceGateway"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_resource_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::ResourcePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::Rule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_service" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::Service"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_service_network" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::ServiceNetwork"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_service_network_resource_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::ServiceNetworkResourceAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_service_network_service_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::ServiceNetworkServiceAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_service_network_vpc_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::ServiceNetworkVpcAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_vpclattice_target_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::VpcLattice::TargetGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wafv2_ip_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WAFv2::IPSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wafv2_logging_configuration" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WAFv2::LoggingConfiguration"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wafv2_regex_pattern_set" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WAFv2::RegexPatternSet"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wafv2_rule_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WAFv2::RuleGroup"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_wafv2_web_acl" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WAFv2::WebACL"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}
resource_schema "aws_wafv2_web_acl_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WAFv2::WebACLAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_ai_agent" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::AIAgent"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_ai_agent_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::AIAgentVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_ai_guardrail" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::AIGuardrail"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_ai_guardrail_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::AIGuardrailVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_ai_prompt" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::AIPrompt"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_ai_prompt_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::AIPromptVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_assistant" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::Assistant"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_assistant_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::AssistantAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_knowledge_base" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::KnowledgeBase"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_message_template" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::MessageTemplate"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_message_template_version" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::MessageTemplateVersion"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_wisdom_quick_response" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::Wisdom::QuickResponse"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspaces_connection_alias" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpaces::ConnectionAlias"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspaces_workspaces_pool" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpaces::WorkspacesPool"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesinstances_volume" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkspacesInstances::Volume"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesinstances_volume_association" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkspacesInstances::VolumeAssociation"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesinstances_workspace_instance" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkspacesInstances::WorkspaceInstance"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesthinclient_environment" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpacesThinClient::Environment"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesweb_browser_settings" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpacesWeb::BrowserSettings"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesweb_data_protection_settings" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpacesWeb::DataProtectionSettings"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesweb_identity_provider" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpacesWeb::IdentityProvider"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = true
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesweb_ip_access_settings" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpacesWeb::IpAccessSettings"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesweb_network_settings" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpacesWeb::NetworkSettings"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesweb_portal" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpacesWeb::Portal"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesweb_trust_store" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpacesWeb::TrustStore"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesweb_user_access_logging_settings" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpacesWeb::UserAccessLoggingSettings"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_workspacesweb_user_settings" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::WorkSpacesWeb::UserSettings"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_xray_group" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::XRay::Group"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_xray_resource_policy" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::XRay::ResourcePolicy"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_xray_sampling_rule" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::XRay::SamplingRule"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
resource_schema "aws_xray_transaction_search_config" {
  cloudformation_schema_path               = ""
  cloudformation_type_name                 = "AWS::XRay::TransactionSearchConfig"
  suppression_reason                       = ""
  suppress_plural_data_source_generation   = false
  suppress_resource_generation             = false
  suppress_singular_data_source_generation = false
}
