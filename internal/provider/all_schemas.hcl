# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

defaults {
  schema_cache_directory     = "../service/cloudformation/schemas"
  terraform_type_name_prefix = "awscc"
}

meta_schema {
  path = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}

# 975 CloudFormation resource types schemas are available for use with the Cloud Control API.

resource_schema "aws_acmpca_certificate" {
  cloudformation_type_name               = "AWS::ACMPCA::Certificate"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_acmpca_certificate_authority" {
  cloudformation_type_name = "AWS::ACMPCA::CertificateAuthority"
}

resource_schema "aws_acmpca_certificate_authority_activation" {
  cloudformation_type_name               = "AWS::ACMPCA::CertificateAuthorityActivation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_acmpca_permission" {
  cloudformation_type_name               = "AWS::ACMPCA::Permission"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_aps_rule_groups_namespace" {
  cloudformation_type_name               = "AWS::APS::RuleGroupsNamespace"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_aps_scraper" {
  cloudformation_type_name = "AWS::APS::Scraper"
}

resource_schema "aws_aps_workspace" {
  cloudformation_type_name = "AWS::APS::Workspace"
}

resource_schema "aws_arczonalshift_zonal_autoshift_configuration" {
  cloudformation_type_name = "AWS::ARCZonalShift::ZonalAutoshiftConfiguration"
}

resource_schema "aws_accessanalyzer_analyzer" {
  cloudformation_type_name = "AWS::AccessAnalyzer::Analyzer"
}

resource_schema "aws_amplify_app" {
  cloudformation_type_name = "AWS::Amplify::App"
}

resource_schema "aws_amplify_branch" {
  cloudformation_type_name               = "AWS::Amplify::Branch"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_amplify_domain" {
  cloudformation_type_name               = "AWS::Amplify::Domain"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_amplifyuibuilder_component" {
  cloudformation_type_name               = "AWS::AmplifyUIBuilder::Component"
  suppress_plural_data_source_generation = true

  # Suppression Reason:
  # Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_amplifyuibuilder_form" {
  cloudformation_type_name               = "AWS::AmplifyUIBuilder::Form"
  suppress_plural_data_source_generation = true

  # Suppression Reason:
  # Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_amplifyuibuilder_theme" {
  cloudformation_type_name               = "AWS::AmplifyUIBuilder::Theme"
  suppress_plural_data_source_generation = true

  # Suppression Reason:
  # Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_apigateway_account" {
  cloudformation_type_name               = "AWS::ApiGateway::Account"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigateway_api_key" {
  cloudformation_type_name = "AWS::ApiGateway::ApiKey"
}

resource_schema "aws_apigateway_authorizer" {
  cloudformation_type_name               = "AWS::ApiGateway::Authorizer"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigateway_base_path_mapping" {
  cloudformation_type_name = "AWS::ApiGateway::BasePathMapping"
}

resource_schema "aws_apigateway_client_certificate" {
  cloudformation_type_name = "AWS::ApiGateway::ClientCertificate"
}

resource_schema "aws_apigateway_deployment" {
  cloudformation_type_name               = "AWS::ApiGateway::Deployment"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigateway_documentation_part" {
  cloudformation_type_name               = "AWS::ApiGateway::DocumentationPart"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigateway_documentation_version" {
  cloudformation_type_name               = "AWS::ApiGateway::DocumentationVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigateway_domain_name" {
  cloudformation_type_name = "AWS::ApiGateway::DomainName"
}

resource_schema "aws_apigateway_method" {
  cloudformation_type_name               = "AWS::ApiGateway::Method"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigateway_model" {
  cloudformation_type_name               = "AWS::ApiGateway::Model"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigateway_request_validator" {
  cloudformation_type_name               = "AWS::ApiGateway::RequestValidator"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigateway_resource" {
  cloudformation_type_name               = "AWS::ApiGateway::Resource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigateway_rest_api" {
  cloudformation_type_name = "AWS::ApiGateway::RestApi"
}

resource_schema "aws_apigateway_stage" {
  cloudformation_type_name               = "AWS::ApiGateway::Stage"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigateway_usage_plan" {
  cloudformation_type_name = "AWS::ApiGateway::UsagePlan"
}

resource_schema "aws_apigateway_usage_plan_key" {
  cloudformation_type_name               = "AWS::ApiGateway::UsagePlanKey"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigateway_vpc_link" {
  cloudformation_type_name = "AWS::ApiGateway::VpcLink"
}

resource_schema "aws_apigatewayv2_api" {
  cloudformation_type_name = "AWS::ApiGatewayV2::Api"
}

resource_schema "aws_apigatewayv2_api_mapping" {
  cloudformation_type_name               = "AWS::ApiGatewayV2::ApiMapping"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigatewayv2_authorizer" {
  cloudformation_type_name               = "AWS::ApiGatewayV2::Authorizer"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigatewayv2_deployment" {
  cloudformation_type_name               = "AWS::ApiGatewayV2::Deployment"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigatewayv2_domain_name" {
  cloudformation_type_name = "AWS::ApiGatewayV2::DomainName"
}

resource_schema "aws_apigatewayv2_integration_response" {
  cloudformation_type_name               = "AWS::ApiGatewayV2::IntegrationResponse"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigatewayv2_model" {
  cloudformation_type_name               = "AWS::ApiGatewayV2::Model"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigatewayv2_route" {
  cloudformation_type_name               = "AWS::ApiGatewayV2::Route"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apigatewayv2_route_response" {
  cloudformation_type_name               = "AWS::ApiGatewayV2::RouteResponse"
  suppress_plural_data_source_generation = true

  # Suppression Reason: ResponseParameters is of unsupported type: ""
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1504
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_apigatewayv2_vpc_link" {
  cloudformation_type_name = "AWS::ApiGatewayV2::VpcLink"
}

resource_schema "aws_appconfig_application" {
  cloudformation_type_name = "AWS::AppConfig::Application"
}

resource_schema "aws_appconfig_configuration_profile" {
  cloudformation_type_name               = "AWS::AppConfig::ConfigurationProfile"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appconfig_environment" {
  cloudformation_type_name               = "AWS::AppConfig::Environment"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appconfig_extension" {
  cloudformation_type_name = "AWS::AppConfig::Extension"

  # Suppression Reason: Actions is of unsupported type: key-value map of set of object.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1508
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_appconfig_extension_association" {
  cloudformation_type_name = "AWS::AppConfig::ExtensionAssociation"
}

resource_schema "aws_appconfig_hosted_configuration_version" {
  cloudformation_type_name               = "AWS::AppConfig::HostedConfigurationVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appflow_connector" {
  cloudformation_type_name = "AWS::AppFlow::Connector"
}

resource_schema "aws_appflow_connector_profile" {
  cloudformation_type_name = "AWS::AppFlow::ConnectorProfile"

  # Latest schema updates are suppressed.
  # git checkout internal/service/cloudformation/schemas/AWS_AppFlow_ConnectorProfile.json
  # Suppression Reason: isSandboxEnvironment overwrites IsSandboxEnvironment for Terraform attribute is_sandbox_environment.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1526
}

resource_schema "aws_appflow_flow" {
  cloudformation_type_name = "AWS::AppFlow::Flow"
}

resource_schema "aws_appintegrations_application" {
  cloudformation_type_name = "AWS::AppIntegrations::Application"
}

resource_schema "aws_appintegrations_data_integration" {
  cloudformation_type_name = "AWS::AppIntegrations::DataIntegration"

  # Suppression Reason: ObjectConfiguration is of unsupported type: key-value map of key-value map.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1509
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_appintegrations_event_integration" {
  cloudformation_type_name = "AWS::AppIntegrations::EventIntegration"
}

resource_schema "aws_applicationautoscaling_scalable_target" {
  cloudformation_type_name = "AWS::ApplicationAutoScaling::ScalableTarget"
}

resource_schema "aws_applicationautoscaling_scaling_policy" {
  cloudformation_type_name               = "AWS::ApplicationAutoScaling::ScalingPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_apprunner_auto_scaling_configuration" {
  cloudformation_type_name = "AWS::AppRunner::AutoScalingConfiguration"
}

resource_schema "aws_apprunner_observability_configuration" {
  cloudformation_type_name = "AWS::AppRunner::ObservabilityConfiguration"
}

resource_schema "aws_apprunner_service" {
  cloudformation_type_name = "AWS::AppRunner::Service"
}

resource_schema "aws_apprunner_vpc_connector" {
  cloudformation_type_name = "AWS::AppRunner::VpcConnector"
}

resource_schema "aws_apprunner_vpc_ingress_connection" {
  cloudformation_type_name = "AWS::AppRunner::VpcIngressConnection"
}

resource_schema "aws_appstream_app_block" {
  cloudformation_type_name               = "AWS::AppStream::AppBlock"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appstream_app_block_builder" {
  cloudformation_type_name               = "AWS::AppStream::AppBlockBuilder"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appstream_application" {
  cloudformation_type_name               = "AWS::AppStream::Application"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appstream_application_entitlement_association" {
  cloudformation_type_name               = "AWS::AppStream::ApplicationEntitlementAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appstream_application_fleet_association" {
  cloudformation_type_name               = "AWS::AppStream::ApplicationFleetAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appstream_directory_config" {
  cloudformation_type_name = "AWS::AppStream::DirectoryConfig"
}

resource_schema "aws_appstream_entitlement" {
  cloudformation_type_name               = "AWS::AppStream::Entitlement"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appstream_image_builder" {
  cloudformation_type_name = "AWS::AppStream::ImageBuilder"
}

resource_schema "aws_appsync_domain_name" {
  cloudformation_type_name = "AWS::AppSync::DomainName"
}

resource_schema "aws_appsync_domain_name_api_association" {
  cloudformation_type_name               = "AWS::AppSync::DomainNameApiAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appsync_function_configuration" {
  cloudformation_type_name               = "AWS::AppSync::FunctionConfiguration"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appsync_resolver" {
  cloudformation_type_name               = "AWS::AppSync::Resolver"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_appsync_source_api_association" {
  cloudformation_type_name               = "AWS::AppSync::SourceApiAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_applicationinsights_application" {
  cloudformation_type_name = "AWS::ApplicationInsights::Application"
}

resource_schema "aws_athena_capacity_reservation" {
  cloudformation_type_name = "AWS::Athena::CapacityReservation"
}

resource_schema "aws_athena_data_catalog" {
  cloudformation_type_name = "AWS::Athena::DataCatalog"
}

resource_schema "aws_athena_named_query" {
  cloudformation_type_name = "AWS::Athena::NamedQuery"
}

resource_schema "aws_athena_prepared_statement" {
  cloudformation_type_name               = "AWS::Athena::PreparedStatement"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_athena_work_group" {
  cloudformation_type_name = "AWS::Athena::WorkGroup"
}

resource_schema "aws_auditmanager_assessment" {
  cloudformation_type_name               = "AWS::AuditManager::Assessment"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_autoscaling_auto_scaling_group" {
  cloudformation_type_name = "AWS::AutoScaling::AutoScalingGroup"
}

resource_schema "aws_autoscaling_launch_configuration" {
  cloudformation_type_name = "AWS::AutoScaling::LaunchConfiguration"
}

resource_schema "aws_autoscaling_lifecycle_hook" {
  cloudformation_type_name               = "AWS::AutoScaling::LifecycleHook"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_autoscaling_scaling_policy" {
  cloudformation_type_name = "AWS::AutoScaling::ScalingPolicy"
}

resource_schema "aws_autoscaling_scheduled_action" {
  cloudformation_type_name = "AWS::AutoScaling::ScheduledAction"
}

resource_schema "aws_autoscaling_warm_pool" {
  cloudformation_type_name               = "AWS::AutoScaling::WarmPool"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_b2bi_capability" {
  cloudformation_type_name = "AWS::B2BI::Capability"
}

resource_schema "aws_b2bi_partnership" {
  cloudformation_type_name = "AWS::B2BI::Partnership"
}

resource_schema "aws_b2bi_profile" {
  cloudformation_type_name = "AWS::B2BI::Profile"
}

resource_schema "aws_b2bi_transformer" {
  cloudformation_type_name = "AWS::B2BI::Transformer"
}

resource_schema "aws_bcmdataexports_export" {
  cloudformation_type_name = "AWS::BCMDataExports::Export"

  # Suppression Reason: Export/DataQuery/TableConfigurations is of unsupported type: key-value map of key-value map.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1509
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_backup_backup_plan" {
  cloudformation_type_name = "AWS::Backup::BackupPlan"
}

resource_schema "aws_backup_backup_selection" {
  cloudformation_type_name = "AWS::Backup::BackupSelection"
}

resource_schema "aws_backup_backup_vault" {
  cloudformation_type_name = "AWS::Backup::BackupVault"
}

resource_schema "aws_backup_framework" {
  cloudformation_type_name = "AWS::Backup::Framework"
}

resource_schema "aws_backup_report_plan" {
  cloudformation_type_name = "AWS::Backup::ReportPlan"
}

resource_schema "aws_backup_restore_testing_plan" {
  cloudformation_type_name = "AWS::Backup::RestoreTestingPlan"
}

resource_schema "aws_backup_restore_testing_selection" {
  cloudformation_type_name = "AWS::Backup::RestoreTestingSelection"
}

resource_schema "aws_backupgateway_hypervisor" {
  cloudformation_type_name = "AWS::BackupGateway::Hypervisor"
}

resource_schema "aws_batch_compute_environment" {
  cloudformation_type_name = "AWS::Batch::ComputeEnvironment"
}

resource_schema "aws_batch_job_definition" {
  cloudformation_type_name               = "AWS::Batch::JobDefinition"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_batch_job_queue" {
  cloudformation_type_name = "AWS::Batch::JobQueue"
}

resource_schema "aws_batch_scheduling_policy" {
  cloudformation_type_name = "AWS::Batch::SchedulingPolicy"
}

resource_schema "aws_bedrock_agent" {
  cloudformation_type_name = "AWS::Bedrock::Agent"
}

resource_schema "aws_bedrock_agent_alias" {
  cloudformation_type_name               = "AWS::Bedrock::AgentAlias"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_bedrock_data_source" {
  cloudformation_type_name               = "AWS::Bedrock::DataSource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_bedrock_knowledge_base" {
  cloudformation_type_name = "AWS::Bedrock::KnowledgeBase"
}

resource_schema "aws_billingconductor_billing_group" {
  cloudformation_type_name               = "AWS::BillingConductor::BillingGroup"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_billingconductor_custom_line_item" {
  cloudformation_type_name               = "AWS::BillingConductor::CustomLineItem"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_billingconductor_pricing_plan" {
  cloudformation_type_name               = "AWS::BillingConductor::PricingPlan"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_billingconductor_pricing_rule" {
  cloudformation_type_name               = "AWS::BillingConductor::PricingRule"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_budgets_budgets_action" {
  cloudformation_type_name = "AWS::Budgets::BudgetsAction"
}

resource_schema "aws_ce_anomaly_monitor" {
  cloudformation_type_name = "AWS::CE::AnomalyMonitor"
}

resource_schema "aws_ce_anomaly_subscription" {
  cloudformation_type_name = "AWS::CE::AnomalySubscription"
}

resource_schema "aws_ce_cost_category" {
  cloudformation_type_name               = "AWS::CE::CostCategory"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cassandra_keyspace" {
  cloudformation_type_name = "AWS::Cassandra::Keyspace"
}

resource_schema "aws_cassandra_table" {
  cloudformation_type_name = "AWS::Cassandra::Table"
}

resource_schema "aws_certificatemanager_account" {
  cloudformation_type_name               = "AWS::CertificateManager::Account"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_chatbot_microsoft_teams_channel_configuration" {
  cloudformation_type_name = "AWS::Chatbot::MicrosoftTeamsChannelConfiguration"
}

resource_schema "aws_chatbot_slack_channel_configuration" {
  cloudformation_type_name = "AWS::Chatbot::SlackChannelConfiguration"
}

resource_schema "aws_cleanrooms_analysis_template" {
  cloudformation_type_name = "AWS::CleanRooms::AnalysisTemplate"
}

resource_schema "aws_cleanrooms_collaboration" {
  cloudformation_type_name = "AWS::CleanRooms::Collaboration"
}

resource_schema "aws_cleanrooms_configured_table" {
  cloudformation_type_name = "AWS::CleanRooms::ConfiguredTable"
}

resource_schema "aws_cleanrooms_configured_table_association" {
  cloudformation_type_name               = "AWS::CleanRooms::ConfiguredTableAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cleanrooms_membership" {
  cloudformation_type_name = "AWS::CleanRooms::Membership"
}

resource_schema "aws_cleanroomsml_training_dataset" {
  cloudformation_type_name = "AWS::CleanRoomsML::TrainingDataset"
}

resource_schema "aws_cloudformation_hook_default_version" {
  cloudformation_type_name               = "AWS::CloudFormation::HookDefaultVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cloudformation_hook_type_config" {
  cloudformation_type_name               = "AWS::CloudFormation::HookTypeConfig"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cloudformation_hook_version" {
  cloudformation_type_name = "AWS::CloudFormation::HookVersion"
}

resource_schema "aws_cloudformation_module_default_version" {
  cloudformation_type_name = "AWS::CloudFormation::ModuleDefaultVersion"
}

resource_schema "aws_cloudformation_module_version" {
  cloudformation_type_name = "AWS::CloudFormation::ModuleVersion"
}

resource_schema "aws_cloudformation_public_type_version" {
  cloudformation_type_name = "AWS::CloudFormation::PublicTypeVersion"
}

resource_schema "aws_cloudformation_publisher" {
  cloudformation_type_name               = "AWS::CloudFormation::Publisher"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cloudformation_resource_default_version" {
  cloudformation_type_name               = "AWS::CloudFormation::ResourceDefaultVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cloudformation_resource_version" {
  cloudformation_type_name               = "AWS::CloudFormation::ResourceVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cloudformation_stack_set" {
  cloudformation_type_name = "AWS::CloudFormation::StackSet"
}

resource_schema "aws_cloudformation_type_activation" {
  cloudformation_type_name = "AWS::CloudFormation::TypeActivation"
}

resource_schema "aws_cloudfront_cache_policy" {
  cloudformation_type_name = "AWS::CloudFront::CachePolicy"
}

resource_schema "aws_cloudfront_cloudfront_origin_access_identity" {
  cloudformation_type_name = "AWS::CloudFront::CloudFrontOriginAccessIdentity"
}

resource_schema "aws_cloudfront_continuous_deployment_policy" {
  cloudformation_type_name = "AWS::CloudFront::ContinuousDeploymentPolicy"
}

resource_schema "aws_cloudfront_distribution" {
  cloudformation_type_name = "AWS::CloudFront::Distribution"
}

resource_schema "aws_cloudfront_function" {
  cloudformation_type_name = "AWS::CloudFront::Function"
}

resource_schema "aws_cloudfront_key_group" {
  cloudformation_type_name = "AWS::CloudFront::KeyGroup"
}

resource_schema "aws_cloudfront_key_value_store" {
  cloudformation_type_name = "AWS::CloudFront::KeyValueStore"
}

resource_schema "aws_cloudfront_monitoring_subscription" {
  cloudformation_type_name               = "AWS::CloudFront::MonitoringSubscription"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cloudfront_origin_access_control" {
  cloudformation_type_name = "AWS::CloudFront::OriginAccessControl"
}

resource_schema "aws_cloudfront_origin_request_policy" {
  cloudformation_type_name = "AWS::CloudFront::OriginRequestPolicy"
}

resource_schema "aws_cloudfront_public_key" {
  cloudformation_type_name = "AWS::CloudFront::PublicKey"
}

resource_schema "aws_cloudfront_realtime_log_config" {
  cloudformation_type_name = "AWS::CloudFront::RealtimeLogConfig"
}

resource_schema "aws_cloudfront_response_headers_policy" {
  cloudformation_type_name = "AWS::CloudFront::ResponseHeadersPolicy"
}

resource_schema "aws_cloudtrail_channel" {
  cloudformation_type_name = "AWS::CloudTrail::Channel"
}

resource_schema "aws_cloudtrail_event_data_store" {
  cloudformation_type_name = "AWS::CloudTrail::EventDataStore"
}

resource_schema "aws_cloudtrail_resource_policy" {
  cloudformation_type_name               = "AWS::CloudTrail::ResourcePolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cloudtrail_trail" {
  cloudformation_type_name = "AWS::CloudTrail::Trail"
}

resource_schema "aws_cloudwatch_alarm" {
  cloudformation_type_name = "AWS::CloudWatch::Alarm"
}

resource_schema "aws_cloudwatch_composite_alarm" {
  cloudformation_type_name = "AWS::CloudWatch::CompositeAlarm"
}

resource_schema "aws_cloudwatch_metric_stream" {
  cloudformation_type_name = "AWS::CloudWatch::MetricStream"
}

resource_schema "aws_codeartifact_domain" {
  cloudformation_type_name = "AWS::CodeArtifact::Domain"
}

resource_schema "aws_codeartifact_package_group" {
  cloudformation_type_name               = "AWS::CodeArtifact::PackageGroup"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_codeartifact_repository" {
  cloudformation_type_name = "AWS::CodeArtifact::Repository"
}

resource_schema "aws_codebuild_fleet" {
  cloudformation_type_name = "AWS::CodeBuild::Fleet"
}

resource_schema "aws_codeconnections_connection" {
  cloudformation_type_name = "AWS::CodeConnections::Connection"
}

resource_schema "aws_codedeploy_deployment_config" {
  cloudformation_type_name = "AWS::CodeDeploy::DeploymentConfig"
}

resource_schema "aws_codeguruprofiler_profiling_group" {
  cloudformation_type_name = "AWS::CodeGuruProfiler::ProfilingGroup"
}

resource_schema "aws_codegurureviewer_repository_association" {
  cloudformation_type_name               = "AWS::CodeGuruReviewer::RepositoryAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_codepipeline_custom_action_type" {
  cloudformation_type_name = "AWS::CodePipeline::CustomActionType"
}

resource_schema "aws_codestarconnections_connection" {
  cloudformation_type_name = "AWS::CodeStarConnections::Connection"
}

resource_schema "aws_codestarconnections_repository_link" {
  cloudformation_type_name = "AWS::CodeStarConnections::RepositoryLink"
}

resource_schema "aws_codestarconnections_sync_configuration" {
  cloudformation_type_name = "AWS::CodeStarConnections::SyncConfiguration"
}

resource_schema "aws_codestarnotifications_notification_rule" {
  cloudformation_type_name = "AWS::CodeStarNotifications::NotificationRule"
}

resource_schema "aws_cognito_identity_pool" {
  cloudformation_type_name = "AWS::Cognito::IdentityPool"
}

resource_schema "aws_cognito_identity_pool_role_attachment" {
  cloudformation_type_name = "AWS::Cognito::IdentityPoolRoleAttachment"

  # Suppression Reason: RoleMappings is of unsupported type: ""
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1505
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_cognito_identity_pool_principal_tag" {
  cloudformation_type_name               = "AWS::Cognito::IdentityPoolPrincipalTag"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cognito_log_delivery_configuration" {
  cloudformation_type_name               = "AWS::Cognito::LogDeliveryConfiguration"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cognito_user_pool" {
  cloudformation_type_name = "AWS::Cognito::UserPool"

  # Suppression Reason: SmsMessage overwrites SMSMessage for Terraform attribute sms_message.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1520
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_cognito_user_pool_client" {
  cloudformation_type_name               = "AWS::Cognito::UserPoolClient"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cognito_user_pool_user" {
  cloudformation_type_name               = "AWS::Cognito::UserPoolUser"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cognito_user_pool_risk_configuration_attachment" {
  cloudformation_type_name               = "AWS::Cognito::UserPoolRiskConfigurationAttachment"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_cognito_user_pool_user_to_group_attachment" {
  cloudformation_type_name               = "AWS::Cognito::UserPoolUserToGroupAttachment"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_comprehend_document_classifier" {
  cloudformation_type_name = "AWS::Comprehend::DocumentClassifier"
}

resource_schema "aws_comprehend_flywheel" {
  cloudformation_type_name = "AWS::Comprehend::Flywheel"
}

resource_schema "aws_config_aggregation_authorization" {
  cloudformation_type_name = "AWS::Config::AggregationAuthorization"
}

resource_schema "aws_config_config_rule" {
  cloudformation_type_name = "AWS::Config::ConfigRule"
}

resource_schema "aws_config_configuration_aggregator" {
  cloudformation_type_name = "AWS::Config::ConfigurationAggregator"
}

resource_schema "aws_config_conformance_pack" {
  cloudformation_type_name = "AWS::Config::ConformancePack"
}

resource_schema "aws_config_organization_conformance_pack" {
  cloudformation_type_name = "AWS::Config::OrganizationConformancePack"
}

resource_schema "aws_config_stored_query" {
  cloudformation_type_name = "AWS::Config::StoredQuery"
}

resource_schema "aws_connect_approved_origin" {
  cloudformation_type_name               = "AWS::Connect::ApprovedOrigin"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_contact_flow" {
  cloudformation_type_name               = "AWS::Connect::ContactFlow"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_contact_flow_module" {
  cloudformation_type_name               = "AWS::Connect::ContactFlowModule"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_evaluation_form" {
  cloudformation_type_name = "AWS::Connect::EvaluationForm"

  # Suppression Reason:
  # Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_connect_hours_of_operation" {
  cloudformation_type_name               = "AWS::Connect::HoursOfOperation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_instance" {
  cloudformation_type_name = "AWS::Connect::Instance"
}

resource_schema "aws_connect_instance_storage_config" {
  cloudformation_type_name               = "AWS::Connect::InstanceStorageConfig"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_integration_association" {
  cloudformation_type_name               = "AWS::Connect::IntegrationAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_phone_number" {
  cloudformation_type_name               = "AWS::Connect::PhoneNumber"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_predefined_attribute" {
  cloudformation_type_name               = "AWS::Connect::PredefinedAttribute"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_prompt" {
  cloudformation_type_name               = "AWS::Connect::Prompt"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_queue" {
  cloudformation_type_name               = "AWS::Connect::Queue"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_quick_connect" {
  cloudformation_type_name               = "AWS::Connect::QuickConnect"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_routing_profile" {
  cloudformation_type_name               = "AWS::Connect::RoutingProfile"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_rule" {
  cloudformation_type_name = "AWS::Connect::Rule"

  # Suppression Reason: Actions/AssignContactCategoryActions is of unsupported type: set of undefined schema.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1510
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_connect_security_key" {
  cloudformation_type_name               = "AWS::Connect::SecurityKey"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_security_profile" {
  cloudformation_type_name               = "AWS::Connect::SecurityProfile"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_task_template" {
  cloudformation_type_name               = "AWS::Connect::TaskTemplate"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_traffic_distribution_group" {
  cloudformation_type_name               = "AWS::Connect::TrafficDistributionGroup"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_user" {
  cloudformation_type_name               = "AWS::Connect::User"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_user_hierarchy_group" {
  cloudformation_type_name               = "AWS::Connect::UserHierarchyGroup"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_view" {
  cloudformation_type_name               = "AWS::Connect::View"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connect_view_version" {
  cloudformation_type_name               = "AWS::Connect::ViewVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_connectcampaigns_campaign" {
  cloudformation_type_name = "AWS::ConnectCampaigns::Campaign"
}

resource_schema "aws_controltower_enabled_baseline" {
  cloudformation_type_name = "AWS::ControlTower::EnabledBaseline"

  # Suppression Reason:
  # Properties with anyOf indicator https://github.com/hashicorp/terraform-provider-awscc/issues/98
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_controltower_enabled_control" {
  cloudformation_type_name               = "AWS::ControlTower::EnabledControl"
  suppress_plural_data_source_generation = true

  # Latest schema updates are suppressed.
  # git checkout internal/service/cloudformation/schemas/AWS_ControlTower_EnabledControl.json
  # Suppression Reason:  Parameters/Value is of unsupported type: .
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1527
}

resource_schema "aws_controltower_landing_zone" {
  cloudformation_type_name = "AWS::ControlTower::LandingZone"

  # Suppression Reason: Manifest is of unsupported type: .
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1519
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_cur_report_definition" {
  cloudformation_type_name = "AWS::CUR::ReportDefinition"
}

resource_schema "aws_customerprofiles_calculated_attribute_definition" {
  cloudformation_type_name = "AWS::CustomerProfiles::CalculatedAttributeDefinition"
}

resource_schema "aws_customerprofiles_domain" {
  cloudformation_type_name = "AWS::CustomerProfiles::Domain"

  # Latest schema updates are suppressed.
  # git checkout internal/service/cloudformation/schemas/AWS_CustomerProfiles_Domain.json
  # Suppression Reason:  Matching/AutoMerging/Consolidation/MatchingAttributesList is of unsupported type: list of array.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1528
}

resource_schema "aws_customerprofiles_event_stream" {
  cloudformation_type_name               = "AWS::CustomerProfiles::EventStream"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_customerprofiles_integration" {
  cloudformation_type_name               = "AWS::CustomerProfiles::Integration"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_customerprofiles_object_type" {
  cloudformation_type_name               = "AWS::CustomerProfiles::ObjectType"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_dms_data_provider" {
  cloudformation_type_name = "AWS::DMS::DataProvider"
}

resource_schema "aws_dms_instance_profile" {
  cloudformation_type_name = "AWS::DMS::InstanceProfile"
}

resource_schema "aws_dms_migration_project" {
  cloudformation_type_name = "AWS::DMS::MigrationProject"
}

resource_schema "aws_dms_replication_config" {
  cloudformation_type_name = "AWS::DMS::ReplicationConfig"
}

resource_schema "aws_databrew_dataset" {
  cloudformation_type_name = "AWS::DataBrew::Dataset"
}

resource_schema "aws_databrew_job" {
  cloudformation_type_name = "AWS::DataBrew::Job"
}

resource_schema "aws_databrew_project" {
  cloudformation_type_name = "AWS::DataBrew::Project"
}

resource_schema "aws_databrew_recipe" {
  cloudformation_type_name = "AWS::DataBrew::Recipe"

  # Suppression Reason:
  # Properties with anyOf indicator https://github.com/hashicorp/terraform-provider-awscc/issues/98
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_databrew_ruleset" {
  cloudformation_type_name = "AWS::DataBrew::Ruleset"
}

resource_schema "aws_databrew_schedule" {
  cloudformation_type_name = "AWS::DataBrew::Schedule"
}

resource_schema "aws_datapipeline_pipeline" {
  cloudformation_type_name = "AWS::DataPipeline::Pipeline"
}

resource_schema "aws_datasync_agent" {
  cloudformation_type_name = "AWS::DataSync::Agent"
}

resource_schema "aws_datasync_location_azure_blob" {
  cloudformation_type_name = "AWS::DataSync::LocationAzureBlob"
}

resource_schema "aws_datasync_location_efs" {
  cloudformation_type_name = "AWS::DataSync::LocationEFS"
}

resource_schema "aws_datasync_location_fsx_lustre" {
  cloudformation_type_name = "AWS::DataSync::LocationFSxLustre"
}

resource_schema "aws_datasync_location_fsx_ontap" {
  cloudformation_type_name = "AWS::DataSync::LocationFSxONTAP"
}

resource_schema "aws_datasync_location_fsx_open_zfs" {
  cloudformation_type_name = "AWS::DataSync::LocationFSxOpenZFS"
}

resource_schema "aws_datasync_location_fsx_windows" {
  cloudformation_type_name = "AWS::DataSync::LocationFSxWindows"
}

resource_schema "aws_datasync_location_hdfs" {
  cloudformation_type_name = "AWS::DataSync::LocationHDFS"
}

resource_schema "aws_datasync_location_nfs" {
  cloudformation_type_name = "AWS::DataSync::LocationNFS"
}

resource_schema "aws_datasync_location_object_storage" {
  cloudformation_type_name = "AWS::DataSync::LocationObjectStorage"
}

resource_schema "aws_datasync_location_s3" {
  cloudformation_type_name = "AWS::DataSync::LocationS3"
}

resource_schema "aws_datasync_location_smb" {
  cloudformation_type_name = "AWS::DataSync::LocationSMB"
}

resource_schema "aws_datasync_storage_system" {
  cloudformation_type_name = "AWS::DataSync::StorageSystem"
}

resource_schema "aws_datasync_task" {
  cloudformation_type_name = "AWS::DataSync::Task"
}

resource_schema "aws_datazone_data_source" {
  cloudformation_type_name               = "AWS::DataZone::DataSource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_datazone_domain" {
  cloudformation_type_name = "AWS::DataZone::Domain"
}

resource_schema "aws_datazone_environment" {
  cloudformation_type_name = "AWS::DataZone::Environment"
}

resource_schema "aws_datazone_environment_blueprint_configuration" {
  cloudformation_type_name               = "AWS::DataZone::EnvironmentBlueprintConfiguration"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_datazone_environment_profile" {
  cloudformation_type_name               = "AWS::DataZone::EnvironmentProfile"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_datazone_project" {
  cloudformation_type_name               = "AWS::DataZone::Project"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_datazone_subscription_target" {
  cloudformation_type_name = "AWS::DataZone::SubscriptionTarget"
}

resource_schema "aws_deadline_farm" {
  cloudformation_type_name = "AWS::Deadline::Farm"
}

resource_schema "aws_deadline_fleet" {
  cloudformation_type_name               = "AWS::Deadline::Fleet"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_deadline_license_endpoint" {
  cloudformation_type_name = "AWS::Deadline::LicenseEndpoint"
}

resource_schema "aws_deadline_metered_product" {
  cloudformation_type_name = "AWS::Deadline::MeteredProduct"
}

resource_schema "aws_deadline_queue" {
  cloudformation_type_name               = "AWS::Deadline::Queue"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_deadline_queue_environment" {
  cloudformation_type_name               = "AWS::Deadline::QueueEnvironment"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_deadline_queue_fleet_association" {
  cloudformation_type_name               = "AWS::Deadline::QueueFleetAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_deadline_storage_profile" {
  cloudformation_type_name               = "AWS::Deadline::StorageProfile"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_detective_graph" {
  cloudformation_type_name = "AWS::Detective::Graph"
}

resource_schema "aws_detective_member_invitation" {
  cloudformation_type_name = "AWS::Detective::MemberInvitation"
}

resource_schema "aws_detective_organization_admin" {
  cloudformation_type_name               = "AWS::Detective::OrganizationAdmin"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_devopsguru_log_anomaly_detection_integration" {
  cloudformation_type_name = "AWS::DevOpsGuru::LogAnomalyDetectionIntegration"
}

resource_schema "aws_devopsguru_notification_channel" {
  cloudformation_type_name = "AWS::DevOpsGuru::NotificationChannel"
}

resource_schema "aws_devopsguru_resource_collection" {
  cloudformation_type_name = "AWS::DevOpsGuru::ResourceCollection"
}

resource_schema "aws_directoryservice_simple_ad" {
  cloudformation_type_name = "AWS::DirectoryService::SimpleAD"
}

resource_schema "aws_docdbelastic_cluster" {
  cloudformation_type_name = "AWS::DocDBElastic::Cluster"
}

resource_schema "aws_dynamodb_global_table" {
  cloudformation_type_name = "AWS::DynamoDB::GlobalTable"
}

resource_schema "aws_dynamodb_table" {
  cloudformation_type_name = "AWS::DynamoDB::Table"
}

resource_schema "aws_ec2_capacity_reservation" {
  cloudformation_type_name = "AWS::EC2::CapacityReservation"
}

resource_schema "aws_ec2_capacity_reservation_fleet" {
  cloudformation_type_name = "AWS::EC2::CapacityReservationFleet"
}

resource_schema "aws_ec2_carrier_gateway" {
  cloudformation_type_name = "AWS::EC2::CarrierGateway"
}

resource_schema "aws_ec2_customer_gateway" {
  cloudformation_type_name = "AWS::EC2::CustomerGateway"
}

resource_schema "aws_ec2_dhcp_options" {
  cloudformation_type_name = "AWS::EC2::DHCPOptions"
}

resource_schema "aws_ec2_ec2_fleet" {
  cloudformation_type_name = "AWS::EC2::EC2Fleet"
}

resource_schema "aws_ec2_eip" {
  cloudformation_type_name = "AWS::EC2::EIP"
}

resource_schema "aws_ec2_eip_association" {
  cloudformation_type_name = "AWS::EC2::EIPAssociation"
}

resource_schema "aws_ec2_egress_only_internet_gateway" {
  cloudformation_type_name = "AWS::EC2::EgressOnlyInternetGateway"
}

resource_schema "aws_ec2_enclave_certificate_iam_role_association" {
  cloudformation_type_name               = "AWS::EC2::EnclaveCertificateIamRoleAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_flow_log" {
  cloudformation_type_name = "AWS::EC2::FlowLog"
}

resource_schema "aws_ec2_gateway_route_table_association" {
  cloudformation_type_name               = "AWS::EC2::GatewayRouteTableAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_host" {
  cloudformation_type_name = "AWS::EC2::Host"
}

resource_schema "aws_ec2_instance" {
  cloudformation_type_name               = "AWS::EC2::Instance"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_internet_gateway" {
  cloudformation_type_name = "AWS::EC2::InternetGateway"
}

resource_schema "aws_ec2_ipam" {
  cloudformation_type_name = "AWS::EC2::IPAM"
}

resource_schema "aws_ec2_ipam_allocation" {
  cloudformation_type_name               = "AWS::EC2::IPAMAllocation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_ipam_pool" {
  cloudformation_type_name = "AWS::EC2::IPAMPool"
}

resource_schema "aws_ec2_ipam_pool_cidr" {
  cloudformation_type_name               = "AWS::EC2::IPAMPoolCidr"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_ipam_resource_discovery" {
  cloudformation_type_name = "AWS::EC2::IPAMResourceDiscovery"
}

resource_schema "aws_ec2_ipam_resource_discovery_association" {
  cloudformation_type_name = "AWS::EC2::IPAMResourceDiscoveryAssociation"
}

resource_schema "aws_ec2_ipam_scope" {
  cloudformation_type_name = "AWS::EC2::IPAMScope"
}

resource_schema "aws_ec2_launch_template" {
  cloudformation_type_name = "AWS::EC2::LaunchTemplate"
}

resource_schema "aws_ec2_local_gateway_route" {
  cloudformation_type_name = "AWS::EC2::LocalGatewayRoute"
}

resource_schema "aws_ec2_local_gateway_route_table" {
  cloudformation_type_name = "AWS::EC2::LocalGatewayRouteTable"
}

resource_schema "aws_ec2_local_gateway_route_table_virtual_interface_group_association" {
  cloudformation_type_name = "AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation"
}

resource_schema "aws_ec2_local_gateway_route_table_vpc_association" {
  cloudformation_type_name = "AWS::EC2::LocalGatewayRouteTableVPCAssociation"
}

resource_schema "aws_ec2_nat_gateway" {
  cloudformation_type_name = "AWS::EC2::NatGateway"
}

resource_schema "aws_ec2_network_acl" {
  cloudformation_type_name = "AWS::EC2::NetworkAcl"
}

resource_schema "aws_ec2_network_insights_access_scope" {
  cloudformation_type_name = "AWS::EC2::NetworkInsightsAccessScope"
}

resource_schema "aws_ec2_network_insights_access_scope_analysis" {
  cloudformation_type_name = "AWS::EC2::NetworkInsightsAccessScopeAnalysis"
}

resource_schema "aws_ec2_network_insights_analysis" {
  cloudformation_type_name = "AWS::EC2::NetworkInsightsAnalysis"

  # Latest schema updates are suppressed.
  # git checkout internal/service/cloudformation/schemas/AWS_EC2_NetworkInsightsAnalysis.json
  # Suppression Reason:  DestinationCidr overwrites destinationCidr for Terraform attribute destination_cidr.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1529
}

resource_schema "aws_ec2_network_insights_path" {
  cloudformation_type_name = "AWS::EC2::NetworkInsightsPath"
}

resource_schema "aws_ec2_network_interface" {
  cloudformation_type_name = "AWS::EC2::NetworkInterface"
}

resource_schema "aws_ec2_network_interface_attachment" {
  cloudformation_type_name = "AWS::EC2::NetworkInterfaceAttachment"
}

resource_schema "aws_ec2_network_performance_metric_subscription" {
  cloudformation_type_name = "AWS::EC2::NetworkPerformanceMetricSubscription"
}

resource_schema "aws_ec2_placement_group" {
  cloudformation_type_name = "AWS::EC2::PlacementGroup"
}

resource_schema "aws_ec2_prefix_list" {
  cloudformation_type_name = "AWS::EC2::PrefixList"
}

resource_schema "aws_ec2_route" {
  cloudformation_type_name               = "AWS::EC2::Route"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_route_table" {
  cloudformation_type_name = "AWS::EC2::RouteTable"
}

resource_schema "aws_ec2_security_group" {
  cloudformation_type_name               = "AWS::EC2::SecurityGroup"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_security_group_egress" {
  cloudformation_type_name = "AWS::EC2::SecurityGroupEgress"
}

resource_schema "aws_ec2_security_group_ingress" {
  cloudformation_type_name = "AWS::EC2::SecurityGroupIngress"
}

resource_schema "aws_ec2_snapshot_block_public_access" {
  cloudformation_type_name = "AWS::EC2::SnapshotBlockPublicAccess"
}

resource_schema "aws_ec2_spot_fleet" {
  cloudformation_type_name = "AWS::EC2::SpotFleet"
}

resource_schema "aws_ec2_subnet" {
  cloudformation_type_name               = "AWS::EC2::Subnet"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_subnet_cidr_block" {
  cloudformation_type_name = "AWS::EC2::SubnetCidrBlock"
}

resource_schema "aws_ec2_subnet_network_acl_association" {
  cloudformation_type_name = "AWS::EC2::SubnetNetworkAclAssociation"
}

resource_schema "aws_ec2_subnet_route_table_association" {
  cloudformation_type_name = "AWS::EC2::SubnetRouteTableAssociation"
}

resource_schema "aws_ec2_transit_gateway" {
  cloudformation_type_name = "AWS::EC2::TransitGateway"
}

resource_schema "aws_ec2_transit_gateway_attachment" {
  cloudformation_type_name = "AWS::EC2::TransitGatewayAttachment"
}

resource_schema "aws_ec2_transit_gateway_connect" {
  cloudformation_type_name = "AWS::EC2::TransitGatewayConnect"
}

resource_schema "aws_ec2_transit_gateway_multicast_domain" {
  cloudformation_type_name = "AWS::EC2::TransitGatewayMulticastDomain"
}

resource_schema "aws_ec2_transit_gateway_multicast_domain_association" {
  cloudformation_type_name               = "AWS::EC2::TransitGatewayMulticastDomainAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_transit_gateway_multicast_group_member" {
  cloudformation_type_name               = "AWS::EC2::TransitGatewayMulticastGroupMember"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_transit_gateway_multicast_group_source" {
  cloudformation_type_name               = "AWS::EC2::TransitGatewayMulticastGroupSource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_transit_gateway_peering_attachment" {
  cloudformation_type_name = "AWS::EC2::TransitGatewayPeeringAttachment"
}

resource_schema "aws_ec2_transit_gateway_route_table" {
  cloudformation_type_name = "AWS::EC2::TransitGatewayRouteTable"
}

resource_schema "aws_ec2_transit_gateway_route_table_association" {
  cloudformation_type_name               = "AWS::EC2::TransitGatewayRouteTableAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_vpc_endpoint_connection_notification" {
  cloudformation_type_name = "AWS::EC2::VPCEndpointConnectionNotification"
}

resource_schema "aws_ec2_transit_gateway_vpc_attachment" {
  cloudformation_type_name = "AWS::EC2::TransitGatewayVpcAttachment"
}

resource_schema "aws_ec2_vpc" {
  cloudformation_type_name = "AWS::EC2::VPC"
}

resource_schema "aws_ec2_vpc_cidr_block" {
  cloudformation_type_name               = "AWS::EC2::VPCCidrBlock"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ec2_vpcdhcp_options_association" {
  cloudformation_type_name = "AWS::EC2::VPCDHCPOptionsAssociation"
}

resource_schema "aws_ec2_vpc_endpoint" {
  cloudformation_type_name = "AWS::EC2::VPCEndpoint"
}

resource_schema "aws_ec2_vpc_endpoint_service" {
  cloudformation_type_name = "AWS::EC2::VPCEndpointService"
}

resource_schema "aws_ec2_vpc_endpoint_service_permissions" {
  cloudformation_type_name = "AWS::EC2::VPCEndpointServicePermissions"
}

resource_schema "aws_ec2_vpc_gateway_attachment" {
  cloudformation_type_name = "AWS::EC2::VPCGatewayAttachment"
}

resource_schema "aws_ec2_vpc_peering_connection" {
  cloudformation_type_name = "AWS::EC2::VPCPeeringConnection"
}

resource_schema "aws_ec2_vpn_connection" {
  cloudformation_type_name = "AWS::EC2::VPNConnection"
}

resource_schema "aws_ec2_vpn_connection_route" {
  cloudformation_type_name = "AWS::EC2::VPNConnectionRoute"
}

resource_schema "aws_ec2_vpn_gateway" {
  cloudformation_type_name = "AWS::EC2::VPNGateway"
}

resource_schema "aws_ec2_verified_access_endpoint" {
  cloudformation_type_name = "AWS::EC2::VerifiedAccessEndpoint"
}

resource_schema "aws_ec2_verified_access_group" {
  cloudformation_type_name = "AWS::EC2::VerifiedAccessGroup"
}

resource_schema "aws_ec2_verified_access_instance" {
  cloudformation_type_name = "AWS::EC2::VerifiedAccessInstance"
}

resource_schema "aws_ec2_verified_access_trust_provider" {
  cloudformation_type_name = "AWS::EC2::VerifiedAccessTrustProvider"
}

resource_schema "aws_ec2_volume" {
  cloudformation_type_name = "AWS::EC2::Volume"
}

resource_schema "aws_ec2_volume_attachment" {
  cloudformation_type_name = "AWS::EC2::VolumeAttachment"
}

resource_schema "aws_ecr_pull_through_cache_rule" {
  cloudformation_type_name = "AWS::ECR::PullThroughCacheRule"
}

resource_schema "aws_ecr_registry_policy" {
  cloudformation_type_name = "AWS::ECR::RegistryPolicy"
}

resource_schema "aws_ecr_replication_configuration" {
  cloudformation_type_name = "AWS::ECR::ReplicationConfiguration"
}

resource_schema "aws_ecr_public_repository" {
  cloudformation_type_name = "AWS::ECR::PublicRepository"
}

resource_schema "aws_ecr_repository" {
  cloudformation_type_name = "AWS::ECR::Repository"
}

resource_schema "aws_ecs_capacity_provider" {
  cloudformation_type_name = "AWS::ECS::CapacityProvider"
}

resource_schema "aws_ecs_cluster" {
  cloudformation_type_name = "AWS::ECS::Cluster"
}

resource_schema "aws_ecs_cluster_capacity_provider_associations" {
  cloudformation_type_name = "AWS::ECS::ClusterCapacityProviderAssociations"
}

resource_schema "aws_ecs_primary_task_set" {
  cloudformation_type_name               = "AWS::ECS::PrimaryTaskSet"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ecs_service" {
  cloudformation_type_name = "AWS::ECS::Service"
}

resource_schema "aws_ecs_task_definition" {
  cloudformation_type_name = "AWS::ECS::TaskDefinition"
}

resource_schema "aws_ecs_task_set" {
  cloudformation_type_name               = "AWS::ECS::TaskSet"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_efs_access_point" {
  cloudformation_type_name = "AWS::EFS::AccessPoint"
}

resource_schema "aws_efs_file_system" {
  cloudformation_type_name = "AWS::EFS::FileSystem"
}

resource_schema "aws_efs_mount_target" {
  cloudformation_type_name               = "AWS::EFS::MountTarget"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_eks_access_entry" {
  cloudformation_type_name               = "AWS::EKS::AccessEntry"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_eks_addon" {
  cloudformation_type_name               = "AWS::EKS::Addon"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_eks_cluster" {
  cloudformation_type_name = "AWS::EKS::Cluster"
}

resource_schema "aws_eks_fargate_profile" {
  cloudformation_type_name               = "AWS::EKS::FargateProfile"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_eks_identity_provider_config" {
  cloudformation_type_name               = "AWS::EKS::IdentityProviderConfig"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_eks_nodegroup" {
  cloudformation_type_name               = "AWS::EKS::Nodegroup"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_eks_pod_identity_association" {
  cloudformation_type_name               = "AWS::EKS::PodIdentityAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_emr_security_configuration" {
  cloudformation_type_name = "AWS::EMR::SecurityConfiguration"
}

resource_schema "aws_emr_studio" {
  cloudformation_type_name = "AWS::EMR::Studio"
}

resource_schema "aws_emr_studio_session_mapping" {
  cloudformation_type_name = "AWS::EMR::StudioSessionMapping"
}

resource_schema "aws_emr_wal_workspace" {
  cloudformation_type_name = "AWS::EMR::WALWorkspace"
}

resource_schema "aws_emrcontainers_virtual_cluster" {
  cloudformation_type_name = "AWS::EMRContainers::VirtualCluster"
}

resource_schema "aws_elasticache_global_replication_group" {
  cloudformation_type_name = "AWS::ElastiCache::GlobalReplicationGroup"
}

resource_schema "aws_elasticache_serverless_cache" {
  cloudformation_type_name = "AWS::ElastiCache::ServerlessCache"
}

resource_schema "aws_elasticache_subnet_group" {
  cloudformation_type_name = "AWS::ElastiCache::SubnetGroup"
}

resource_schema "aws_elasticache_user" {
  cloudformation_type_name = "AWS::ElastiCache::User"
}

resource_schema "aws_elasticache_user_group" {
  cloudformation_type_name = "AWS::ElastiCache::UserGroup"
}

resource_schema "aws_elasticbeanstalk_application" {
  cloudformation_type_name = "AWS::ElasticBeanstalk::Application"
}

resource_schema "aws_elasticbeanstalk_application_version" {
  cloudformation_type_name = "AWS::ElasticBeanstalk::ApplicationVersion"
}

resource_schema "aws_elasticbeanstalk_configuration_template" {
  cloudformation_type_name = "AWS::ElasticBeanstalk::ConfigurationTemplate"
}

resource_schema "aws_elasticbeanstalk_environment" {
  cloudformation_type_name = "AWS::ElasticBeanstalk::Environment"
}

resource_schema "aws_elasticloadbalancingv2_listener" {
  cloudformation_type_name               = "AWS::ElasticLoadBalancingV2::Listener"
  suppress_plural_data_source_generation = true

  # Suppression Reason: error creating write-only attribute path (/properties/DefaultActions/*/AuthenticateOidcConfig/ClientSecret): invalid property path segment: "*"
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1521
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_elasticloadbalancingv2_listener_rule" {
  cloudformation_type_name               = "AWS::ElasticLoadBalancingV2::ListenerRule"
  suppress_plural_data_source_generation = true

  # Suppression Reason:  error creating write-only attribute path (/properties/Actions/*/AuthenticateOidcConfig/ClientSecret): invalid property path segment: "*"
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1521
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_elasticloadbalancingv2_load_balancer" {
  cloudformation_type_name = "AWS::ElasticLoadBalancingV2::LoadBalancer"
}

resource_schema "aws_elasticloadbalancingv2_target_group" {
  cloudformation_type_name = "AWS::ElasticLoadBalancingV2::TargetGroup"
}

resource_schema "aws_elasticloadbalancingv2_trust_store" {
  cloudformation_type_name               = "AWS::ElasticLoadBalancingV2::TrustStore"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_elasticloadbalancingv2_trust_store_revocation" {
  cloudformation_type_name               = "AWS::ElasticLoadBalancingV2::TrustStoreRevocation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_emrserverless_application" {
  cloudformation_type_name = "AWS::EMRServerless::Application"

  # Latest schema updates are suppressed.
  # git checkout internal/service/cloudformation/schemas/AWS_EC2_NetworkInsightsAnalysis.json
  # Suppression Reason:
  # Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95
}

resource_schema "aws_entityresolution_id_mapping_workflow" {
  cloudformation_type_name = "AWS::EntityResolution::IdMappingWorkflow"
}

resource_schema "aws_entityresolution_id_namespace" {
  cloudformation_type_name = "AWS::EntityResolution::IdNamespace"
}

resource_schema "aws_entityresolution_policy_statement" {
  cloudformation_type_name               = "AWS::EntityResolution::PolicyStatement"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_eventschemas_registry_policy" {
  cloudformation_type_name               = "AWS::EventSchemas::RegistryPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_eventschemas_schema" {
  cloudformation_type_name               = "AWS::EventSchemas::Schema"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_events_api_destination" {
  cloudformation_type_name = "AWS::Events::ApiDestination"
}

resource_schema "aws_events_archive" {
  cloudformation_type_name = "AWS::Events::Archive"
}

resource_schema "aws_events_connection" {
  cloudformation_type_name = "AWS::Events::Connection"

  # Suppression Reason: error creating write-only attribute path (/definitions/BasicAuthParameters/Password): expected "properties" for the second property path segment, got: "definitions"
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1521
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_events_endpoint" {
  cloudformation_type_name = "AWS::Events::Endpoint"
}

resource_schema "aws_events_rule" {
  cloudformation_type_name = "AWS::Events::Rule"
}

resource_schema "aws_events_event_bus" {
  cloudformation_type_name = "AWS::Events::EventBus"
}

resource_schema "aws_eventschemas_discoverer" {
  cloudformation_type_name = "AWS::EventSchemas::Discoverer"
}

resource_schema "aws_eventschemas_registry" {
  cloudformation_type_name = "AWS::EventSchemas::Registry"
}

resource_schema "aws_evidently_experiment" {
  cloudformation_type_name               = "AWS::Evidently::Experiment"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_evidently_feature" {
  cloudformation_type_name               = "AWS::Evidently::Feature"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_evidently_launch" {
  cloudformation_type_name               = "AWS::Evidently::Launch"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_evidently_project" {
  cloudformation_type_name               = "AWS::Evidently::Project"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_evidently_segment" {
  cloudformation_type_name               = "AWS::Evidently::Segment"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_fis_experiment_template" {
  cloudformation_type_name = "AWS::FIS::ExperimentTemplate"
}

resource_schema "aws_fis_target_account_configuration" {
  cloudformation_type_name               = "AWS::FIS::TargetAccountConfiguration"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_forecast_dataset" {
  cloudformation_type_name = "AWS::Forecast::Dataset"
}

resource_schema "aws_forecast_dataset_group" {
  cloudformation_type_name = "AWS::Forecast::DatasetGroup"
}

resource_schema "aws_fms_notification_channel" {
  cloudformation_type_name               = "AWS::FMS::NotificationChannel"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_fms_policy" {
  cloudformation_type_name               = "AWS::FMS::Policy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_fms_resource_set" {
  cloudformation_type_name               = "AWS::FMS::ResourceSet"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_fsx_data_repository_association" {
  cloudformation_type_name = "AWS::FSx::DataRepositoryAssociation"
}

resource_schema "aws_finspace_environment" {
  cloudformation_type_name = "AWS::FinSpace::Environment"
}

resource_schema "aws_frauddetector_detector" {
  cloudformation_type_name = "AWS::FraudDetector::Detector"
}

resource_schema "aws_frauddetector_entity_type" {
  cloudformation_type_name = "AWS::FraudDetector::EntityType"
}

resource_schema "aws_frauddetector_event_type" {
  cloudformation_type_name = "AWS::FraudDetector::EventType"
}

resource_schema "aws_frauddetector_label" {
  cloudformation_type_name = "AWS::FraudDetector::Label"
}

resource_schema "aws_frauddetector_list" {
  cloudformation_type_name = "AWS::FraudDetector::List"
}

resource_schema "aws_frauddetector_outcome" {
  cloudformation_type_name = "AWS::FraudDetector::Outcome"
}

resource_schema "aws_frauddetector_variable" {
  cloudformation_type_name = "AWS::FraudDetector::Variable"
}

resource_schema "aws_gamelift_alias" {
  cloudformation_type_name = "AWS::GameLift::Alias"
}

resource_schema "aws_gamelift_build" {
  cloudformation_type_name = "AWS::GameLift::Build"
}

resource_schema "aws_gamelift_fleet" {
  cloudformation_type_name = "AWS::GameLift::Fleet"
}

resource_schema "aws_gamelift_game_server_group" {
  cloudformation_type_name = "AWS::GameLift::GameServerGroup"
}

resource_schema "aws_gamelift_game_session_queue" {
  cloudformation_type_name = "AWS::GameLift::GameSessionQueue"
}

resource_schema "aws_gamelift_location" {
  cloudformation_type_name = "AWS::GameLift::Location"
}

resource_schema "aws_gamelift_matchmaking_configuration" {
  cloudformation_type_name = "AWS::GameLift::MatchmakingConfiguration"
}

resource_schema "aws_gamelift_matchmaking_rule_set" {
  cloudformation_type_name = "AWS::GameLift::MatchmakingRuleSet"
}

resource_schema "aws_gamelift_script" {
  cloudformation_type_name = "AWS::GameLift::Script"
}

resource_schema "aws_globalaccelerator_accelerator" {
  cloudformation_type_name = "AWS::GlobalAccelerator::Accelerator"
}

resource_schema "aws_globalaccelerator_cross_account_attachment" {
  cloudformation_type_name = "AWS::GlobalAccelerator::CrossAccountAttachment"
}

resource_schema "aws_globalaccelerator_endpoint_group" {
  cloudformation_type_name               = "AWS::GlobalAccelerator::EndpointGroup"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_globalaccelerator_listener" {
  cloudformation_type_name               = "AWS::GlobalAccelerator::Listener"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_glue_registry" {
  cloudformation_type_name = "AWS::Glue::Registry"
}

resource_schema "aws_glue_schema" {
  cloudformation_type_name = "AWS::Glue::Schema"
}

resource_schema "aws_glue_schema_version" {
  cloudformation_type_name               = "AWS::Glue::SchemaVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_glue_schema_version_metadata" {
  cloudformation_type_name               = "AWS::Glue::SchemaVersionMetadata"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_grafana_workspace" {
  cloudformation_type_name = "AWS::Grafana::Workspace"
}

resource_schema "aws_greengrassv2_component_version" {
  cloudformation_type_name = "AWS::GreengrassV2::ComponentVersion"
}

resource_schema "aws_greengrassv2_deployment" {
  cloudformation_type_name = "AWS::GreengrassV2::Deployment"
}

resource_schema "aws_groundstation_config" {
  cloudformation_type_name = "AWS::GroundStation::Config"
}

resource_schema "aws_groundstation_dataflow_endpoint_group" {
  cloudformation_type_name = "AWS::GroundStation::DataflowEndpointGroup"
}

resource_schema "aws_groundstation_mission_profile" {
  cloudformation_type_name = "AWS::GroundStation::MissionProfile"
}

resource_schema "aws_guardduty_detector" {
  cloudformation_type_name = "AWS::GuardDuty::Detector"
}

resource_schema "aws_guardduty_filter" {
  cloudformation_type_name               = "AWS::GuardDuty::Filter"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_guardduty_ip_set" {
  cloudformation_type_name               = "AWS::GuardDuty::IPSet"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_guardduty_master" {
  cloudformation_type_name               = "AWS::GuardDuty::Master"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_guardduty_member" {
  cloudformation_type_name               = "AWS::GuardDuty::Member"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_guardduty_threat_intel_set" {
  cloudformation_type_name               = "AWS::GuardDuty::ThreatIntelSet"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_healthimaging_datastore" {
  cloudformation_type_name = "AWS::HealthImaging::Datastore"
}

resource_schema "aws_healthlake_fhir_datastore" {
  cloudformation_type_name = "AWS::HealthLake::FHIRDatastore"
}

resource_schema "aws_iam_group" {
  cloudformation_type_name = "AWS::IAM::Group"
}

resource_schema "aws_iam_group_policy" {
  cloudformation_type_name               = "AWS::IAM::GroupPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_iam_managed_policy" {
  cloudformation_type_name = "AWS::IAM::ManagedPolicy"
}

resource_schema "aws_iam_oidc_provider" {
  cloudformation_type_name = "AWS::IAM::OIDCProvider"
}

resource_schema "aws_iam_role" {
  cloudformation_type_name = "AWS::IAM::Role"
}

resource_schema "aws_iam_role_policy" {
  cloudformation_type_name               = "AWS::IAM::RolePolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_iam_saml_provider" {
  cloudformation_type_name = "AWS::IAM::SAMLProvider"
}

resource_schema "aws_iam_service_linked_role" {
  cloudformation_type_name               = "AWS::IAM::ServiceLinkedRole"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_iam_server_certificate" {
  cloudformation_type_name = "AWS::IAM::ServerCertificate"
}

resource_schema "aws_iam_user" {
  cloudformation_type_name = "AWS::IAM::User"
}

resource_schema "aws_iam_user_policy" {
  cloudformation_type_name               = "AWS::IAM::UserPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_iam_virtual_mfa_device" {
  cloudformation_type_name = "AWS::IAM::VirtualMFADevice"
}

resource_schema "aws_ivs_channel" {
  cloudformation_type_name = "AWS::IVS::Channel"
}

resource_schema "aws_ivs_encoder_configuration" {
  cloudformation_type_name = "AWS::IVS::EncoderConfiguration"
}

resource_schema "aws_ivs_playback_key_pair" {
  cloudformation_type_name = "AWS::IVS::PlaybackKeyPair"
}

resource_schema "aws_ivs_playback_restriction_policy" {
  cloudformation_type_name = "AWS::IVS::PlaybackRestrictionPolicy"
}

resource_schema "aws_ivs_recording_configuration" {
  cloudformation_type_name = "AWS::IVS::RecordingConfiguration"
}

resource_schema "aws_ivs_stage" {
  cloudformation_type_name = "AWS::IVS::Stage"
}

resource_schema "aws_ivs_storage_configuration" {
  cloudformation_type_name = "AWS::IVS::StorageConfiguration"
}

resource_schema "aws_ivs_stream_key" {
  cloudformation_type_name = "AWS::IVS::StreamKey"
}

resource_schema "aws_ivschat_logging_configuration" {
  cloudformation_type_name = "AWS::IVSChat::LoggingConfiguration"
}

resource_schema "aws_ivschat_room" {
  cloudformation_type_name = "AWS::IVSChat::Room"
}

resource_schema "aws_identitystore_group" {
  cloudformation_type_name               = "AWS::IdentityStore::Group"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_identitystore_group_membership" {
  cloudformation_type_name               = "AWS::IdentityStore::GroupMembership"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_inspector_assessment_target" {
  cloudformation_type_name = "AWS::Inspector::AssessmentTarget"
}

resource_schema "aws_inspector_assessment_template" {
  cloudformation_type_name = "AWS::Inspector::AssessmentTemplate"
}

resource_schema "aws_inspector_resource_group" {
  cloudformation_type_name               = "AWS::Inspector::ResourceGroup"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_imagebuilder_component" {
  cloudformation_type_name               = "AWS::ImageBuilder::Component"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_imagebuilder_container_recipe" {
  cloudformation_type_name = "AWS::ImageBuilder::ContainerRecipe"
}

resource_schema "aws_imagebuilder_distribution_configuration" {
  cloudformation_type_name = "AWS::ImageBuilder::DistributionConfiguration"
}

resource_schema "aws_imagebuilder_image" {
  cloudformation_type_name = "AWS::ImageBuilder::Image"
}

resource_schema "aws_imagebuilder_image_pipeline" {
  cloudformation_type_name = "AWS::ImageBuilder::ImagePipeline"
}

resource_schema "aws_imagebuilder_image_recipe" {
  cloudformation_type_name = "AWS::ImageBuilder::ImageRecipe"
}

resource_schema "aws_imagebuilder_infrastructure_configuration" {
  cloudformation_type_name = "AWS::ImageBuilder::InfrastructureConfiguration"
}

resource_schema "aws_imagebuilder_lifecycle_policy" {
  cloudformation_type_name = "AWS::ImageBuilder::LifecyclePolicy"
}

resource_schema "aws_inspectorv2_cis_scan_configuration" {
  cloudformation_type_name = "AWS::InspectorV2::CisScanConfiguration"

  # Suppression Reason: Targets/TargetResourceTags is of unsupported type: ""
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1507
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_inspectorv2_filter" {
  cloudformation_type_name = "AWS::InspectorV2::Filter"
}

resource_schema "aws_internetmonitor_monitor" {
  cloudformation_type_name = "AWS::InternetMonitor::Monitor"
}

resource_schema "aws_iot_account_audit_configuration" {
  cloudformation_type_name = "AWS::IoT::AccountAuditConfiguration"
}

resource_schema "aws_iot_authorizer" {
  cloudformation_type_name = "AWS::IoT::Authorizer"
}

resource_schema "aws_iot_billing_group" {
  cloudformation_type_name = "AWS::IoT::BillingGroup"
}

resource_schema "aws_iot_ca_certificate" {
  cloudformation_type_name = "AWS::IoT::CACertificate"
}

resource_schema "aws_iot_certificate" {
  cloudformation_type_name = "AWS::IoT::Certificate"
}

resource_schema "aws_iot_certificate_provider" {
  cloudformation_type_name = "AWS::IoT::CertificateProvider"
}

resource_schema "aws_iot_custom_metric" {
  cloudformation_type_name = "AWS::IoT::CustomMetric"
}

resource_schema "aws_iot_dimension" {
  cloudformation_type_name = "AWS::IoT::Dimension"
}

resource_schema "aws_iot_domain_configuration" {
  cloudformation_type_name = "AWS::IoT::DomainConfiguration"
}

resource_schema "aws_iot_fleet_metric" {
  cloudformation_type_name = "AWS::IoT::FleetMetric"
}

resource_schema "aws_iot_job_template" {
  cloudformation_type_name = "AWS::IoT::JobTemplate"
}

resource_schema "aws_iot_logging" {
  cloudformation_type_name = "AWS::IoT::Logging"
}

resource_schema "aws_iot_mitigation_action" {
  cloudformation_type_name = "AWS::IoT::MitigationAction"
}

resource_schema "aws_iot_policy" {
  cloudformation_type_name = "AWS::IoT::Policy"
}

resource_schema "aws_iot_provisioning_template" {
  cloudformation_type_name = "AWS::IoT::ProvisioningTemplate"
}

resource_schema "aws_iot_resource_specific_logging" {
  cloudformation_type_name = "AWS::IoT::ResourceSpecificLogging"
}

resource_schema "aws_iot_role_alias" {
  cloudformation_type_name = "AWS::IoT::RoleAlias"
}

resource_schema "aws_iot_scheduled_audit" {
  cloudformation_type_name = "AWS::IoT::ScheduledAudit"
}

resource_schema "aws_iot_security_profile" {
  cloudformation_type_name = "AWS::IoT::SecurityProfile"
}

resource_schema "aws_iot_software_package" {
  cloudformation_type_name = "AWS::IoT::SoftwarePackage"
}

resource_schema "aws_iot_software_package_version" {
  cloudformation_type_name               = "AWS::IoT::SoftwarePackageVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_iot_thing" {
  cloudformation_type_name = "AWS::IoT::Thing"
}

resource_schema "aws_iot_thing_group" {
  cloudformation_type_name = "AWS::IoT::ThingGroup"
}

resource_schema "aws_iot_thing_type" {
  cloudformation_type_name = "AWS::IoT::ThingType"
}

resource_schema "aws_iot_topic_rule" {
  cloudformation_type_name = "AWS::IoT::TopicRule"
}

resource_schema "aws_iot_topic_rule_destination" {
  cloudformation_type_name = "AWS::IoT::TopicRuleDestination"
}

resource_schema "aws_iotanalytics_channel" {
  cloudformation_type_name = "AWS::IoTAnalytics::Channel"
}

resource_schema "aws_iotanalytics_dataset" {
  cloudformation_type_name = "AWS::IoTAnalytics::Dataset"
}

resource_schema "aws_iotanalytics_datastore" {
  cloudformation_type_name = "AWS::IoTAnalytics::Datastore"
}

resource_schema "aws_iotanalytics_pipeline" {
  cloudformation_type_name = "AWS::IoTAnalytics::Pipeline"
}

resource_schema "aws_iotcoredeviceadvisor_suite_definition" {
  cloudformation_type_name = "AWS::IoTCoreDeviceAdvisor::SuiteDefinition"
}

resource_schema "aws_iotevents_alarm_model" {
  cloudformation_type_name = "AWS::IoTEvents::AlarmModel"
}

resource_schema "aws_iotevents_detector_model" {
  cloudformation_type_name = "AWS::IoTEvents::DetectorModel"
}

resource_schema "aws_iotevents_input" {
  cloudformation_type_name = "AWS::IoTEvents::Input"
}

resource_schema "aws_iotfleethub_application" {
  cloudformation_type_name = "AWS::IoTFleetHub::Application"
}

resource_schema "aws_iotfleetwise_campaign" {
  cloudformation_type_name = "AWS::IoTFleetWise::Campaign"
}

resource_schema "aws_iotfleetwise_decoder_manifest" {
  cloudformation_type_name = "AWS::IoTFleetWise::DecoderManifest"

  # Suppression Reason: NetworkInterfaces is of unsupported type: list of ""
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1515
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_iotfleetwise_fleet" {
  cloudformation_type_name = "AWS::IoTFleetWise::Fleet"
}

resource_schema "aws_iotfleetwise_model_manifest" {
  cloudformation_type_name = "AWS::IoTFleetWise::ModelManifest"
}

resource_schema "aws_iotfleetwise_signal_catalog" {
  cloudformation_type_name = "AWS::IoTFleetWise::SignalCatalog"
}

resource_schema "aws_iotfleetwise_vehicle" {
  cloudformation_type_name = "AWS::IoTFleetWise::Vehicle"
}

resource_schema "aws_iotsitewise_access_policy" {
  cloudformation_type_name               = "AWS::IoTSiteWise::AccessPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_iotsitewise_asset" {
  cloudformation_type_name = "AWS::IoTSiteWise::Asset"
}

resource_schema "aws_iotsitewise_asset_model" {
  cloudformation_type_name = "AWS::IoTSiteWise::AssetModel"
}

resource_schema "aws_iotsitewise_dashboard" {
  cloudformation_type_name               = "AWS::IoTSiteWise::Dashboard"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_iotsitewise_gateway" {
  cloudformation_type_name = "AWS::IoTSiteWise::Gateway"
}

resource_schema "aws_iotsitewise_portal" {
  cloudformation_type_name = "AWS::IoTSiteWise::Portal"
}

resource_schema "aws_iotsitewise_project" {
  cloudformation_type_name               = "AWS::IoTSiteWise::Project"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_iottwinmaker_component_type" {
  cloudformation_type_name               = "AWS::IoTTwinMaker::ComponentType"
  suppress_plural_data_source_generation = true

  # Suppression Reason:
  # Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_iottwinmaker_entity" {
  cloudformation_type_name               = "AWS::IoTTwinMaker::Entity"
  suppress_plural_data_source_generation = true

  # Suppression Reason:
  # Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_iottwinmaker_scene" {
  cloudformation_type_name               = "AWS::IoTTwinMaker::Scene"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_iottwinmaker_sync_job" {
  cloudformation_type_name               = "AWS::IoTTwinMaker::SyncJob"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_iottwinmaker_workspace" {
  cloudformation_type_name = "AWS::IoTTwinMaker::Workspace"
}

resource_schema "aws_iotwireless_destination" {
  cloudformation_type_name = "AWS::IoTWireless::Destination"
}

resource_schema "aws_iotwireless_device_profile" {
  cloudformation_type_name = "AWS::IoTWireless::DeviceProfile"
}

resource_schema "aws_iotwireless_fuota_task" {
  cloudformation_type_name = "AWS::IoTWireless::FuotaTask"
}

resource_schema "aws_iotwireless_multicast_group" {
  cloudformation_type_name = "AWS::IoTWireless::MulticastGroup"
}

resource_schema "aws_iotwireless_network_analyzer_configuration" {
  cloudformation_type_name = "AWS::IoTWireless::NetworkAnalyzerConfiguration"
}

resource_schema "aws_iotwireless_partner_account" {
  cloudformation_type_name = "AWS::IoTWireless::PartnerAccount"
}

resource_schema "aws_iotwireless_service_profile" {
  cloudformation_type_name = "AWS::IoTWireless::ServiceProfile"
}

resource_schema "aws_iotwireless_task_definition" {
  cloudformation_type_name = "AWS::IoTWireless::TaskDefinition"
}

resource_schema "aws_iotwireless_wireless_device" {
  cloudformation_type_name = "AWS::IoTWireless::WirelessDevice"
}

resource_schema "aws_iotwireless_wireless_device_import_task" {
  cloudformation_type_name = "AWS::IoTWireless::WirelessDeviceImportTask"
}

resource_schema "aws_iotwireless_wireless_gateway" {
  cloudformation_type_name = "AWS::IoTWireless::WirelessGateway"
}

resource_schema "aws_kms_alias" {
  cloudformation_type_name = "AWS::KMS::Alias"
}

resource_schema "aws_kms_key" {
  cloudformation_type_name = "AWS::KMS::Key"
}

resource_schema "aws_kms_replica_key" {
  cloudformation_type_name = "AWS::KMS::ReplicaKey"
}

resource_schema "aws_kafkaconnect_connector" {
  cloudformation_type_name = "AWS::KafkaConnect::Connector"
}

resource_schema "aws_kafkaconnect_custom_plugin" {
  cloudformation_type_name = "AWS::KafkaConnect::CustomPlugin"
}

resource_schema "aws_kafkaconnect_worker_configuration" {
  cloudformation_type_name = "AWS::KafkaConnect::WorkerConfiguration"
}

resource_schema "aws_kendra_data_source" {
  cloudformation_type_name               = "AWS::Kendra::DataSource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_kendra_faq" {
  cloudformation_type_name               = "AWS::Kendra::Faq"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_kendra_index" {
  cloudformation_type_name = "AWS::Kendra::Index"
}

resource_schema "aws_kendraranking_execution_plan" {
  cloudformation_type_name = "AWS::KendraRanking::ExecutionPlan"
}

resource_schema "aws_kinesis_stream" {
  cloudformation_type_name = "AWS::Kinesis::Stream"
}

resource_schema "aws_kinesisanalyticsv2_application" {
  cloudformation_type_name = "AWS::KinesisAnalyticsV2::Application"
}

resource_schema "aws_kinesisfirehose_delivery_stream" {
  cloudformation_type_name = "AWS::KinesisFirehose::DeliveryStream"
}

resource_schema "aws_kinesisvideo_signaling_channel" {
  cloudformation_type_name               = "AWS::KinesisVideo::SignalingChannel"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_kinesisvideo_stream" {
  cloudformation_type_name               = "AWS::KinesisVideo::Stream"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_lakeformation_data_cells_filter" {
  cloudformation_type_name = "AWS::LakeFormation::DataCellsFilter"
}

resource_schema "aws_lakeformation_principal_permissions" {
  cloudformation_type_name               = "AWS::LakeFormation::PrincipalPermissions"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_lakeformation_tag" {
  cloudformation_type_name = "AWS::LakeFormation::Tag"
}

resource_schema "aws_lakeformation_tag_association" {
  cloudformation_type_name               = "AWS::LakeFormation::TagAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_lambda_code_signing_config" {
  cloudformation_type_name = "AWS::Lambda::CodeSigningConfig"
}

resource_schema "aws_lambda_event_invoke_config" {
  cloudformation_type_name               = "AWS::Lambda::EventInvokeConfig"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_lambda_event_source_mapping" {
  cloudformation_type_name = "AWS::Lambda::EventSourceMapping"
}

resource_schema "aws_lambda_function" {
  cloudformation_type_name = "AWS::Lambda::Function"
}

resource_schema "aws_lambda_layer_version_permission" {
  cloudformation_type_name               = "AWS::Lambda::LayerVersionPermission"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_lambda_permission" {
  cloudformation_type_name               = "AWS::Lambda::Permission"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_lambda_url" {
  cloudformation_type_name               = "AWS::Lambda::Url"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_lambda_version" {
  cloudformation_type_name               = "AWS::Lambda::Version"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_lex_bot" {
  cloudformation_type_name = "AWS::Lex::Bot"

  # Latest schema updates are suppressed.
  # git checkout internal/service/cloudformation/schemas/AWS_Lex_Bot.json
  # Suppression Reason:
  # Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95
}

resource_schema "aws_lex_bot_alias" {
  cloudformation_type_name               = "AWS::Lex::BotAlias"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_lex_bot_version" {
  cloudformation_type_name               = "AWS::Lex::BotVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_lex_resource_policy" {
  cloudformation_type_name               = "AWS::Lex::ResourcePolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_licensemanager_grant" {
  cloudformation_type_name = "AWS::LicenseManager::Grant"
}

resource_schema "aws_licensemanager_license" {
  cloudformation_type_name = "AWS::LicenseManager::License"
}

resource_schema "aws_lightsail_alarm" {
  cloudformation_type_name = "AWS::Lightsail::Alarm"
}

resource_schema "aws_lightsail_bucket" {
  cloudformation_type_name = "AWS::Lightsail::Bucket"
}

resource_schema "aws_lightsail_certificate" {
  cloudformation_type_name = "AWS::Lightsail::Certificate"
}

resource_schema "aws_lightsail_container" {
  cloudformation_type_name = "AWS::Lightsail::Container"
}

resource_schema "aws_lightsail_database" {
  cloudformation_type_name = "AWS::Lightsail::Database"
}

resource_schema "aws_lightsail_disk" {
  cloudformation_type_name = "AWS::Lightsail::Disk"
}

resource_schema "aws_lightsail_distribution" {
  cloudformation_type_name = "AWS::Lightsail::Distribution"
}

resource_schema "aws_lightsail_instance" {
  cloudformation_type_name = "AWS::Lightsail::Instance"
}

resource_schema "aws_lightsail_load_balancer" {
  cloudformation_type_name = "AWS::Lightsail::LoadBalancer"
}

resource_schema "aws_lightsail_load_balancer_tls_certificate" {
  cloudformation_type_name               = "AWS::Lightsail::LoadBalancerTlsCertificate"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_lightsail_static_ip" {
  cloudformation_type_name = "AWS::Lightsail::StaticIp"
}

resource_schema "aws_location_geofence_collection" {
  cloudformation_type_name = "AWS::Location::GeofenceCollection"
}

resource_schema "aws_location_api_key" {
  cloudformation_type_name = "AWS::Location::APIKey"
}

resource_schema "aws_location_map" {
  cloudformation_type_name = "AWS::Location::Map"
}

resource_schema "aws_location_place_index" {
  cloudformation_type_name = "AWS::Location::PlaceIndex"
}

resource_schema "aws_location_route_calculator" {
  cloudformation_type_name = "AWS::Location::RouteCalculator"
}

resource_schema "aws_location_tracker" {
  cloudformation_type_name = "AWS::Location::Tracker"
}

resource_schema "aws_location_tracker_consumer" {
  cloudformation_type_name               = "AWS::Location::TrackerConsumer"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_logs_account_policy" {
  cloudformation_type_name               = "AWS::Logs::AccountPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_logs_delivery" {
  cloudformation_type_name = "AWS::Logs::Delivery"
}

resource_schema "aws_logs_delivery_destination" {
  cloudformation_type_name = "AWS::Logs::DeliveryDestination"
}

resource_schema "aws_logs_delivery_source" {
  cloudformation_type_name = "AWS::Logs::DeliverySource"
}

resource_schema "aws_logs_destination" {
  cloudformation_type_name = "AWS::Logs::Destination"
}

resource_schema "aws_logs_log_anomaly_detector" {
  cloudformation_type_name = "AWS::Logs::LogAnomalyDetector"
}

resource_schema "aws_logs_log_group" {
  cloudformation_type_name = "AWS::Logs::LogGroup"
}

resource_schema "aws_logs_log_stream" {
  cloudformation_type_name               = "AWS::Logs::LogStream"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_logs_metric_filter" {
  cloudformation_type_name = "AWS::Logs::MetricFilter"
}

resource_schema "aws_logs_query_definition" {
  cloudformation_type_name = "AWS::Logs::QueryDefinition"
}

resource_schema "aws_logs_resource_policy" {
  cloudformation_type_name = "AWS::Logs::ResourcePolicy"
}

resource_schema "aws_lookoutequipment_inference_scheduler" {
  cloudformation_type_name = "AWS::LookoutEquipment::InferenceScheduler"
}

resource_schema "aws_lookoutmetrics_alert" {
  cloudformation_type_name = "AWS::LookoutMetrics::Alert"
}

resource_schema "aws_lookoutmetrics_anomaly_detector" {
  cloudformation_type_name = "AWS::LookoutMetrics::AnomalyDetector"
}

resource_schema "aws_lookoutvision_project" {
  cloudformation_type_name = "AWS::LookoutVision::Project"
}

resource_schema "aws_m2_application" {
  cloudformation_type_name = "AWS::M2::Application"
}

resource_schema "aws_m2_environment" {
  cloudformation_type_name = "AWS::M2::Environment"
}

resource_schema "aws_msk_batch_scram_secret" {
  cloudformation_type_name               = "AWS::MSK::BatchScramSecret"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_msk_cluster" {
  cloudformation_type_name = "AWS::MSK::Cluster"
}

resource_schema "aws_msk_cluster_policy" {
  cloudformation_type_name               = "AWS::MSK::ClusterPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_msk_configuration" {
  cloudformation_type_name = "AWS::MSK::Configuration"
}

resource_schema "aws_msk_replicator" {
  cloudformation_type_name = "AWS::MSK::Replicator"
}

resource_schema "aws_msk_serverless_cluster" {
  cloudformation_type_name = "AWS::MSK::ServerlessCluster"
}

resource_schema "aws_msk_vpc_connection" {
  cloudformation_type_name = "AWS::MSK::VpcConnection"
}

resource_schema "aws_mwaa_environment" {
  cloudformation_type_name = "AWS::MWAA::Environment"
}

resource_schema "aws_macie_allow_list" {
  cloudformation_type_name               = "AWS::Macie::AllowList"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_macie_custom_data_identifier" {
  cloudformation_type_name               = "AWS::Macie::CustomDataIdentifier"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_macie_findings_filter" {
  cloudformation_type_name               = "AWS::Macie::FindingsFilter"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_macie_session" {
  cloudformation_type_name = "AWS::Macie::Session"
}

resource_schema "aws_managedblockchain_accessor" {
  cloudformation_type_name = "AWS::ManagedBlockchain::Accessor"
}

resource_schema "aws_mediaconnect_bridge" {
  cloudformation_type_name = "AWS::MediaConnect::Bridge"
}

resource_schema "aws_mediaconnect_bridge_output" {
  cloudformation_type_name               = "AWS::MediaConnect::BridgeOutput"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediaconnect_bridge_source" {
  cloudformation_type_name               = "AWS::MediaConnect::BridgeSource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediaconnect_flow" {
  cloudformation_type_name = "AWS::MediaConnect::Flow"
}

resource_schema "aws_mediaconnect_flow_entitlement" {
  cloudformation_type_name               = "AWS::MediaConnect::FlowEntitlement"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediaconnect_flow_output" {
  cloudformation_type_name               = "AWS::MediaConnect::FlowOutput"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediaconnect_flow_source" {
  cloudformation_type_name               = "AWS::MediaConnect::FlowSource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediaconnect_flow_vpc_interface" {
  cloudformation_type_name               = "AWS::MediaConnect::FlowVpcInterface"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediaconnect_gateway" {
  cloudformation_type_name = "AWS::MediaConnect::Gateway"
}

resource_schema "aws_medialive_multiplex" {
  cloudformation_type_name = "AWS::MediaLive::Multiplex"
}

resource_schema "aws_medialive_multiplexprogram" {
  cloudformation_type_name               = "AWS::MediaLive::Multiplexprogram"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediapackage_asset" {
  cloudformation_type_name               = "AWS::MediaPackage::Asset"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediapackage_channel" {
  cloudformation_type_name = "AWS::MediaPackage::Channel"
}

resource_schema "aws_mediapackage_origin_endpoint" {
  cloudformation_type_name = "AWS::MediaPackage::OriginEndpoint"
}

resource_schema "aws_mediapackage_packaging_configuration" {
  cloudformation_type_name               = "AWS::MediaPackage::PackagingConfiguration"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediapackage_packaging_group" {
  cloudformation_type_name = "AWS::MediaPackage::PackagingGroup"
}

resource_schema "aws_mediapackagev2_channel" {
  cloudformation_type_name               = "AWS::MediaPackageV2::Channel"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediapackagev2_channel_group" {
  cloudformation_type_name = "AWS::MediaPackageV2::ChannelGroup"
}

resource_schema "aws_mediapackagev2_channel_policy" {
  cloudformation_type_name               = "AWS::MediaPackageV2::ChannelPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediapackagev2_origin_endpoint" {
  cloudformation_type_name               = "AWS::MediaPackageV2::OriginEndpoint"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediapackagev2_origin_endpoint_policy" {
  cloudformation_type_name               = "AWS::MediaPackageV2::OriginEndpointPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediatailor_channel" {
  cloudformation_type_name = "AWS::MediaTailor::Channel"
}

resource_schema "aws_mediatailor_channel_policy" {
  cloudformation_type_name               = "AWS::MediaTailor::ChannelPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediatailor_live_source" {
  cloudformation_type_name               = "AWS::MediaTailor::LiveSource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_mediatailor_playback_configuration" {
  cloudformation_type_name = "AWS::MediaTailor::PlaybackConfiguration"

  # Suppression Reason: ConfigurationAliases is of unsupported type: ""
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1506
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_mediatailor_vod_source" {
  cloudformation_type_name               = "AWS::MediaTailor::VodSource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_memorydb_acl" {
  cloudformation_type_name = "AWS::MemoryDB::ACL"
}

resource_schema "aws_memorydb_cluster" {
  cloudformation_type_name = "AWS::MemoryDB::Cluster"
}

resource_schema "aws_memorydb_parameter_group" {
  cloudformation_type_name = "AWS::MemoryDB::ParameterGroup"
}

resource_schema "aws_memorydb_subnet_group" {
  cloudformation_type_name = "AWS::MemoryDB::SubnetGroup"
}

resource_schema "aws_memorydb_user" {
  cloudformation_type_name = "AWS::MemoryDB::User"
}

resource_schema "aws_neptune_db_cluster" {
  cloudformation_type_name = "AWS::Neptune::DBCluster"
}

resource_schema "aws_neptunegraph_graph" {
  cloudformation_type_name = "AWS::NeptuneGraph::Graph"
}

resource_schema "aws_neptunegraph_private_graph_endpoint" {
  cloudformation_type_name               = "AWS::NeptuneGraph::PrivateGraphEndpoint"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_networkfirewall_firewall" {
  cloudformation_type_name = "AWS::NetworkFirewall::Firewall"
}

resource_schema "aws_networkfirewall_firewall_policy" {
  cloudformation_type_name = "AWS::NetworkFirewall::FirewallPolicy"
}

resource_schema "aws_networkfirewall_logging_configuration" {
  cloudformation_type_name               = "AWS::NetworkFirewall::LoggingConfiguration"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_networkfirewall_rule_group" {
  cloudformation_type_name = "AWS::NetworkFirewall::RuleGroup"
}

resource_schema "aws_networkfirewall_tls_inspection_configuration" {
  cloudformation_type_name = "AWS::NetworkFirewall::TLSInspectionConfiguration"
}

resource_schema "aws_networkmanager_connect_attachment" {
  cloudformation_type_name = "AWS::NetworkManager::ConnectAttachment"
}

resource_schema "aws_networkmanager_connect_peer" {
  cloudformation_type_name = "AWS::NetworkManager::ConnectPeer"
}

resource_schema "aws_networkmanager_core_network" {
  cloudformation_type_name = "AWS::NetworkManager::CoreNetwork"
}

resource_schema "aws_networkmanager_customer_gateway_association" {
  cloudformation_type_name               = "AWS::NetworkManager::CustomerGatewayAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_networkmanager_device" {
  cloudformation_type_name               = "AWS::NetworkManager::Device"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_networkmanager_global_network" {
  cloudformation_type_name = "AWS::NetworkManager::GlobalNetwork"
}

resource_schema "aws_networkmanager_link" {
  cloudformation_type_name = "AWS::NetworkManager::Link"
}

resource_schema "aws_networkmanager_link_association" {
  cloudformation_type_name               = "AWS::NetworkManager::LinkAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_networkmanager_site" {
  cloudformation_type_name               = "AWS::NetworkManager::Site"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_networkmanager_site_to_site_vpn_attachment" {
  cloudformation_type_name = "AWS::NetworkManager::SiteToSiteVpnAttachment"
}

resource_schema "aws_networkmanager_transit_gateway_peering" {
  cloudformation_type_name = "AWS::NetworkManager::TransitGatewayPeering"
}

resource_schema "aws_networkmanager_transit_gateway_registration" {
  cloudformation_type_name               = "AWS::NetworkManager::TransitGatewayRegistration"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_networkmanager_transit_gateway_route_table_attachment" {
  cloudformation_type_name = "AWS::NetworkManager::TransitGatewayRouteTableAttachment"
}

resource_schema "aws_networkmanager_vpc_attachment" {
  cloudformation_type_name = "AWS::NetworkManager::VpcAttachment"
}

resource_schema "aws_nimblestudio_launch_profile" {
  cloudformation_type_name               = "AWS::NimbleStudio::LaunchProfile"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_nimblestudio_streaming_image" {
  cloudformation_type_name               = "AWS::NimbleStudio::StreamingImage"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_nimblestudio_studio" {
  cloudformation_type_name = "AWS::NimbleStudio::Studio"
}

resource_schema "aws_nimblestudio_studio_component" {
  cloudformation_type_name               = "AWS::NimbleStudio::StudioComponent"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_osis_pipeline" {
  cloudformation_type_name = "AWS::OSIS::Pipeline"
}

resource_schema "aws_oam_link" {
  cloudformation_type_name = "AWS::Oam::Link"
}

resource_schema "aws_oam_sink" {
  cloudformation_type_name = "AWS::Oam::Sink"
}

resource_schema "aws_omics_annotation_store" {
  cloudformation_type_name = "AWS::Omics::AnnotationStore"

  # Suppression Reason: StoreOptions/TsvStoreOptions/Schema is of unsupported type: list of key-value map.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1517
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_omics_reference_store" {
  cloudformation_type_name = "AWS::Omics::ReferenceStore"
}

resource_schema "aws_omics_run_group" {
  cloudformation_type_name = "AWS::Omics::RunGroup"
}

resource_schema "aws_omics_sequence_store" {
  cloudformation_type_name = "AWS::Omics::SequenceStore"
}

resource_schema "aws_omics_variant_store" {
  cloudformation_type_name = "AWS::Omics::VariantStore"
}

resource_schema "aws_omics_workflow" {
  cloudformation_type_name = "AWS::Omics::Workflow"
}

resource_schema "aws_opensearchserverless_access_policy" {
  cloudformation_type_name = "AWS::OpenSearchServerless::AccessPolicy"
}

resource_schema "aws_opensearchserverless_collection" {
  cloudformation_type_name = "AWS::OpenSearchServerless::Collection"
}

resource_schema "aws_opensearchserverless_lifecycle_policy" {
  cloudformation_type_name               = "AWS::OpenSearchServerless::LifecyclePolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_opensearchserverless_security_config" {
  cloudformation_type_name = "AWS::OpenSearchServerless::SecurityConfig"
}

resource_schema "aws_opensearchserverless_security_policy" {
  cloudformation_type_name = "AWS::OpenSearchServerless::SecurityPolicy"
}

resource_schema "aws_opensearchserverless_vpc_endpoint" {
  cloudformation_type_name = "AWS::OpenSearchServerless::VpcEndpoint"
}


resource_schema "aws_opensearchservice_domain" {
  cloudformation_type_name = "AWS::OpenSearchService::Domain"
}

resource_schema "aws_opsworkscm_server" {
  cloudformation_type_name = "AWS::OpsWorksCM::Server"
}

resource_schema "aws_organizations_account" {
  cloudformation_type_name               = "AWS::Organizations::Account"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_organizations_organization" {
  cloudformation_type_name = "AWS::Organizations::Organization"
}

resource_schema "aws_organizations_organizational_unit" {
  cloudformation_type_name               = "AWS::Organizations::OrganizationalUnit"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_organizations_policy" {
  cloudformation_type_name               = "AWS::Organizations::Policy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_organizations_resource_policy" {
  cloudformation_type_name               = "AWS::Organizations::ResourcePolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_pcaconnectorad_connector" {
  cloudformation_type_name = "AWS::PCAConnectorAD::Connector"
}

resource_schema "aws_pcaconnectorad_directory_registration" {
  cloudformation_type_name = "AWS::PCAConnectorAD::DirectoryRegistration"
}

resource_schema "aws_pcaconnectorad_service_principal_name" {
  cloudformation_type_name               = "AWS::PCAConnectorAD::ServicePrincipalName"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_pcaconnectorad_template" {
  cloudformation_type_name               = "AWS::PCAConnectorAD::Template"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_pcaconnectorad_template_group_access_control_entry" {
  cloudformation_type_name               = "AWS::PCAConnectorAD::TemplateGroupAccessControlEntry"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_panorama_application_instance" {
  cloudformation_type_name = "AWS::Panorama::ApplicationInstance"
}

resource_schema "aws_panorama_package" {
  cloudformation_type_name = "AWS::Panorama::Package"
}

resource_schema "aws_panorama_package_version" {
  cloudformation_type_name               = "AWS::Panorama::PackageVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_personalize_dataset" {
  cloudformation_type_name = "AWS::Personalize::Dataset"
}

resource_schema "aws_personalize_dataset_group" {
  cloudformation_type_name = "AWS::Personalize::DatasetGroup"
}

resource_schema "aws_personalize_schema" {
  cloudformation_type_name = "AWS::Personalize::Schema"
}

resource_schema "aws_personalize_solution" {
  cloudformation_type_name = "AWS::Personalize::Solution"
}

resource_schema "aws_pinpoint_in_app_template" {
  cloudformation_type_name = "AWS::Pinpoint::InAppTemplate"
}

resource_schema "aws_pipes_pipe" {
  cloudformation_type_name = "AWS::Pipes::Pipe"
}

resource_schema "aws_proton_environment_account_connection" {
  cloudformation_type_name = "AWS::Proton::EnvironmentAccountConnection"
}

resource_schema "aws_proton_environment_template" {
  cloudformation_type_name = "AWS::Proton::EnvironmentTemplate"
}

resource_schema "aws_proton_service_template" {
  cloudformation_type_name = "AWS::Proton::ServiceTemplate"
}

resource_schema "aws_qldb_stream" {
  cloudformation_type_name               = "AWS::QLDB::Stream"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_quicksight_analysis" {
  cloudformation_type_name               = "AWS::QuickSight::Analysis"
  suppress_plural_data_source_generation = true

  # Latest schema updates are suppressed.
  # git checkout internal/service/cloudformation/schemas/AWS_QuickSight_Analysis.json
  # Suppression Reason: internal compiler error: NewBulk too big.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1531
}

resource_schema "aws_quicksight_dashboard" {
  cloudformation_type_name               = "AWS::QuickSight::Dashboard"
  suppress_plural_data_source_generation = true

  # Latest schema updates are suppressed.
  # git checkout internal/service/cloudformation/schemas/AWS_QuickSight_Dashboard.json
  # Suppression Reason: internal compiler error: NewBulk too big.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1531
}

resource_schema "aws_quicksight_data_set" {
  cloudformation_type_name               = "AWS::QuickSight::DataSet"
  suppress_plural_data_source_generation = true

  # Latest schema updates are suppressed.
  # git checkout internal/service/cloudformation/schemas/AWS_QuickSight_DataSet.json
  # Suppression Reason: internal compiler error: NewBulk too big.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1531
}

resource_schema "aws_quicksight_data_source" {
  cloudformation_type_name               = "AWS::QuickSight::DataSource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_quicksight_refresh_schedule" {
  cloudformation_type_name               = "AWS::QuickSight::RefreshSchedule"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_quicksight_template" {
  cloudformation_type_name               = "AWS::QuickSight::Template"
  suppress_plural_data_source_generation = true

  # Latest schema updates are suppressed.
  # git checkout internal/service/cloudformation/schemas/AWS_QuickSight_Template.json
  # Suppression Reason: internal compiler error: NewBulk too big.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1531
}

resource_schema "aws_quicksight_theme" {
  cloudformation_type_name               = "AWS::QuickSight::Theme"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_quicksight_topic" {
  cloudformation_type_name               = "AWS::QuickSight::Topic"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_quicksight_vpc_connection" {
  cloudformation_type_name               = "AWS::QuickSight::VPCConnection"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ram_permission" {
  cloudformation_type_name = "AWS::RAM::Permission"
}

resource_schema "aws_rds_custom_db_engine_version" {
  cloudformation_type_name = "AWS::RDS::CustomDBEngineVersion"
}

resource_schema "aws_rds_db_cluster" {
  cloudformation_type_name = "AWS::RDS::DBCluster"
}

resource_schema "aws_rds_db_cluster_parameter_group" {
  cloudformation_type_name = "AWS::RDS::DBClusterParameterGroup"
}

resource_schema "aws_rds_db_instance" {
  cloudformation_type_name = "AWS::RDS::DBInstance"
}

resource_schema "aws_rds_db_parameter_group" {
  cloudformation_type_name = "AWS::RDS::DBParameterGroup"
}

resource_schema "aws_rds_db_proxy" {
  cloudformation_type_name = "AWS::RDS::DBProxy"
}

resource_schema "aws_rds_db_proxy_endpoint" {
  cloudformation_type_name = "AWS::RDS::DBProxyEndpoint"
}

resource_schema "aws_rds_db_proxy_target_group" {
  cloudformation_type_name               = "AWS::RDS::DBProxyTargetGroup"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_rds_db_subnet_group" {
  cloudformation_type_name = "AWS::RDS::DBSubnetGroup"
}

resource_schema "aws_rds_event_subscription" {
  cloudformation_type_name = "AWS::RDS::EventSubscription"
}

resource_schema "aws_rds_global_cluster" {
  cloudformation_type_name = "AWS::RDS::GlobalCluster"
}

resource_schema "aws_rds_integration" {
  cloudformation_type_name = "AWS::RDS::Integration"
}

resource_schema "aws_rds_option_group" {
  cloudformation_type_name = "AWS::RDS::OptionGroup"
}

resource_schema "aws_rum_app_monitor" {
  cloudformation_type_name = "AWS::RUM::AppMonitor"
}

resource_schema "aws_redshift_cluster" {
  cloudformation_type_name = "AWS::Redshift::Cluster"
}

resource_schema "aws_redshift_cluster_parameter_group" {
  cloudformation_type_name = "AWS::Redshift::ClusterParameterGroup"
}

resource_schema "aws_redshift_cluster_subnet_group" {
  cloudformation_type_name = "AWS::Redshift::ClusterSubnetGroup"
}

resource_schema "aws_redshift_endpoint_access" {
  cloudformation_type_name = "AWS::Redshift::EndpointAccess"
}

resource_schema "aws_redshift_endpoint_authorization" {
  cloudformation_type_name = "AWS::Redshift::EndpointAuthorization"
}

resource_schema "aws_redshift_event_subscription" {
  cloudformation_type_name = "AWS::Redshift::EventSubscription"
}

resource_schema "aws_redshift_scheduled_action" {
  cloudformation_type_name = "AWS::Redshift::ScheduledAction"
}

resource_schema "aws_redshiftserverless_namespace" {
  cloudformation_type_name = "AWS::RedshiftServerless::Namespace"
}

resource_schema "aws_redshiftserverless_workgroup" {
  cloudformation_type_name = "AWS::RedshiftServerless::Workgroup"
}

resource_schema "aws_refactorspaces_application" {
  cloudformation_type_name               = "AWS::RefactorSpaces::Application"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_refactorspaces_environment" {
  cloudformation_type_name = "AWS::RefactorSpaces::Environment"
}

resource_schema "aws_refactorspaces_route" {
  cloudformation_type_name               = "AWS::RefactorSpaces::Route"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_refactorspaces_service" {
  cloudformation_type_name               = "AWS::RefactorSpaces::Service"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_rekognition_collection" {
  cloudformation_type_name = "AWS::Rekognition::Collection"
}

resource_schema "aws_rekognition_project" {
  cloudformation_type_name = "AWS::Rekognition::Project"
}

resource_schema "aws_rekognition_stream_processor" {
  cloudformation_type_name = "AWS::Rekognition::StreamProcessor"

  # Suppression Reason: PolygonRegionsOfInterest is of unsupported type: set of array.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1516
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_resiliencehub_app" {
  cloudformation_type_name = "AWS::ResilienceHub::App"
}

resource_schema "aws_resiliencehub_resiliency_policy" {
  cloudformation_type_name = "AWS::ResilienceHub::ResiliencyPolicy"
}

resource_schema "aws_resourceexplorer2_default_view_association" {
  cloudformation_type_name               = "AWS::ResourceExplorer2::DefaultViewAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_resourceexplorer2_index" {
  cloudformation_type_name = "AWS::ResourceExplorer2::Index"
}

resource_schema "aws_resourceexplorer2_view" {
  cloudformation_type_name = "AWS::ResourceExplorer2::View"
}

resource_schema "aws_resourcegroups_group" {
  cloudformation_type_name = "AWS::ResourceGroups::Group"
}

resource_schema "aws_robomaker_fleet" {
  cloudformation_type_name = "AWS::RoboMaker::Fleet"
}

resource_schema "aws_robomaker_robot" {
  cloudformation_type_name = "AWS::RoboMaker::Robot"
}

resource_schema "aws_robomaker_robot_application" {
  cloudformation_type_name = "AWS::RoboMaker::RobotApplication"
}

resource_schema "aws_robomaker_robot_application_version" {
  cloudformation_type_name               = "AWS::RoboMaker::RobotApplicationVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_robomaker_simulation_application" {
  cloudformation_type_name = "AWS::RoboMaker::SimulationApplication"
}

resource_schema "aws_robomaker_simulation_application_version" {
  cloudformation_type_name               = "AWS::RoboMaker::SimulationApplicationVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_rolesanywhere_crl" {
  cloudformation_type_name = "AWS::RolesAnywhere::CRL"
}

resource_schema "aws_rolesanywhere_profile" {
  cloudformation_type_name = "AWS::RolesAnywhere::Profile"
}

resource_schema "aws_rolesanywhere_trust_anchor" {
  cloudformation_type_name               = "AWS::RolesAnywhere::TrustAnchor"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_route53_cidr_collection" {
  cloudformation_type_name = "AWS::Route53::CidrCollection"
}

resource_schema "aws_route53_dnssec" {
  cloudformation_type_name = "AWS::Route53::DNSSEC"
}

resource_schema "aws_route53_health_check" {
  cloudformation_type_name = "AWS::Route53::HealthCheck"
}

resource_schema "aws_route53_hosted_zone" {
  cloudformation_type_name = "AWS::Route53::HostedZone"
}

resource_schema "aws_route53_key_signing_key" {
  cloudformation_type_name = "AWS::Route53::KeySigningKey"
}

resource_schema "aws_route53recoverycontrol_cluster" {
  cloudformation_type_name = "AWS::Route53RecoveryControl::Cluster"
}

resource_schema "aws_route53recoverycontrol_control_panel" {
  cloudformation_type_name = "AWS::Route53RecoveryControl::ControlPanel"
}

resource_schema "aws_route53recoverycontrol_routing_control" {
  cloudformation_type_name               = "AWS::Route53RecoveryControl::RoutingControl"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_route53recoverycontrol_safety_rule" {
  cloudformation_type_name               = "AWS::Route53RecoveryControl::SafetyRule"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_route53recoveryreadiness_cell" {
  cloudformation_type_name = "AWS::Route53RecoveryReadiness::Cell"
}

resource_schema "aws_route53recoveryreadiness_readiness_check" {
  cloudformation_type_name = "AWS::Route53RecoveryReadiness::ReadinessCheck"
}

resource_schema "aws_route53recoveryreadiness_recovery_group" {
  cloudformation_type_name = "AWS::Route53RecoveryReadiness::RecoveryGroup"
}

resource_schema "aws_route53recoveryreadiness_resource_set" {
  cloudformation_type_name = "AWS::Route53RecoveryReadiness::ResourceSet"
}

resource_schema "aws_route53resolver_firewall_domain_list" {
  cloudformation_type_name = "AWS::Route53Resolver::FirewallDomainList"
}

resource_schema "aws_route53resolver_firewall_rule_group" {
  cloudformation_type_name = "AWS::Route53Resolver::FirewallRuleGroup"
}

resource_schema "aws_route53resolver_firewall_rule_group_association" {
  cloudformation_type_name = "AWS::Route53Resolver::FirewallRuleGroupAssociation"
}

resource_schema "aws_route53resolver_resolver_config" {
  cloudformation_type_name = "AWS::Route53Resolver::ResolverConfig"
}

resource_schema "aws_route53resolver_resolver_dnssec_config" {
  cloudformation_type_name = "AWS::Route53Resolver::ResolverDNSSECConfig"
}

resource_schema "aws_route53resolver_resolver_rule_association" {
  cloudformation_type_name = "AWS::Route53Resolver::ResolverRuleAssociation"
}

resource_schema "aws_route53resolver_resolver_query_logging_config" {
  cloudformation_type_name = "AWS::Route53Resolver::ResolverQueryLoggingConfig"
}

resource_schema "aws_route53resolver_resolver_query_logging_config_association" {
  cloudformation_type_name = "AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation"
}

resource_schema "aws_route53resolver_resolver_rule" {
  cloudformation_type_name = "AWS::Route53Resolver::ResolverRule"
}

resource_schema "aws_s3_access_grant" {
  cloudformation_type_name               = "AWS::S3::AccessGrant"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_s3_access_grants_instance" {
  cloudformation_type_name = "AWS::S3::AccessGrantsInstance"
}

resource_schema "aws_s3_access_grants_location" {
  cloudformation_type_name               = "AWS::S3::AccessGrantsLocation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_s3_access_point" {
  cloudformation_type_name = "AWS::S3::AccessPoint"
}

resource_schema "aws_s3_bucket" {
  cloudformation_type_name = "AWS::S3::Bucket"
}

resource_schema "aws_s3_bucket_policy" {
  cloudformation_type_name = "AWS::S3::BucketPolicy"
}

resource_schema "aws_s3_multi_region_access_point" {
  cloudformation_type_name = "AWS::S3::MultiRegionAccessPoint"
}

resource_schema "aws_s3_multi_region_access_point_policy" {
  cloudformation_type_name               = "AWS::S3::MultiRegionAccessPointPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_s3_storage_lens" {
  cloudformation_type_name = "AWS::S3::StorageLens"
}

resource_schema "aws_s3_storage_lens_group" {
  cloudformation_type_name = "AWS::S3::StorageLensGroup"
}

resource_schema "aws_s3express_bucket_policy" {
  cloudformation_type_name = "AWS::S3Express::BucketPolicy"
}

resource_schema "aws_s3express_directory_bucket" {
  cloudformation_type_name = "AWS::S3Express::DirectoryBucket"
}

resource_schema "aws_s3objectlambda_access_point" {
  cloudformation_type_name = "AWS::S3ObjectLambda::AccessPoint"
}

resource_schema "aws_s3objectlambda_access_point_policy" {
  cloudformation_type_name               = "AWS::S3ObjectLambda::AccessPointPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_s3outposts_access_point" {
  cloudformation_type_name               = "AWS::S3Outposts::AccessPoint"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_s3outposts_bucket" {
  cloudformation_type_name               = "AWS::S3Outposts::Bucket"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_s3outposts_bucket_policy" {
  cloudformation_type_name               = "AWS::S3Outposts::BucketPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_s3outposts_endpoint" {
  cloudformation_type_name = "AWS::S3Outposts::Endpoint"
}

resource_schema "aws_ses_configuration_set" {
  cloudformation_type_name = "AWS::SES::ConfigurationSet"
}

resource_schema "aws_ses_configuration_set_event_destination" {
  cloudformation_type_name               = "AWS::SES::ConfigurationSetEventDestination"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ses_contact_list" {
  cloudformation_type_name = "AWS::SES::ContactList"
}

resource_schema "aws_ses_dedicated_ip_pool" {
  cloudformation_type_name = "AWS::SES::DedicatedIpPool"
}

resource_schema "aws_ses_email_identity" {
  cloudformation_type_name = "AWS::SES::EmailIdentity"
}

resource_schema "aws_ses_template" {
  cloudformation_type_name = "AWS::SES::Template"
}

resource_schema "aws_ses_vdm_attributes" {
  cloudformation_type_name               = "AWS::SES::VdmAttributes"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_sns_topic" {
  cloudformation_type_name = "AWS::SNS::Topic"
}

resource_schema "aws_sns_topic_inline_policy" {
  cloudformation_type_name               = "AWS::SNS::TopicInlinePolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_sns_topic_policy" {
  cloudformation_type_name               = "AWS::SNS::TopicPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ssm_association" {
  cloudformation_type_name               = "AWS::SSM::Association"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ssm_document" {
  cloudformation_type_name = "AWS::SSM::Document"
}

resource_schema "aws_ssm_parameter" {
  cloudformation_type_name               = "AWS::SSM::Parameter"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ssm_patch_baseline" {
  cloudformation_type_name = "AWS::SSM::PatchBaseline"
}

resource_schema "aws_ssm_resource_data_sync" {
  cloudformation_type_name = "AWS::SSM::ResourceDataSync"
}

resource_schema "aws_ssm_resource_policy" {
  cloudformation_type_name = "AWS::SSM::ResourcePolicy"
}

resource_schema "aws_ssmcontacts_contact" {
  cloudformation_type_name = "AWS::SSMContacts::Contact"
}

resource_schema "aws_ssmcontacts_contact_channel" {
  cloudformation_type_name               = "AWS::SSMContacts::ContactChannel"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ssmcontacts_plan" {
  cloudformation_type_name               = "AWS::SSMContacts::Plan"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ssmcontacts_rotation" {
  cloudformation_type_name               = "AWS::SSMContacts::Rotation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_ssmincidents_replication_set" {
  cloudformation_type_name = "AWS::SSMIncidents::ReplicationSet"
}

resource_schema "aws_ssmincidents_response_plan" {
  cloudformation_type_name = "AWS::SSMIncidents::ResponsePlan"
}

resource_schema "aws_sso_assignment" {
  cloudformation_type_name               = "AWS::SSO::Assignment"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_sso_instance_access_control_attribute_configuration" {
  cloudformation_type_name               = "AWS::SSO::InstanceAccessControlAttributeConfiguration"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_sso_permission_set" {
  cloudformation_type_name               = "AWS::SSO::PermissionSet"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_sqs_queue" {
  cloudformation_type_name = "AWS::SQS::Queue"
}

resource_schema "aws_sqs_queue_inline_policy" {
  cloudformation_type_name               = "AWS::SQS::QueueInlinePolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_sagemaker_app" {
  cloudformation_type_name = "AWS::SageMaker::App"
}

resource_schema "aws_sagemaker_app_image_config" {
  cloudformation_type_name = "AWS::SageMaker::AppImageConfig"
}

resource_schema "aws_sagemaker_data_quality_job_definition" {
  cloudformation_type_name = "AWS::SageMaker::DataQualityJobDefinition"
}

resource_schema "aws_sagemaker_device" {
  cloudformation_type_name               = "AWS::SageMaker::Device"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_sagemaker_device_fleet" {
  cloudformation_type_name               = "AWS::SageMaker::DeviceFleet"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_sagemaker_domain" {
  cloudformation_type_name = "AWS::SageMaker::Domain"
}

resource_schema "aws_sagemaker_feature_group" {
  cloudformation_type_name = "AWS::SageMaker::FeatureGroup"
}

resource_schema "aws_sagemaker_image" {
  cloudformation_type_name = "AWS::SageMaker::Image"
}

resource_schema "aws_sagemaker_image_version" {
  cloudformation_type_name               = "AWS::SageMaker::ImageVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_sagemaker_inference_component" {
  cloudformation_type_name = "AWS::SageMaker::InferenceComponent"
}

resource_schema "aws_sagemaker_inference_experiment" {
  cloudformation_type_name = "AWS::SageMaker::InferenceExperiment"
}

resource_schema "aws_sagemaker_model_bias_job_definition" {
  cloudformation_type_name = "AWS::SageMaker::ModelBiasJobDefinition"
}

resource_schema "aws_sagemaker_model_card" {
  cloudformation_type_name = "AWS::SageMaker::ModelCard"

  # Suppression Reason: Properties with anyOf indicator
  # Content/EvaluationDetails/MetricGroups/MetricData is of unsupported type: list of ""
  # https://github.com/hashicorp/terraform-provider-awscc/issues/98
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_sagemaker_model_explainability_job_definition" {
  cloudformation_type_name = "AWS::SageMaker::ModelExplainabilityJobDefinition"
}

resource_schema "aws_sagemaker_model_package" {
  cloudformation_type_name = "AWS::SageMaker::ModelPackage"
}

resource_schema "aws_sagemaker_model_package_group" {
  cloudformation_type_name = "AWS::SageMaker::ModelPackageGroup"
}

resource_schema "aws_sagemaker_model_quality_job_definition" {
  cloudformation_type_name = "AWS::SageMaker::ModelQualityJobDefinition"
}

resource_schema "aws_sagemaker_monitoring_schedule" {
  cloudformation_type_name = "AWS::SageMaker::MonitoringSchedule"
}

resource_schema "aws_sagemaker_pipeline" {
  cloudformation_type_name = "AWS::SageMaker::Pipeline"
}

resource_schema "aws_sagemaker_project" {
  cloudformation_type_name = "AWS::SageMaker::Project"
}

resource_schema "aws_sagemaker_space" {
  cloudformation_type_name = "AWS::SageMaker::Space"
}

resource_schema "aws_sagemaker_user_profile" {
  cloudformation_type_name = "AWS::SageMaker::UserProfile"
}

resource_schema "aws_scheduler_schedule" {
  cloudformation_type_name = "AWS::Scheduler::Schedule"

  # Suppression Reason: Target/EcsParameters/Tags is of unsupported type: list of key-value map.
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1517
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_scheduler_schedule_group" {
  cloudformation_type_name = "AWS::Scheduler::ScheduleGroup"
}

resource_schema "aws_secretsmanager_secret" {
  cloudformation_type_name = "AWS::SecretsManager::Secret"
}

resource_schema "aws_securityhub_automation_rule" {
  cloudformation_type_name = "AWS::SecurityHub::AutomationRule"

  # Suppression Reason:
  # Properties with anyOf indicator https://github.com/hashicorp/terraform-provider-awscc/issues/98
  # Actions/FindingFieldsUpdate/Note/UpdatedBy
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
  suppress_plural_data_source_generation   = true
}

resource_schema "aws_securityhub_hub" {
  cloudformation_type_name = "AWS::SecurityHub::Hub"
}

resource_schema "aws_securityhub_delegated_admin" {
  cloudformation_type_name               = "AWS::SecurityHub::DelegatedAdmin"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_securityhub_insight" {
  cloudformation_type_name               = "AWS::SecurityHub::Insight"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_securityhub_product_subscription" {
  cloudformation_type_name               = "AWS::SecurityHub::ProductSubscription"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_securityhub_standard" {
  cloudformation_type_name               = "AWS::SecurityHub::Standard"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_securitylake_aws_log_source" {
  cloudformation_type_name               = "AWS::SecurityLake::AwsLogSource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_securitylake_data_lake" {
  cloudformation_type_name               = "AWS::SecurityLake::DataLake"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_securitylake_subscriber" {
  cloudformation_type_name               = "AWS::SecurityLake::Subscriber"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_servicecatalog_cloudformation_provisioned_product" {
  cloudformation_type_name               = "AWS::ServiceCatalog::CloudFormationProvisionedProduct"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_servicecatalog_service_action" {
  cloudformation_type_name = "AWS::ServiceCatalog::ServiceAction"
}

resource_schema "aws_servicecatalog_service_action_association" {
  cloudformation_type_name               = "AWS::ServiceCatalog::ServiceActionAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_servicecatalogappregistry_application" {
  cloudformation_type_name = "AWS::ServiceCatalogAppRegistry::Application"
}

resource_schema "aws_servicecatalogappregistry_attribute_group" {
  cloudformation_type_name = "AWS::ServiceCatalogAppRegistry::AttributeGroup"
}

resource_schema "aws_servicecatalogappregistry_attribute_group_association" {
  cloudformation_type_name               = "AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_servicecatalogappregistry_resource_association" {
  cloudformation_type_name               = "AWS::ServiceCatalogAppRegistry::ResourceAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_simspaceweaver_simulation" {
  cloudformation_type_name = "AWS::SimSpaceWeaver::Simulation"
}

resource_schema "aws_signer_profile_permission" {
  cloudformation_type_name               = "AWS::Signer::ProfilePermission"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_signer_signing_profile" {
  cloudformation_type_name = "AWS::Signer::SigningProfile"
}

resource_schema "aws_shield_drt_access" {
  cloudformation_type_name = "AWS::Shield::DRTAccess"
}

resource_schema "aws_shield_proactive_engagement" {
  cloudformation_type_name = "AWS::Shield::ProactiveEngagement"
}

resource_schema "aws_shield_protection" {
  cloudformation_type_name = "AWS::Shield::Protection"
}

resource_schema "aws_shield_protection_group" {
  cloudformation_type_name = "AWS::Shield::ProtectionGroup"
}

resource_schema "aws_stepfunctions_activity" {
  cloudformation_type_name               = "AWS::StepFunctions::Activity"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_stepfunctions_state_machine" {
  cloudformation_type_name = "AWS::StepFunctions::StateMachine"

  # Latest schema updates are suppressed.
  # git checkout internal/service/cloudformation/schemas/AWS_StepFunctions_StateMachine.json
  # Suppression Reason: DefinitionSubstitutions is of unsupported type: key-value map of .
  # https://github.com/hashicorp/terraform-provider-awscc/issues/1530
}

resource_schema "aws_stepfunctions_state_machine_alias" {
  cloudformation_type_name               = "AWS::StepFunctions::StateMachineAlias"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_stepfunctions_state_machine_version" {
  cloudformation_type_name               = "AWS::StepFunctions::StateMachineVersion"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_supportapp_account_alias" {
  cloudformation_type_name = "AWS::SupportApp::AccountAlias"
}

resource_schema "aws_supportapp_slack_channel_configuration" {
  cloudformation_type_name = "AWS::SupportApp::SlackChannelConfiguration"
}

resource_schema "aws_supportapp_slack_workspace_configuration" {
  cloudformation_type_name = "AWS::SupportApp::SlackWorkspaceConfiguration"
}

resource_schema "aws_synthetics_canary" {
  cloudformation_type_name = "AWS::Synthetics::Canary"
}

resource_schema "aws_synthetics_group" {
  cloudformation_type_name = "AWS::Synthetics::Group"
}

resource_schema "aws_systemsmanagersap_application" {
  cloudformation_type_name = "AWS::SystemsManagerSAP::Application"
}

resource_schema "aws_timestream_database" {
  cloudformation_type_name = "AWS::Timestream::Database"
}

resource_schema "aws_timestream_scheduled_query" {
  cloudformation_type_name = "AWS::Timestream::ScheduledQuery"
}

resource_schema "aws_timestream_table" {
  cloudformation_type_name = "AWS::Timestream::Table"
}

resource_schema "aws_transfer_agreement" {
  cloudformation_type_name               = "AWS::Transfer::Agreement"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_transfer_certificate" {
  cloudformation_type_name = "AWS::Transfer::Certificate"
}

resource_schema "aws_transfer_connector" {
  cloudformation_type_name = "AWS::Transfer::Connector"
}

resource_schema "aws_transfer_profile" {
  cloudformation_type_name = "AWS::Transfer::Profile"
}

resource_schema "aws_transfer_workflow" {
  cloudformation_type_name = "AWS::Transfer::Workflow"
}

resource_schema "aws_voiceid_domain" {
  cloudformation_type_name = "AWS::VoiceID::Domain"
}

resource_schema "aws_verifiedpermissions_identity_source" {
  cloudformation_type_name               = "AWS::VerifiedPermissions::IdentitySource"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_verifiedpermissions_policy" {
  cloudformation_type_name               = "AWS::VerifiedPermissions::Policy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_verifiedpermissions_policy_store" {
  cloudformation_type_name = "AWS::VerifiedPermissions::PolicyStore"
}

resource_schema "aws_verifiedpermissions_policy_template" {
  cloudformation_type_name               = "AWS::VerifiedPermissions::PolicyTemplate"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_vpclattice_access_log_subscription" {
  cloudformation_type_name = "AWS::VpcLattice::AccessLogSubscription"
}

resource_schema "aws_vpclattice_auth_policy" {
  cloudformation_type_name               = "AWS::VpcLattice::AuthPolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_vpclattice_listener" {
  cloudformation_type_name = "AWS::VpcLattice::Listener"
}

resource_schema "aws_vpclattice_resource_policy" {
  cloudformation_type_name               = "AWS::VpcLattice::ResourcePolicy"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_vpclattice_rule" {
  cloudformation_type_name = "AWS::VpcLattice::Rule"
}

resource_schema "aws_vpclattice_service" {
  cloudformation_type_name = "AWS::VpcLattice::Service"
}

resource_schema "aws_vpclattice_service_network" {
  cloudformation_type_name = "AWS::VpcLattice::ServiceNetwork"
}

resource_schema "aws_vpclattice_service_network_service_association" {
  cloudformation_type_name = "AWS::VpcLattice::ServiceNetworkServiceAssociation"
}

resource_schema "aws_vpclattice_service_network_vpc_association" {
  cloudformation_type_name = "AWS::VpcLattice::ServiceNetworkVpcAssociation"
}

resource_schema "aws_vpclattice_target_group" {
  cloudformation_type_name = "AWS::VpcLattice::TargetGroup"
}

resource_schema "aws_wafv2_ip_set" {
  cloudformation_type_name               = "AWS::WAFv2::IPSet"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_wafv2_logging_configuration" {
  cloudformation_type_name = "AWS::WAFv2::LoggingConfiguration"
}

resource_schema "aws_wafv2_regex_pattern_set" {
  cloudformation_type_name               = "AWS::WAFv2::RegexPatternSet"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_wafv2_rule_group" {
  cloudformation_type_name               = "AWS::WAFv2::RuleGroup"
  suppress_plural_data_source_generation = true

  # Suppression Reason:
  # Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_wafv2_web_acl" {
  cloudformation_type_name               = "AWS::WAFv2::WebACL"
  suppress_plural_data_source_generation = true

  # Suppression Reason:
  # Recursive Attribute Definitions https://github.com/hashicorp/terraform-provider-awscc/issues/95
  suppress_resource_generation             = true
  suppress_singular_data_source_generation = true
}

resource_schema "aws_wafv2_web_acl_association" {
  cloudformation_type_name               = "AWS::WAFv2::WebACLAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_wisdom_assistant" {
  cloudformation_type_name = "AWS::Wisdom::Assistant"
}

resource_schema "aws_wisdom_assistant_association" {
  cloudformation_type_name               = "AWS::Wisdom::AssistantAssociation"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_wisdom_knowledge_base" {
  cloudformation_type_name = "AWS::Wisdom::KnowledgeBase"
}

resource_schema "aws_workspaces_connection_alias" {
  cloudformation_type_name               = "AWS::WorkSpaces::ConnectionAlias"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_workspacesthinclient_environment" {
  cloudformation_type_name = "AWS::WorkSpacesThinClient::Environment"
}

resource_schema "aws_xray_group" {
  cloudformation_type_name               = "AWS::XRay::Group"
  suppress_plural_data_source_generation = true
}

resource_schema "aws_xray_resource_policy" {
  cloudformation_type_name = "AWS::XRay::ResourcePolicy"
}

resource_schema "aws_xray_sampling_rule" {
  cloudformation_type_name = "AWS::XRay::SamplingRule"
}