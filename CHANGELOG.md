## 0.21.0 (Unreleased)
## [0.20.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.20.0) (May  5, 2022)

FEATURES:

* **New Data Source:** `awscc_ec2_transit_gateway_attachment`
* **New Data Source:** `awscc_ec2_transit_gateway_attachments`
* **New Data Source:** `awscc_voiceid_domain`
* **New Data Source:** `awscc_voiceid_domains`
* **New Resource:** `awscc_ec2_transit_gateway_attachment`
* **New Resource:** `awscc_voiceid_domain`
* Adds support for assuming IAM role with web identity.

BUG FIXES:

* Provider parameter `skip_medatadata_api_check = false` now correctly overrides environment variable `AWS_EC2_METADATA_DISABLED`

## [0.19.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.19.0) (April 21, 2022)

FEATURES:

* **New Data Source:** `awscc_connect_phone_number`
* **New Data Source:** `awscc_iottwinmaker_scene`
* **New Data Source:** `awscc_iottwinmaker_workspace`
* **New Data Source:** `awscc_iottwinmaker_workspaces`
* **New Resource:** `awscc_connect_phone_number`
* **New Resource:** `awscc_iottwinmaker_scene`
* **New Resource:** `awscc_iottwinmaker_workspace`

## [0.18.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.18.0) (April 14, 2022)

FEATURES:

* **New Data Source:** `awscc_apprunner_observability_configuration`
* **New Data Source:** `awscc_apprunner_observability_configurations`
* **New Data Source:** `awscc_ssm_documents`
* **New Resource:** `awscc_apprunner_observability_configuration`

## [0.17.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.17.0) (April  7, 2022)

FEATURES:

* **New Data Source:** `awscc_datasync_location_fsx_open_zfs`
* **New Data Source:** `awscc_datasync_location_fsx_open_zfs_plural`
* **New Data Source:** `awscc_events_endpoint`
* **New Data Source:** `awscc_events_endpoints`
* **New Data Source:** `awscc_lambda_url`
* **New Resource:** `awscc_datasync_location_fsx_open_zfs`
* **New Resource:** `awscc_events_endpoint`
* **New Resource:** `awscc_lambda_url`

## [0.16.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.16.0) (March 24, 2022)

FEATURES:

* **New Data Source:** `awscc_iotevents_alarm_model`
* **New Data Source:** `awscc_iotevents_alarm_models`
* **New Resource:** `awscc_iotevents_alarm_model`

## [0.15.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.15.0) (March 17, 2022)

FEATURES:

* **New Data Source:** `awscc_billingconductor_billing_group`
* **New Data Source:** `awscc_billingconductor_custom_line_item`
* **New Data Source:** `awscc_billingconductor_pricing_plan`
* **New Data Source:** `awscc_billingconductor_pricing_rule`
* **New Data Source:** `awscc_inspector_assessment_target`
* **New Data Source:** `awscc_inspector_assessment_targets`
* **New Data Source:** `awscc_inspector_assessment_template`
* **New Data Source:** `awscc_inspector_assessment_templates`
* **New Data Source:** `awscc_inspector_resource_group`
* **New Resource:** `awscc_billingconductor_billing_group`
* **New Resource:** `awscc_billingconductor_custom_line_item`
* **New Resource:** `awscc_billingconductor_pricing_plan`
* **New Resource:** `awscc_billingconductor_pricing_rule`
* **New Resource:** `awscc_inspector_assessment_target`
* **New Resource:** `awscc_inspector_assessment_template`
* **New Resource:** `awscc_inspector_resource_group`

## [0.14.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.14.0) (March 11, 2022)

FEATURES:

* **New Data Source:** `awscc_eks_identity_provider_config`
* **New Data Source:** `awscc_msk_batch_scram_secret`
* **New Data Source:** `awscc_msk_cluster`
* **New Data Source:** `awscc_msk_clusters`
* **New Data Source:** `awscc_msk_configuration`
* **New Data Source:** `awscc_msk_configurations`
* **New Data Source:** `awscc_personalize_dataset`
* **New Data Source:** `awscc_personalize_dataset_group`
* **New Data Source:** `awscc_personalize_dataset_groups`
* **New Data Source:** `awscc_personalize_datasets`
* **New Data Source:** `awscc_personalize_schema`
* **New Data Source:** `awscc_personalize_schemas`
* **New Data Source:** `awscc_personalize_solution`
* **New Data Source:** `awscc_personalize_solutions`
* **New Resource:** `awscc_eks_identity_provider_config`
* **New Resource:** `awscc_msk_batch_scram_secret`
* **New Resource:** `awscc_msk_cluster`
* **New Resource:** `awscc_msk_configuration`
* **New Resource:** `awscc_personalize_dataset`
* **New Resource:** `awscc_personalize_dataset_group`
* **New Resource:** `awscc_personalize_schema`
* **New Resource:** `awscc_personalize_solution`

## [0.13.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.13.0) (February 24, 2022)

FEATURES:

* **New Data Source:** `awscc_datasync_location_fsx_lustre`
* **New Data Source:** `awscc_datasync_location_fsx_lustres`
* **New Resource:** `awscc_datasync_location_fsx_lustre`
* Support property `pattern` validation ([#88](https://github.com/hashicorp/terraform-provider-awscc/issues/88))

## [0.12.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.12.0) (February 17, 2022)

BREAKING CHANGES:

* data-source/awscc_cloudfront_distribution: `distribution_config.cnam_es` attribute renamed to `distribution_config.cnames`
* resource/awscc_cloudfront_distribution: `distribution_config.cnam_es` attribute renamed to `distribution_config.cnames`

FEATURES:

* **New Data Source:** `awscc_amplify_apps.md`
* **New Data Source:** `awscc_apprunner_vpc_connector.md`
* **New Data Source:** `awscc_apprunner_vpc_connectors.md`
* **New Data Source:** `awscc_batch_compute_environment.md`
* **New Data Source:** `awscc_batch_compute_environments.md`
* **New Data Source:** `awscc_batch_job_queue.md`
* **New Data Source:** `awscc_batch_job_queues.md`
* **New Data Source:** `awscc_cloudformation_hook_default_version.md`
* **New Data Source:** `awscc_cloudformation_hook_type_config.md`
* **New Data Source:** `awscc_cloudformation_hook_version.md`
* **New Data Source:** `awscc_ecr_pull_through_cache_rule.md`
* **New Data Source:** `awscc_ecr_pull_through_cache_rules.md`
* **New Data Source:** `awscc_eks_nodegroup.md`
* **New Data Source:** `awscc_frauddetector_event_types.md`
* **New Data Source:** `awscc_gamelift_game_server_groups.md`
* **New Data Source:** `awscc_robomaker_robot_application.md`
* **New Data Source:** `awscc_robomaker_robot_applications.md`
* **New Data Source:** `awscc_ses_configuration_set_event_destination.md`
* **New Data Source:** `awscc_ses_template.md`
* **New Data Source:** `awscc_ses_templates.md`
* **New Data Source:** `awscc_sqs_queue.md`
* **New Data Source:** `awscc_sqs_queues.md`
* **New Resource:** `awscc_apprunner_vpc_connector.md`
* **New Resource:** `awscc_batch_compute_environment.md`
* **New Resource:** `awscc_batch_job_queue.md`
* **New Resource:** `awscc_cloudformation_hook_default_version.md`
* **New Resource:** `awscc_cloudformation_hook_type_config.md`
* **New Resource:** `awscc_cloudformation_hook_version.md`
* **New Resource:** `awscc_ecr_pull_through_cache_rule.md`
* **New Resource:** `awscc_eks_nodegroup.md`
* **New Resource:** `awscc_robomaker_robot_application.md`
* **New Resource:** `awscc_ses_configuration_set_event_destination.md`
* **New Resource:** `awscc_ses_template.md`
* **New Resource:** `awscc_sqs_queue.md`

## [0.11.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.11.0) (January 27, 2022)

BREAKING CHANGES:

* data-source/awscc_frauddetector_event_types: Temporarily removed

FEATURES:

* **New Data Source:** `awscc_appintegrations_data_integration`
* **New Data Source:** `awscc_appintegrations_data_integrations`
* **New Data Source:** `awscc_athena_data_catalogs`
* **New Data Source:** `awscc_cloudformation_module_default_versions`
* **New Data Source:** `awscc_iotcoredeviceadvisor_suite_definitions`
* **New Data Source:** `awscc_forecast_dataset`
* **New Data Source:** `awscc_forecast_datasets`
* **New Data Source:** `awscc_kafkaconnect_connector`
* **New Data Source:** `awscc_kafkaconnect_connectors`
* **New Data Source:** `awscc_lightsail_certificate`
* **New Data Source:** `awscc_lightsail_certificates`
* **New Data Source:** `awscc_lightsail_container`
* **New Data Source:** `awscc_lightsail_containers`
* **New Data Source:** `awscc_lightsail_distribution`
* **New Data Source:** `awscc_lightsail_distributions`
* **New Data Source:** `awscc_rekognition_collection`
* **New Data Source:** `awscc_rekognition_collections`
* **New Data Source:** `awscc_route53_dnssecs`
* **New Data Source:** `awscc_route53_key_signing_keys`
* **New Data Source:** `awscc_route53recoveryreadiness_readiness_checks`
* **New Data Source:** `awscc_servicecatalogappregistry_applications`
* **New Data Source:** `awscc_servicecatalogappregistry_attribute_groups`
* **New Resource:** `awscc_appintegrations_data_integration`
* **New Resource:** `awscc_forecast_dataset`
* **New Resource:** `awscc_kafkaconnect_connector`
* **New Resource:** `awscc_lightsail_certificate`
* **New Resource:** `awscc_lightsail_container`
* **New Resource:** `awscc_lightsail_distribution`
* **New Resource:** `awscc_rekognition_collection`

BUG FIXES:

* Prevent errors like `planned value ... for a non-computed attribute` for list arguments during `terraform plan` ([#368](https://github.com/hashicorp/terraform-provider-awscc/issues/368))

## [0.10.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.10.0) (January 13, 2022)

FEATURES:

* **New Data Source:** `awscc_appstream_application_entitlement_association`
* **New Data Source:** `awscc_appstream_entitlement`
* **New Data Source:** `awscc_ec2_host`
* **New Data Source:** `awscc_ec2_hosts`
* **New Data Source:** `awscc_ec2_network_insights_access_scope`
* **New Data Source:** `awscc_ec2_network_insights_access_scope_analyses`
* **New Data Source:** `awscc_ec2_network_insights_access_scope_analysis`
* **New Data Source:** `awscc_ec2_network_insights_access_scopes`
* **New Data Source:** `awscc_forecast_dataset_group`
* **New Data Source:** `awscc_forecast_dataset_groups`
* **New Data Source:** `awscc_inspectorv2_filter`
* **New Data Source:** `awscc_inspectorv2_filters`
* **New Data Source:** `awscc_kinesisvideo_signaling_channel`
* **New Data Source:** `awscc_kinesisvideo_stream`
* **New Data Source:** `awscc_lightsail_alarm`
* **New Data Source:** `awscc_lightsail_alarms`
* **New Data Source:** `awscc_lightsail_bucket`
* **New Data Source:** `awscc_lightsail_buckets`
* **New Data Source:** `awscc_lightsail_load_balancer`
* **New Data Source:** `awscc_lightsail_load_balancer_tls_certificate`
* **New Data Source:** `awscc_lightsail_load_balancers`
* **New Data Source:** `awscc_route53resolver_resolver_rule_association`
* **New Data Source:** `awscc_route53resolver_resolver_rule_associations`
* **New Resource:** `awscc_appstream_application_entitlement_association`
* **New Resource:** `awscc_appstream_entitlement`
* **New Resource:** `awscc_ec2_host`
* **New Resource:** `awscc_ec2_network_insights_access_scope`
* **New Resource:** `awscc_ec2_network_insights_access_scope_analysis`
* **New Resource:** `awscc_forecast_dataset_group`
* **New Resource:** `awscc_inspectorv2_filter`
* **New Resource:** `awscc_kinesisvideo_signaling_channel`
* **New Resource:** `awscc_kinesisvideo_stream`
* **New Resource:** `awscc_lightsail_alarm`
* **New Resource:** `awscc_lightsail_bucket`
* **New Resource:** `awscc_lightsail_load_balancer`
* **New Resource:** `awscc_lightsail_load_balancer_tls_certificate`
* **New Resource:** `awscc_route53resolver_resolver_rule_association`

BUG FIXES:

* Prevent errors like `Unable to set Terraform State Unknown values from Cloud Control API Properties.` during `terraform apply` ([#331](https://github.com/hashicorp/terraform-provider-awscc/issues/331))

## [0.9.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.9.0) (December 16, 2021)

FEATURES:

* **New Data Source:** `awscc_apigateway_deployment`
* **New Data Source:** `awscc_appsync_domain_name`
* **New Data Source:** `awscc_appsync_domain_name_api_association`
* **New Data Source:** `awscc_lex_bot`
* **New Data Source:** `awscc_lex_bot_alias`
* **New Data Source:** `awscc_lex_bot_version`
* **New Data Source:** `awscc_lex_bots`
* **New Data Source:** `awscc_lex_resource_policy`
* **New Resource:** `awscc_apigateway_deployment`
* **New Resource:** `awscc_appsync_domain_name`
* **New Resource:** `awscc_appsync_domain_name_api_association`
* **New Resource:** `awscc_lex_bot`
* **New Resource:** `awscc_lex_bot_alias`
* **New Resource:** `awscc_lex_bot_version`
* **New Resource:** `awscc_lex_resource_policy`

BUG FIXES:

* provider: Ensure `darwin/arm64` platform is included in releases
* Prevent `terraform plan` showing that a resource must be replaced immediately after creation ([#306](https://github.com/hashicorp/terraform-provider-awscc/issues/306))
* Prevent errors like `An unexpected error was encountered trying to read an attribute from the state.` during resource read ([#306](https://github.com/hashicorp/terraform-provider-awscc/issues/306))

## [0.8.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.8.0) (December  3, 2021)

FEATURES:

* **New Data Source:** `awscc_chatbot_slack_channel_configuration`
* **New Data Source:** `awscc_chatbot_slack_channel_configurations`
* **New Data Source:** `awscc_connect_contact_flow`
* **New Data Source:** `awscc_connect_contact_flow_module`
* **New Data Source:** `awscc_ec2_ipam`
* **New Data Source:** `awscc_ec2_ipam_allocation`
* **New Data Source:** `awscc_ec2_ipam_pool`
* **New Data Source:** `awscc_ec2_ipam_pools`
* **New Data Source:** `awscc_ec2_ipam_scope`
* **New Data Source:** `awscc_ec2_ipam_scopes`
* **New Data Source:** `awscc_ec2_ipams`
* **New Data Source:** `awscc_ec2_vpc_endpoint`
* **New Data Source:** `awscc_ec2_vpc_endpoints`
* **New Data Source:** `awscc_evidently_experiment`
* **New Data Source:** `awscc_evidently_feature`
* **New Data Source:** `awscc_evidently_launch`
* **New Data Source:** `awscc_evidently_project`
* **New Data Source:** `awscc_refactorspaces_environment`
* **New Data Source:** `awscc_refactorspaces_environments`
* **New Data Source:** `awscc_refactorspaces_route`
* **New Data Source:** `awscc_refactorspaces_service`
* **New Data Source:** `awscc_resiliencehub_app`
* **New Data Source:** `awscc_resiliencehub_apps`
* **New Data Source:** `awscc_resiliencehub_resiliency_policies`
* **New Data Source:** `awscc_resiliencehub_resiliency_policy`
* **New Data Source:** `awscc_rum_app_monitor`
* **New Data Source:** `awscc_rum_app_monitors`
* **New Data Source:** `awscc_timestream_scheduled_queries`
* **New Data Source:** `awscc_timestream_scheduled_query`
* **New Data Source:** `awscc_transfer_workflow`
* **New Data Source:** `awscc_transfer_workflows`
* **New Resource:** `awscc_chatbot_slack_channel_configuration`
* **New Resource:** `awscc_connect_contact_flow`
* **New Resource:** `awscc_connect_contact_flow_module`
* **New Resource:** `awscc_ec2_ipam`
* **New Resource:** `awscc_ec2_ipam_allocation`
* **New Resource:** `awscc_ec2_ipam_pool`
* **New Resource:** `awscc_ec2_ipam_scope`
* **New Resource:** `awscc_ec2_vpc_endpoint`
* **New Resource:** `awscc_evidently_experiment`
* **New Resource:** `awscc_evidently_feature`
* **New Resource:** `awscc_evidently_launch`
* **New Resource:** `awscc_evidently_project`
* **New Resource:** `awscc_refactorspaces_environment`
* **New Resource:** `awscc_refactorspaces_route`
* **New Resource:** `awscc_refactorspaces_service`
* **New Resource:** `awscc_resiliencehub_app`
* **New Resource:** `awscc_resiliencehub_resiliency_policy`
* **New Resource:** `awscc_rum_app_monitor`
* **New Resource:** `awscc_timestream_scheduled_query`
* **New Resource:** `awscc_transfer_workflow`

## [0.7.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.7.0) (November 18, 2021)

FEATURES:

* **New Data Source:** `awscc_appstream_app_block`
* **New Data Source:** `awscc_appstream_application`
* **New Data Source:** `awscc_appstream_application_fleet_association`
* **New Data Source:** `awscc_databrew_ruleset`
* **New Data Source:** `awscc_databrew_rulesets`
* **New Resource:** `awscc_appstream_app_block`
* **New Resource:** `awscc_appstream_application`
* **New Resource:** `awscc_appstream_application_fleet_association`
* **New Resource:** `awscc_databrew_ruleset`

## [0.6.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.6.0) (November 11, 2021)

FEATURES:

* **New Data Source:** `awscc_batch_scheduling_policies`
* **New Data Source:** `awscc_batch_scheduling_policy`
* **New Data Source:** `awscc_datasync_location_hdfs`
* **New Data Source:** `awscc_datasync_location_hdfs_plural`
* **New Data Source:** `awscc_iotwireless_fuota_task`
* **New Data Source:** `awscc_iotwireless_fuota_tasks`
* **New Data Source:** `awscc_iotwireless_multicast_group`
* **New Data Source:** `awscc_iotwireless_multicast_groups`
* **New Resource:** `awscc_batch_scheduling_policy`
* **New Resource:** `awscc_datasync_location_hdfs`
* **New Resource:** `awscc_iotwireless_fuota_task`
* **New Resource:** `awscc_iotwireless_multicast_group`

## [0.5.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.5.0) (November 4, 2021)

FEATURES:

* **New Data Source:** `awscc_cloudfront_response_headers_policies`
* **New Data Source:** `awscc_cloudfront_response_headers_policy`
* **New Data Source:** `awscc_ec2_capacity_reservation_fleet`
* **New Data Source:** `awscc_ec2_capacity_reservation_fleets`
* **New Data Source:** `awscc_ec2_internet_gateway`
* **New Data Source:** `awscc_ec2_internet_gateways`
* **New Data Source:** `awscc_ec2_network_interface`
* **New Data Source:** `awscc_ec2_network_interfaces`
* **New Data Source:** `awscc_ec2_subnet`
* **New Data Source:** `awscc_eks_cluster`
* **New Data Source:** `awscc_eks_clusters`
* **New Data Source:** `awscc_iot_logging`
* **New Data Source:** `awscc_iot_loggings`
* **New Data Source:** `awscc_iot_resource_specific_logging`
* **New Data Source:** `awscc_iot_resource_specific_loggings`
* **New Data Source:** `awscc_lightsail_database`
* **New Data Source:** `awscc_lightsail_databases`
* **New Data Source:** `awscc_lightsail_static_ip`
* **New Data Source:** `awscc_lightsail_static_ips`
* **New Data Source:** `awscc_pinpoint_in_app_template`
* **New Data Source:** `awscc_pinpoint_in_app_templates`
* **New Data Source:** `awscc_redshift_endpoint_access`
* **New Data Source:** `awscc_redshift_endpoint_accesses`
* **New Data Source:** `awscc_redshift_endpoint_authorization`
* **New Data Source:** `awscc_redshift_endpoint_authorizations`
* **New Data Source:** `awscc_redshift_event_subscription`
* **New Data Source:** `awscc_redshift_event_subscriptions`
* **New Data Source:** `awscc_redshift_scheduled_action`
* **New Data Source:** `awscc_redshift_scheduled_actions`
* **New Data Source:** `awscc_route53resolver_resolver_config`
* **New Data Source:** `awscc_route53resolver_resolver_configs`
* **New Data Source:** `awscc_route53resolver_resolver_rule`
* **New Data Source:** `awscc_route53resolver_resolver_rules`
* **New Resource:** `awscc_cloudfront_response_headers_policy`
* **New Resource:** `awscc_ec2_capacity_reservation_fleet`
* **New Resource:** `awscc_ec2_internet_gateway`
* **New Resource:** `awscc_ec2_network_interface`
* **New Resource:** `awscc_ec2_subnet`
* **New Resource:** `awscc_eks_cluster`
* **New Resource:** `awscc_iot_logging`
* **New Resource:** `awscc_iot_resource_specific_logging`
* **New Resource:** `awscc_lightsail_database`
* **New Resource:** `awscc_lightsail_static_ip`
* **New Resource:** `awscc_pinpoint_in_app_template`
* **New Resource:** `awscc_redshift_endpoint_access`
* **New Resource:** `awscc_redshift_endpoint_authorization`
* **New Resource:** `awscc_redshift_event_subscription`
* **New Resource:** `awscc_redshift_scheduled_action`
* **New Resource:** `awscc_route53resolver_resolver_config`
* **New Resource:** `awscc_route53resolver_resolver_rule`
* provider: Adds `user_agent` argument ([#247](https://github.com/hashicorp/terraform-provider-awscc/issues/247))

## [0.4.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.4.0) (October 21, 2021)

BREAKING CHANGES:

* data-source/awscc_ec2_subnet: Temporarily removed
* resource/awscc_ec2_subnet: Temporarily removed

FEATURES:

* **New Data Source:** `awscc_autoscaling_lifecycle_hook`
* **New Data Source:** `awscc_connect_hours_of_operation`
* **New Data Source:** `awscc_connect_user`
* **New Data Source:** `awscc_connect_user_hierarchy_group`
* **New Data Source:** `awscc_panorama_application_instance`
* **New Data Source:** `awscc_panorama_application_instances`
* **New Data Source:** `awscc_panorama_package`
* **New Data Source:** `awscc_panorama_packages`
* **New Data Source:** `awscc_panorama_package_version`
* **New Data Source:** `awscc_rekognition_project`
* **New Data Source:** `awscc_rekognition_projects`
* **New Data Source:** `awscc_s3outposts_bucket`
* **New Resource:** `awscc_autoscaling_lifecycle_hook`
* **New Resource:** `awscc_connect_hours_of_operation`
* **New Resource:** `awscc_connect_user`
* **New Resource:** `awscc_connect_user_hierarchy_group`
* **New Resource:** `awscc_panorama_application_instance`
* **New Resource:** `awscc_panorama_package`
* **New Resource:** `awscc_panorama_package_version`
* **New Resource:** `awscc_rekognition_project`
* **New Resource:** `awscc_s3outposts_bucket`

BUG FIXES:

* Persist any resource `id` to state if Create fails while waiting for async operation completion ([#252](https://github.com/hashicorp/terraform-provider-awscc/issues/252))
* data-source/awscc_s3_storagelens: Fix incorrectly generated schema ([#255](https://github.com/hashicorp/terraform-provider-awscc/issues/255))
* data-source/awscc_s3objectlambda_access_point: Fix incorrectly generated schema ([#255](https://github.com/hashicorp/terraform-provider-awscc/issues/255))
* data-source/awscc_sagemaker_pipeline: Fix incorrectly generated schema ([#255](https://github.com/hashicorp/terraform-provider-awscc/issues/255))
* resource/awscc_s3_storagelens: Fix incorrectly generated schema ([#255](https://github.com/hashicorp/terraform-provider-awscc/issues/255))
* resource/awscc_s3objectlambda_access_point: Fix incorrectly generated schema ([#255](https://github.com/hashicorp/terraform-provider-awscc/issues/255))
* resource/awscc_sagemaker_pipeline: Fix incorrectly generated schema ([#255](https://github.com/hashicorp/terraform-provider-awscc/issues/255))

## [0.3.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.3.0) (October 14, 2021)

BREAKING CHANGES:

* data-source/awscc_ec2_host: Temporarily removed
* data-source/awscc_ec2_hosts: Temporarily removed
* resource/awscc_ec2_host: Temporarily removed

FEATURES:

* **New Data Source:** `awscc_acmpca_certificate_authorities`
* **New Data Source:** `awscc_backup_framework`
* **New Data Source:** `awscc_backup_frameworks`
* **New Data Source:** `awscc_ce_anomaly_monitor`
* **New Data Source:** `awscc_ce_anomaly_monitors`
* **New Data Source:** `awscc_ce_anomaly_subscription`
* **New Data Source:** `awscc_ce_anomaly_subscriptions`
* **New Data Source:** `awscc_codestarnotifications_notification_rules`
* **New Data Source:** `awscc_cur_report_definition`
* **New Data Source:** `awscc_cur_report_definitions`
* **New Data Source:** `awscc_ec2_subnet`
* **New Data Source:** `awscc_ecr_public_repository`
* **New Data Source:** `awscc_ecr_public_repositories`
* **New Data Source:** `awscc_globalaccelerator_accelerators`
* **New Data Source:** `awscc_iot_job_template`
* **New Data Source:** `awscc_iot_job_templates`
* **New Data Source:** `awscc_iotwireless_partner_account`
* **New Data Source:** `awscc_iotwireless_partner_accounts`
* **New Data Source:** `awscc_kms_keys`
* **New Data Source:** `awscc_lambda_event_source_mappings`
* **New Data Source:** `awscc_lightsail_instance`
* **New Data Source:** `awscc_lightsail_instances`
* **New Data Source:** `awscc_lookoutequipment_inference_scheduler`
* **New Data Source:** `awscc_lookoutequipment_inference_schedulers`
* **New Data Source:** `awscc_memorydb_acl`
* **New Data Source:** `awscc_memorydb_acls`
* **New Data Source:** `awscc_memorydb_cluster`
* **New Data Source:** `awscc_memorydb_clusters`
* **New Data Source:** `awscc_memorydb_parameter_group`
* **New Data Source:** `awscc_memorydb_parameter_groups`
* **New Data Source:** `awscc_memorydb_subnet_group`
* **New Data Source:** `awscc_memorydb_subnet_groups`
* **New Data Source:** `awscc_memorydb_user`
* **New Data Source:** `awscc_memorydb_users`
* **New Data Source:** `awscc_wisdom_assistant`
* **New Data Source:** `awscc_wisdom_assistants`
* **New Data Source:** `awscc_wisdom_assistant_association`
* **New Data Source:** `awscc_wisdom_knowledge_base`
* **New Data Source:** `awscc_wisdom_knowledge_bases`
* **New Resource:** `awscc_backup_framework`
* **New Resource:** `awscc_ce_anomaly_monitor`
* **New Resource:** `awscc_ce_anomaly_subscription`
* **New Resource:** `awscc_cur_report_definition`
* **New Resource:** `awscc_ec2_subnet`
* **New Resource:** `awscc_ecr_public_repository`
* **New Resource:** `awscc_iot_job_template`
* **New Resource:** `awscc_iotwireless_partner_account`
* **New Resource:** `awscc_lightsail_instance`
* **New Resource:** `awscc_lookoutequipment_inference_scheduler`
* **New Resource:** `awscc_memorydb_acl`
* **New Resource:** `awscc_memorydb_cluster`
* **New Resource:** `awscc_memorydb_parameter_group`
* **New Resource:** `awscc_memorydb_subnet_group`
* **New Resource:** `awscc_memorydb_user`
* **New Resource:** `awscc_wisdom_assistant`
* **New Resource:** `awscc_wisdom_assistant_association`
* **New Resource:** `awscc_wisdom_knowledge_base`

## [0.2.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.2.0) (October 7, 2021)

BREAKING CHANGES:

* data-source/awscc_ec2_subnet: Temporarily removed
* data-source/awscc_ec2_subnets: Temporarily removed
* data-source/awscc_eks_cluster: Temporarily removed
* data-source/awscc_eks_clusters: Temporarily removed
* resource/awscc_ec2_subnet: Temporarily removed
* resource/awscc_eks_cluster: Temporarily removed

FEATURES:

* **New Data Source:** `awscc_apigateway_authorizer`
* **New Data Source:** `awscc_autoscaling_launch_configuration`
* **New Data Source:** `awscc_autoscaling_launch_configurations`
* **New Data Source:** `awscc_backup_report_plan`
* **New Data Source:** `awscc_backup_report_plans`
* **New Data Source:** `awscc_ec2_host`
* **New Data Source:** `awscc_ec2_hosts`
* **New Data Source:** `awscc_ec2_route_table`
* **New Data Source:** `awscc_ec2_subnet_network_acl_association`
* **New Data Source:** `awscc_ec2_subnet_network_acl_associations`
* **New Data Source:** `awscc_iam_role`
* **New Data Source:** `awscc_iam_roles`
* **New Data Source:** `awscc_lightsail_disk`
* **New Data Source:** `awscc_lightsail_disks`
* **New Resource:** `awscc_apigateway_authorizer`
* **New Resource:** `awscc_autoscaling_launch_configuration`
* **New Resource:** `awscc_backup_report_plan`
* **New Resource:** `awscc_ec2_host`
* **New Resource:** `awscc_ec2_subnet_network_acl_association`
* **New Resource:** `awscc_iam_role`
* **New Resource:** `awscc_lightsail_disk`

BUG FIXES:

* Correctly create resource with CloudFormation schema defined top-level `id` attribute ([#229](https://github.com/hashicorp/terraform-provider-awscc/issues/229))
* Don't attempt to perform Terraform attribute to CloudFormation property name mapping for map keys ([#232](https://github.com/hashicorp/terraform-provider-awscc/issues/232))

## [0.1.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.1.0) (September 30, 2021)

Public Technical Preview.

## [0.0.15](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.15) (September 30, 2021)

11 additional CloudFormation resource schemas.

## [0.0.14](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.14) (September 29, 2021)

Populate any Unknown values after resource update.
Set APN HTTP User-Agent header value.

## [0.0.13](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.13) (September 24, 2021)

Validate array entries in List/Set of primitive types.

## [0.0.12](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.12) (September 23, 2021)

Refresh CloudFormation resource type schemas.
Add `max_retries`, `assume_role.policy` and `assume_role.policy_arns` provider configuration options.

## [0.0.11](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.11) (September 21, 2021)

Add resource import support.
Validate `NestedAttribute` array lengths.

## [0.0.10](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.10) (September 18, 2021)

Migrate to `cloudcontrol` as the AWS service package.

## [0.0.9](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.9) (September 15, 2021)

Support attribute default values.
Support CloudFormation multisets.
Use `terraform-plugin-framework` Set implementation.

## [0.0.8](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.8) (September 10, 2021)

Documentation generation.
Support `ForceNew` attributes.

## [0.0.7](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.7) (September 9, 2021)

Singular data sources added.

## [0.0.6](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.6) (September 7, 2021)

Plural data sources added.

## [0.0.5](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.5) (August 31, 2021)

Attribute validation support added.
Additional AWS authnentication methods supported.

## [0.0.4](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.4) (August 19, 2021)

Provider renamed to `terraform-provider-awscc`.

## [0.0.3](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.3) (August 16, 2021)

Attempt to generate Terraform resources for all available CloudFormation resource types.
Simple resources (no nested attributes) working with [Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework) v0.2.0+ (forked version).

## [0.0.2](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.2) (July 14, 2021)

`aws_logs_log_group` working with [Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework) v0.1.0.

## [0.0.1](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.0.1) (June 7, 2021)

`aws_logs_log_group` working with [Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk) v2.
