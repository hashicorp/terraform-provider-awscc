## 0.74.0 (Unreleased)

BREAKING CHANGES:

* resource/awscc_ssmguiconnect_preferences: Temporarily removed

## 0.73.0 (April  4, 2024)

FEATURES:

* **New Data Source:** `awscc_apigateway_base_path_mapping`
* **New Data Source:** `awscc_apigateway_base_path_mappings`
* **New Data Source:** `awscc_appintegrations_application`
* **New Data Source:** `awscc_appintegrations_applications`
* **New Data Source:** `awscc_applicationautoscaling_scalable_target`
* **New Data Source:** `awscc_applicationautoscaling_scalable_targets`
* **New Data Source:** `awscc_aps_scraper`
* **New Data Source:** `awscc_aps_scrapers`
* **New Data Source:** `awscc_bedrock_agent_alias`
* **New Data Source:** `awscc_bedrock_agent`
* **New Data Source:** `awscc_bedrock_agents`
* **New Data Source:** `awscc_bedrock_data_source`
* **New Data Source:** `awscc_bedrock_knowledge_base`
* **New Data Source:** `awscc_bedrock_knowledge_bases`
* **New Data Source:** `awscc_cleanroomsml_training_dataset`
* **New Data Source:** `awscc_cleanroomsml_training_datasets`
* **New Data Source:** `awscc_cloudfront_key_value_store`
* **New Data Source:** `awscc_cloudfront_key_value_stores`
* **New Data Source:** `awscc_codeartifact_package_group`
* **New Data Source:** `awscc_codeconnections_connection`
* **New Data Source:** `awscc_codeconnections_connections`
* **New Data Source:** `awscc_codepipeline_custom_action_type`
* **New Data Source:** `awscc_codepipeline_custom_action_types`
* **New Data Source:** `awscc_connect_instance`
* **New Data Source:** `awscc_connect_instances`
* **New Data Source:** `awscc_datazone_environment`
* **New Data Source:** `awscc_datazone_environments`
* **New Data Source:** `awscc_datazone_subscription_target`
* **New Data Source:** `awscc_datazone_subscription_targets`
* **New Data Source:** `awscc_deadline_farm`
* **New Data Source:** `awscc_deadline_farms`
* **New Data Source:** `awscc_deadline_fleet`
* **New Data Source:** `awscc_deadline_license_endpoint`
* **New Data Source:** `awscc_deadline_license_endpoints`
* **New Data Source:** `awscc_deadline_metered_product`
* **New Data Source:** `awscc_deadline_metered_products`
* **New Data Source:** `awscc_deadline_queue_environment`
* **New Data Source:** `awscc_deadline_queue_fleet_association`
* **New Data Source:** `awscc_deadline_queue`
* **New Data Source:** `awscc_deadline_storage_profile`
* **New Data Source:** `awscc_ec2_transit_gateway_route_table_association`
* **New Data Source:** `awscc_ec2_vpcdhcp_options_association`
* **New Data Source:** `awscc_ec2_vpcdhcp_options_associations`
* **New Data Source:** `awscc_entityresolution_id_namespace`
* **New Data Source:** `awscc_entityresolution_id_namespaces`
* **New Data Source:** `awscc_entityresolution_policy_statement`
* **New Data Source:** `awscc_globalaccelerator_cross_account_attachment`
* **New Data Source:** `awscc_globalaccelerator_cross_account_attachments`
* **New Data Source:** `awscc_groundstation_config`
* **New Data Source:** `awscc_groundstation_configs`
* **New Data Source:** `awscc_iot_billing_group`
* **New Data Source:** `awscc_iot_billing_groups`
* **New Data Source:** `awscc_iot_thing_group`
* **New Data Source:** `awscc_iot_thing_groups`
* **New Data Source:** `awscc_iot_thing_type`
* **New Data Source:** `awscc_iot_thing_types`
* **New Data Source:** `awscc_iot_thing`
* **New Data Source:** `awscc_iot_things`
* **New Data Source:** `awscc_iotanalytics_channel`
* **New Data Source:** `awscc_iotanalytics_channels`
* **New Data Source:** `awscc_iotanalytics_dataset`
* **New Data Source:** `awscc_iotanalytics_datasets`
* **New Data Source:** `awscc_iotanalytics_datastore`
* **New Data Source:** `awscc_iotanalytics_datastores`
* **New Data Source:** `awscc_iotanalytics_pipeline`
* **New Data Source:** `awscc_iotanalytics_pipelines`
* **New Data Source:** `awscc_ivs_encoder_configuration`
* **New Data Source:** `awscc_ivs_encoder_configurations`
* **New Data Source:** `awscc_ivs_playback_restriction_policies`
* **New Data Source:** `awscc_ivs_playback_restriction_policy`
* **New Data Source:** `awscc_ivs_storage_configuration`
* **New Data Source:** `awscc_ivs_storage_configurations`
* **New Data Source:** `awscc_ivschat_logging_configuration`
* **New Data Source:** `awscc_ivschat_logging_configurations`
* **New Data Source:** `awscc_ivschat_room`
* **New Data Source:** `awscc_ivschat_rooms`
* **New Data Source:** `awscc_kafkaconnect_custom_plugin`
* **New Data Source:** `awscc_kafkaconnect_custom_plugins`
* **New Data Source:** `awscc_kafkaconnect_worker_configuration`
* **New Data Source:** `awscc_kafkaconnect_worker_configurations`
* **New Data Source:** `awscc_networkmanager_link`
* **New Data Source:** `awscc_networkmanager_links`
* **New Data Source:** `awscc_omics_variant_store`
* **New Data Source:** `awscc_omics_variant_stores`
* **New Data Source:** `awscc_opensearchservice_domain`
* **New Data Source:** `awscc_opensearchservice_domains`
* **New Data Source:** `awscc_opsworkscm_server`
* **New Data Source:** `awscc_opsworkscm_servers`
* **New Data Source:** `awscc_proton_environment_account_connection`
* **New Data Source:** `awscc_proton_environment_account_connections`
* **New Data Source:** `awscc_redshift_cluster`
* **New Data Source:** `awscc_redshift_clusters`
* **New Data Source:** `awscc_s3outposts_endpoint`
* **New Data Source:** `awscc_s3outposts_endpoints`
* **New Data Source:** `awscc_securityhub_delegated_admin`
* **New Data Source:** `awscc_securityhub_insight`
* **New Data Source:** `awscc_securityhub_product_subscription`
* **New Data Source:** `awscc_securitylake_aws_log_source`
* **New Data Source:** `awscc_securitylake_data_lake`
* **New Data Source:** `awscc_securitylake_subscriber`
* **New Data Source:** `awscc_sns_topic`
* **New Data Source:** `awscc_sns_topics`
* **New Data Source:** `awscc_synthetics_canaries`
* **New Data Source:** `awscc_synthetics_canary`
* **New Data Source:** `awscc_synthetics_group`
* **New Data Source:** `awscc_synthetics_groups`
* **New Data Source:** `awscc_vpclattice_access_log_subscription`
* **New Data Source:** `awscc_vpclattice_access_log_subscriptions`
* **New Data Source:** `awscc_vpclattice_listener`
* **New Data Source:** `awscc_vpclattice_listeners`
* **New Data Source:** `awscc_vpclattice_rule`
* **New Data Source:** `awscc_vpclattice_rules`
* **New Data Source:** `awscc_vpclattice_service_network_service_association`
* **New Data Source:** `awscc_vpclattice_service_network_service_associations`
* **New Data Source:** `awscc_vpclattice_service_network_vpc_association`
* **New Data Source:** `awscc_vpclattice_service_network_vpc_associations`
* **New Data Source:** `awscc_vpclattice_service_network`
* **New Data Source:** `awscc_vpclattice_service_networks`
* **New Data Source:** `awscc_vpclattice_service`
* **New Data Source:** `awscc_vpclattice_services`
* **New Data Source:** `awscc_vpclattice_target_group`
* **New Data Source:** `awscc_vpclattice_target_groups`
* **New Resource:** `awscc_apigateway_base_path_mapping`
* **New Resource:** `awscc_appintegrations_application`
* **New Resource:** `awscc_applicationautoscaling_scalable_target`
* **New Resource:** `awscc_aps_scraper`
* **New Resource:** `awscc_bedrock_agent_alias`
* **New Resource:** `awscc_bedrock_agent`
* **New Resource:** `awscc_bedrock_data_source`
* **New Resource:** `awscc_bedrock_knowledge_base`
* **New Resource:** `awscc_cleanroomsml_training_dataset`
* **New Resource:** `awscc_cloudfront_key_value_store`
* **New Resource:** `awscc_codeartifact_package_group`
* **New Resource:** `awscc_codeconnections_connection`
* **New Resource:** `awscc_codepipeline_custom_action_type`
* **New Resource:** `awscc_connect_instance`
* **New Resource:** `awscc_datazone_environment`
* **New Resource:** `awscc_datazone_subscription_target`
* **New Resource:** `awscc_deadline_farm`
* **New Resource:** `awscc_deadline_fleet`
* **New Resource:** `awscc_deadline_license_endpoint`
* **New Resource:** `awscc_deadline_metered_product`
* **New Resource:** `awscc_deadline_queue_environment`
* **New Resource:** `awscc_deadline_queue_fleet_association`
* **New Resource:** `awscc_deadline_queue`
* **New Resource:** `awscc_deadline_storage_profile`
* **New Resource:** `awscc_ec2_transit_gateway_route_table_association`
* **New Resource:** `awscc_ec2_vpcdhcp_options_association`
* **New Resource:** `awscc_entityresolution_id_namespace`
* **New Resource:** `awscc_entityresolution_policy_statement`
* **New Resource:** `awscc_globalaccelerator_cross_account_attachment`
* **New Resource:** `awscc_groundstation_config`
* **New Resource:** `awscc_iot_billing_group`
* **New Resource:** `awscc_iot_thing_group`
* **New Resource:** `awscc_iot_thing_type`
* **New Resource:** `awscc_iot_thing`
* **New Resource:** `awscc_iotanalytics_channel`
* **New Resource:** `awscc_iotanalytics_dataset`
* **New Resource:** `awscc_iotanalytics_datastore`
* **New Resource:** `awscc_iotanalytics_pipeline`
* **New Resource:** `awscc_ivs_encoder_configuration`
* **New Resource:** `awscc_ivs_playback_restriction_policy`
* **New Resource:** `awscc_ivs_storage_configuration`
* **New Resource:** `awscc_ivschat_logging_configuration`
* **New Resource:** `awscc_ivschat_room`
* **New Resource:** `awscc_kafkaconnect_custom_plugin`
* **New Resource:** `awscc_kafkaconnect_worker_configuration`
* **New Resource:** `awscc_networkmanager_link`
* **New Resource:** `awscc_omics_variant_store`
* **New Resource:** `awscc_opensearchservice_domain`
* **New Resource:** `awscc_opsworkscm_server`
* **New Resource:** `awscc_proton_environment_account_connection`
* **New Resource:** `awscc_redshift_cluster`
* **New Resource:** `awscc_s3outposts_endpoint`
* **New Resource:** `awscc_securityhub_delegated_admin`
* **New Resource:** `awscc_securityhub_insight`
* **New Resource:** `awscc_securityhub_product_subscription`
* **New Resource:** `awscc_securitylake_aws_log_source`
* **New Resource:** `awscc_securitylake_data_lake`
* **New Resource:** `awscc_securitylake_subscriber`
* **New Resource:** `awscc_sns_topic`
* **New Resource:** `awscc_synthetics_canary`
* **New Resource:** `awscc_synthetics_group`
* **New Resource:** `awscc_vpclattice_access_log_subscription`
* **New Resource:** `awscc_vpclattice_listener`
* **New Resource:** `awscc_vpclattice_rule`
* **New Resource:** `awscc_vpclattice_service_network_service_association`
* **New Resource:** `awscc_vpclattice_service_network_vpc_association`
* **New Resource:** `awscc_vpclattice_service_network`
* **New Resource:** `awscc_vpclattice_service`
* **New Resource:** `awscc_vpclattice_target_group`

BUG FIXES:

* Fix `ValidationException: Identifier ... is not valid for identifier` errors ([#1501](https://github.com/hashicorp/terraform-provider-awscc/issues/1501))

## 0.72.1 (March 22, 2024)

BUG FIXES:

* Fix `Unable to Convert Configuration` errors ([#1490](https://github.com/hashicorp/terraform-provider-awscc/issues/1490))

## 0.72.0 (March 21, 2024)

FEATURES:

* **New Data Source:** `awscc_ec2_vpc_cidr_block`
* **New Resource:** `awscc_ec2_vpc_cidr_block`

## 0.71.0 (February 22, 2024)

FEATURES:

* **New Data Source:** `awscc_cognito_user_pool_risk_configuration_attachment`
* **New Data Source:** `awscc_ec2_security_group_ingress`
* **New Data Source:** `awscc_ec2_security_group_ingresses`
* **New Data Source:** `awscc_guardduty_master`
* **New Data Source:** `awscc_guardduty_member`
* **New Resource:** `awscc_cognito_user_pool_risk_configuration_attachment`
* **New Resource:** `awscc_ec2_security_group_ingress`
* **New Resource:** `awscc_guardduty_master`
* **New Resource:** `awscc_guardduty_member`

## 0.70.0 (February  8, 2024)

FEATURES:

* **New Data Source:** `awscc_appconfig_environment`
* **New Data Source:** `awscc_appconfig_hosted_configuration_version`
* **New Data Source:** `awscc_rds_integration`
* **New Data Source:** `awscc_rds_integrations`
* **New Resource:** `awscc_appconfig_environment`
* **New Resource:** `awscc_appconfig_hosted_configuration_version`
* **New Resource:** `awscc_rds_integration`

## 0.69.0 (January 25, 2024)

FEATURES:

* **New Data Source:** `awscc_batch_job_definition`
* **New Data Source:** `awscc_codebuild_fleet`
* **New Data Source:** `awscc_codebuild_fleets`
* **New Data Source:** `awscc_cognito_identity_pool`
* **New Data Source:** `awscc_cognito_identity_pools`
* **New Data Source:** `awscc_datazone_data_source`
* **New Data Source:** `awscc_datazone_domain`
* **New Data Source:** `awscc_datazone_domains`
* **New Data Source:** `awscc_datazone_environment_blueprint_configuration`
* **New Data Source:** `awscc_datazone_environment_profile`
* **New Data Source:** `awscc_datazone_project`
* **New Data Source:** `awscc_guardduty_filter`
* **New Data Source:** `awscc_ivs_stage`
* **New Data Source:** `awscc_ivs_stages`
* **New Data Source:** `awscc_ssmguiconnect_preferences`
* **New Resource:** `awscc_batch_job_definition`
* **New Resource:** `awscc_codebuild_fleet`
* **New Resource:** `awscc_cognito_identity_pool`
* **New Resource:** `awscc_datazone_data_source`
* **New Resource:** `awscc_datazone_domain`
* **New Resource:** `awscc_datazone_environment_blueprint_configuration`
* **New Resource:** `awscc_datazone_environment_profile`
* **New Resource:** `awscc_datazone_project`
* **New Resource:** `awscc_guardduty_filter`
* **New Resource:** `awscc_ivs_stage`
* **New Resource:** `awscc_ssmguiconnect_preferences`

BUG FIXES:

* Fix `Provider returned invalid result object after apply` errors ([#1363](https://github.com/hashicorp/terraform-provider-awscc/issues/1363))

## 0.68.0 (January 11, 2024)

FEATURES:

* **New Data Source:** `awscc_b2bi_partnership`
* **New Data Source:** `awscc_b2bi_partnerships`
* **New Data Source:** `awscc_connect_predefined_attribute`
* **New Data Source:** `awscc_ec2_instance`
* **New Data Source:** `awscc_ec2_security_group`
* **New Data Source:** `awscc_eks_access_entry`
* **New Data Source:** `awscc_emr_wal_workspace`
* **New Data Source:** `awscc_emr_wal_workspaces`
* **New Data Source:** `awscc_eventschemas_schema`
* **New Data Source:** `awscc_iot_certificate_provider`
* **New Data Source:** `awscc_iot_certificate_providers`
* **New Data Source:** `awscc_location_api_key`
* **New Data Source:** `awscc_location_api_keys`
* **New Data Source:** `awscc_neptunegraph_graph`
* **New Data Source:** `awscc_neptunegraph_graphs`
* **New Data Source:** `awscc_neptunegraph_private_graph_endpoint`
* **New Data Source:** `awscc_networkfirewall_tls_inspection_configuration`
* **New Data Source:** `awscc_networkfirewall_tls_inspection_configurations`
* **New Data Source:** `awscc_ssm_patch_baseline`
* **New Data Source:** `awscc_ssm_patch_baselines`
* **New Resource:** `awscc_b2bi_partnership`
* **New Resource:** `awscc_connect_predefined_attribute`
* **New Resource:** `awscc_ec2_instance`
* **New Resource:** `awscc_ec2_security_group`
* **New Resource:** `awscc_eks_access_entry`
* **New Resource:** `awscc_emr_wal_workspace`
* **New Resource:** `awscc_eventschemas_schema`
* **New Resource:** `awscc_iot_certificate_provider`
* **New Resource:** `awscc_location_api_key`
* **New Resource:** `awscc_neptunegraph_graph`
* **New Resource:** `awscc_neptunegraph_private_graph_endpoint`
* **New Resource:** `awscc_networkfirewall_tls_inspection_configuration`
* **New Resource:** `awscc_ssm_patch_baseline`

## 0.67.0 (December 14, 2023)

FEATURES:

* provider: Add `https_proxy` and `no_proxy` arguments
* **New Data Source:** `awscc_arczonalshift_zonal_autoshift_configuration`
* **New Data Source:** `awscc_arczonalshift_zonal_autoshift_configurations`
* **New Data Source:** `awscc_b2bi_capabilities`
* **New Data Source:** `awscc_b2bi_capability`
* **New Data Source:** `awscc_b2bi_profile`
* **New Data Source:** `awscc_b2bi_profiles`
* **New Data Source:** `awscc_b2bi_transformer`
* **New Data Source:** `awscc_b2bi_transformers`
* **New Data Source:** `awscc_dms_data_provider`
* **New Data Source:** `awscc_dms_data_providers`
* **New Data Source:** `awscc_dms_instance_profile`
* **New Data Source:** `awscc_dms_instance_profiles`
* **New Data Source:** `awscc_dms_migration_project`
* **New Data Source:** `awscc_dms_migration_projects`
* **New Data Source:** `awscc_ec2_snapshot_block_public_access`
* **New Data Source:** `awscc_ec2_snapshot_block_public_accesses`
* **New Data Source:** `awscc_eventschemas_discoverer`
* **New Data Source:** `awscc_eventschemas_discoverers`
* **New Data Source:** `awscc_eventschemas_registries`
* **New Data Source:** `awscc_eventschemas_registry`
* **New Data Source:** `awscc_fis_target_account_configuration`
* **New Data Source:** `awscc_lambda_event_invoke_config`
* **New Data Source:** `awscc_securityhub_hub`
* **New Data Source:** `awscc_securityhub_hubs`
* **New Data Source:** `awscc_workspacesthinclient_environment`
* **New Data Source:** `awscc_workspacesthinclient_environments`
* **New Resource:** `awscc_arczonalshift_zonal_autoshift_configuration`
* **New Resource:** `awscc_b2bi_capability`
* **New Resource:** `awscc_b2bi_profile`
* **New Resource:** `awscc_b2bi_transformer`
* **New Resource:** `awscc_dms_data_provider`
* **New Resource:** `awscc_dms_instance_profile`
* **New Resource:** `awscc_dms_migration_project`
* **New Resource:** `awscc_ec2_snapshot_block_public_access`
* **New Resource:** `awscc_eventschemas_discoverer`
* **New Resource:** `awscc_eventschemas_registry`
* **New Resource:** `awscc_fis_target_account_configuration`
* **New Resource:** `awscc_lambda_event_invoke_config`
* **New Resource:** `awscc_securityhub_hub`
* **New Resource:** `awscc_workspacesthinclient_environment`

## 0.66.0 (November 30, 2023)

FEATURES:

* **New Data Source:** `awscc_backup_restore_testing_plan`
* **New Data Source:** `awscc_backup_restore_testing_plans`
* **New Data Source:** `awscc_backup_restore_testing_selection`
* **New Data Source:** `awscc_backup_restore_testing_selections`
* **New Data Source:** `awscc_codestarconnections_repository_link`
* **New Data Source:** `awscc_codestarconnections_repository_links`
* **New Data Source:** `awscc_codestarconnections_sync_configuration`
* **New Data Source:** `awscc_codestarconnections_sync_configurations`
* **New Data Source:** `awscc_eks_pod_identity_association`
* **New Data Source:** `awscc_elasticache_serverless_cache`
* **New Data Source:** `awscc_elasticache_serverless_caches`
* **New Data Source:** `awscc_elasticloadbalancingv2_trust_store`
* **New Data Source:** `awscc_elasticloadbalancingv2_trust_store_revocation`
* **New Data Source:** `awscc_guardduty_ip_set`
* **New Data Source:** `awscc_imagebuilder_lifecycle_policies`
* **New Data Source:** `awscc_imagebuilder_lifecycle_policy`
* **New Data Source:** `awscc_logs_deliveries`
* **New Data Source:** `awscc_logs_delivery`
* **New Data Source:** `awscc_logs_delivery_destination`
* **New Data Source:** `awscc_logs_delivery_destinations`
* **New Data Source:** `awscc_logs_delivery_source`
* **New Data Source:** `awscc_logs_delivery_sources`
* **New Data Source:** `awscc_logs_log_anomaly_detector`
* **New Data Source:** `awscc_logs_log_anomaly_detectors`
* **New Data Source:** `awscc_opensearchserverless_lifecycle_policy`
* **New Data Source:** `awscc_s3_access_grant`
* **New Data Source:** `awscc_s3_access_grants_instance`
* **New Data Source:** `awscc_s3_access_grants_instances`
* **New Data Source:** `awscc_s3_access_grants_location`
* **New Data Source:** `awscc_s3express_bucket_policies`
* **New Data Source:** `awscc_s3express_bucket_policy`
* **New Data Source:** `awscc_s3express_directory_bucket`
* **New Data Source:** `awscc_s3express_directory_buckets`
* **New Data Source:** `awscc_sagemaker_inference_component`
* **New Data Source:** `awscc_sagemaker_inference_components`
* **New Resource:** `awscc_backup_restore_testing_plan`
* **New Resource:** `awscc_backup_restore_testing_selection`
* **New Resource:** `awscc_codestarconnections_repository_link`
* **New Resource:** `awscc_codestarconnections_sync_configuration`
* **New Resource:** `awscc_eks_pod_identity_association`
* **New Resource:** `awscc_elasticache_serverless_cache`
* **New Resource:** `awscc_elasticloadbalancingv2_trust_store`
* **New Resource:** `awscc_elasticloadbalancingv2_trust_store_revocation`
* **New Resource:** `awscc_guardduty_ip_set`
* **New Resource:** `awscc_imagebuilder_lifecycle_policy`
* **New Resource:** `awscc_logs_delivery`
* **New Resource:** `awscc_logs_delivery_destination`
* **New Resource:** `awscc_logs_delivery_source`
* **New Resource:** `awscc_logs_log_anomaly_detector`
* **New Resource:** `awscc_opensearchserverless_lifecycle_policy`
* **New Resource:** `awscc_s3_access_grant`
* **New Resource:** `awscc_s3_access_grants_instance`
* **New Resource:** `awscc_s3_access_grants_location`
* **New Resource:** `awscc_s3express_bucket_policy`
* **New Resource:** `awscc_s3express_directory_bucket`
* **New Resource:** `awscc_sagemaker_inference_component`

## 0.65.0 (November 16, 2023)

FEATURES:

* **New Data Source:** `awscc_applicationautoscaling_scaling_policy`
* **New Data Source:** `awscc_ec2_security_group_egress`
* **New Data Source:** `awscc_ec2_security_group_egresses`
* **New Data Source:** `awscc_gamelift_game_session_queue`
* **New Data Source:** `awscc_gamelift_game_session_queues`
* **New Data Source:** `awscc_gamelift_matchmaking_configuration`
* **New Data Source:** `awscc_gamelift_matchmaking_configurations`
* **New Data Source:** `awscc_gamelift_matchmaking_rule_set`
* **New Data Source:** `awscc_gamelift_matchmaking_rule_sets`
* **New Data Source:** `awscc_gamelift_script`
* **New Data Source:** `awscc_gamelift_scripts`
* **New Data Source:** `awscc_guardduty_threat_intel_set`
* **New Data Source:** `awscc_s3_storage_lens_group`
* **New Data Source:** `awscc_s3_storage_lens_groups`
* **New Resource:** `awscc_applicationautoscaling_scaling_policy`
* **New Resource:** `awscc_ec2_security_group_egress`
* **New Resource:** `awscc_gamelift_game_session_queue`
* **New Resource:** `awscc_gamelift_matchmaking_configuration`
* **New Resource:** `awscc_gamelift_matchmaking_rule_set`
* **New Resource:** `awscc_gamelift_script`
* **New Resource:** `awscc_guardduty_threat_intel_set`
* **New Resource:** `awscc_s3_storage_lens_group`

## 0.64.0 (November  9, 2023)

FEATURES:

* **New Data Source:** `awscc_appconfig_configuration_profile`
* **New Data Source:** `awscc_cognito_user_pool_user_to_group_attachment`
* **New Data Source:** `awscc_ec2_vpc_gateway_attachment`
* **New Data Source:** `awscc_ec2_vpc_gateway_attachments`
* **New Data Source:** `awscc_events_rule`
* **New Data Source:** `awscc_events_rules`
* **New Data Source:** `awscc_iam_user`
* **New Data Source:** `awscc_iam_users`
* **New Data Source:** `awscc_medialive_multiplex`
* **New Data Source:** `awscc_medialive_multiplexes`
* **New Data Source:** `awscc_medialive_multiplexprogram`
* **New Data Source:** `awscc_route53resolver_firewall_domain_lists`
* **New Resource:** `awscc_appconfig_configuration_profile`
* **New Resource:** `awscc_cognito_user_pool_user_to_group_attachment`
* **New Resource:** `awscc_ec2_vpc_gateway_attachment`
* **New Resource:** `awscc_events_rule`
* **New Resource:** `awscc_iam_user`
* **New Resource:** `awscc_medialive_multiplex`
* **New Resource:** `awscc_medialive_multiplexprogram`

BREAKING CHANGES:

* resource/awscc_apigatewayv2_api: `body` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_apigatewayv2_integration_response: `response_parameters` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_apigatewayv2_integration_response: `response_templates` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_apigatewayv2_model: `schema` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_apigatewayv2_route: `request_models` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_apigatewayv2_route: `request_parameters` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_backup_backup_plan: `backup_plan.advance_backup_settings.backup_options` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_codeartifact_domain: `permissions_policy_document` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_codeartifact_repository: `permissions_policy_document` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_cognito_identity_pool_principal_tag: `principal_tags` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_connect_view: `template` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_dms_replication_config: `replication_settings` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_dms_replication_config: `supplemental_settings` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_dms_replication_config: `table_mappings` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_dynamodb_table: `key_schema` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_ecr_registry_policy: `policy_text` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_efs_file_system: `file_system_policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_eventschemas_registry_policy: `policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_events_archive: `event_pattern` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_iam_group_policy: `policy_document` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_iam_role_policy: `policy_document` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_iam_user_policy: `policy_document` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_lakeformation_data_cells_filter: `all_rows_wildcard` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_lakeformation_principal_permissions: `catalog` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_lakeformation_principal_permissions: `table_wildcard` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_lakeformation_tag_association: `catalog` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_lakeformation_tag_association: `table_wildcard` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_lambda_function: `policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_lex_resource_policy: `policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_logs_log_group: `data_protection_policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_msk_cluster_policy: `policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_mwaa_environment: `airflow_configuration_options` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_mwaa_environment: `tags` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_memorydb_parameter_group: `parameters` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_oam_sink: `policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_pinpoint_in_app_template: `custom_config` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_pinpoint_in_app_template: `tags` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_ram_permission: `policy_template` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_rds_db_cluster_parameter_group: `parameters` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_rds_db_parameter_group: `parameters` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_s3_access_point: `policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_s3_multi_region_access_point_policy: `policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_s3_storage_lens: `storage_lens_configuration.data_export.s3_bucket_destination.encryption.sses3` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_s3objectlambda_access_point_policy: `policy_document` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_s3outposts_access_point: `policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_s3outposts_bucket_policy: `policy_document` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_sns_topic_inline_policy: `policy_document` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_sqs_queue_inline_policy: `policy_document` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_servicecatalogappregistry_attribute_group: `attributes` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_shield_protection: `block` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_shield_protection: `count` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_stepfunctions_state_machine: `definition` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_vpclattice_auth_policy: `policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_vpclattice_resource_policy: `policy` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_wafv2_logging_configuration: `redacted_fields.match_pattern.all` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_wafv2_logging_configuration: `redacted_fields.method` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_wafv2_logging_configuration: `redacted_fields.query_string` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* resource/awscc_wafv2_logging_configuration: `redacted_fields.uri_path` attribute's type has changed from  `Map of String` to `String` and must be valid JSON
* data-source/awscc_apigatewayv2_api: `body` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_apigatewayv2_integration_response: `response_parameters` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_apigatewayv2_integration_response: `response_templates` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_apigatewayv2_model: `schema` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_apigatewayv2_route: `request_models` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_apigatewayv2_route: `request_parameters` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_backup_backup_plan: `backup_plan.advance_backup_settings.backup_options` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_codeartifact_domain: `permissions_policy_document` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_codeartifact_repository: `permissions_policy_document` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_cognito_identity_pool_principal_tag: `principal_tags` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_connect_view: `template` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_dms_replication_config: `replication_settings` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_dms_replication_config: `supplemental_settings` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_dms_replication_config: `table_mappings` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_dynamodb_table: `key_schema` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_ecr_registry_policy: `policy_text` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_efs_file_system: `file_system_policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_eventschemas_registry_policy: `policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_events_archive: `event_pattern` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_iam_group_policy: `policy_document` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_iam_role_policy: `policy_document` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_iam_user_policy: `policy_document` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_lakeformation_data_cells_filter: `all_rows_wildcard` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_lakeformation_principal_permissions: `catalog` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_lakeformation_principal_permissions: `table_wildcard` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_lakeformation_tag_association: `catalog` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_lakeformation_tag_association: `table_wildcard` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_lambda_function: `policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_lex_resource_policy: `policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_logs_log_group: `data_protection_policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_msk_cluster_policy: `policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_mwaa_environment: `airflow_configuration_options` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_mwaa_environment: `tags` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_memorydb_parameter_group: `parameters` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_oam_sink: `policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_pinpoint_in_app_template: `custom_config` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_pinpoint_in_app_template: `tags` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_ram_permission: `policy_template` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_rds_db_cluster_parameter_group: `parameters` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_rds_db_parameter_group: `parameters` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_s3_access_point: `policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_s3_multi_region_access_point_policy: `policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_s3_storage_lens: `storage_lens_configuration.data_export.s3_bucket_destination.encryption.sses3` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_s3objectlambda_access_point_policy: `policy_document` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_s3outposts_access_point: `policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_s3outposts_bucket_policy: `policy_document` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_sns_topic_inline_policy: `policy_document` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_sqs_queue_inline_policy: `policy_document` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_servicecatalogappregistry_attribute_group: `attributes` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_shield_protection: `block` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_shield_protection: `count` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_stepfunctions_state_machine: `definition` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_vpclattice_auth_policy: `policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_vpclattice_resource_policy: `policy` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_wafv2_logging_configuration: `redacted_fields.match_pattern.all` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_wafv2_logging_configuration: `redacted_fields.method` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_wafv2_logging_configuration: `redacted_fields.query_string` attribute's type has changed from  `Map of String` to `String`
* data-source/awscc_wafv2_logging_configuration: `redacted_fields.uri_path` attribute's type has changed from  `Map of String` to `String`

## 0.63.0 (October 26, 2023)

FEATURES:

* **New Data Source:** `awscc_appconfig_application`
* **New Data Source:** `awscc_appconfig_applications`
* **New Data Source:** `awscc_appsync_function_configuration`
* **New Data Source:** `awscc_appsync_resolver`
* **New Data Source:** `awscc_autoscaling_auto_scaling_group`
* **New Data Source:** `awscc_autoscaling_auto_scaling_groups`
* **New Data Source:** `awscc_cognito_log_delivery_configuration`
* **New Data Source:** `awscc_cognito_user_pool_client`
* **New Data Source:** `awscc_entityresolution_id_mapping_workflow`
* **New Data Source:** `awscc_entityresolution_id_mapping_workflows`
* **New Data Source:** `awscc_events_event_bus`
* **New Data Source:** `awscc_events_event_buses`
* **New Data Source:** `awscc_iam_group`
* **New Data Source:** `awscc_iam_groups`
* **New Data Source:** `awscc_lambda_version`
* **New Data Source:** `awscc_msk_replicator`
* **New Data Source:** `awscc_msk_replicators`
* **New Data Source:** `awscc_route53resolver_firewall_rule_group_associations`
* **New Data Source:** `awscc_route53resolver_firewall_rule_groups`
* **New Resource:** `awscc_appconfig_application`
* **New Resource:** `awscc_appsync_function_configuration`
* **New Resource:** `awscc_appsync_resolver`
* **New Resource:** `awscc_autoscaling_auto_scaling_group`
* **New Resource:** `awscc_cognito_log_delivery_configuration`
* **New Resource:** `awscc_cognito_user_pool_client`
* **New Resource:** `awscc_entityresolution_id_mapping_workflow`
* **New Resource:** `awscc_events_event_bus`
* **New Resource:** `awscc_iam_group`
* **New Resource:** `awscc_lambda_version`
* **New Resource:** `awscc_msk_replicator`

## 0.62.0 (October 12, 2023)

FEATURES:

* **New Data Source:** `awscc_apigatewayv2_domain_name`
* **New Data Source:** `awscc_apigatewayv2_domain_names`
* **New Data Source:** `awscc_appsync_domain_names`
* **New Data Source:** `awscc_cognito_user_pool_user`
* **New Data Source:** `awscc_connect_security_profile`
* **New Data Source:** `awscc_elasticloadbalancingv2_load_balancer`
* **New Data Source:** `awscc_elasticloadbalancingv2_load_balancers`
* **New Data Source:** `awscc_healthimaging_datastore`
* **New Data Source:** `awscc_healthimaging_datastores`
* **New Data Source:** `awscc_iot_software_package`
* **New Data Source:** `awscc_iot_software_package_version`
* **New Data Source:** `awscc_iot_software_packages`
* **New Data Source:** `awscc_s3_bucket_policies`
* **New Data Source:** `awscc_s3_bucket_policy`
* **New Data Source:** `awscc_ssm_parameter`
* **New Resource:** `awscc_apigatewayv2_domain_name`
* **New Resource:** `awscc_cognito_user_pool_user`
* **New Resource:** `awscc_connect_security_profile`
* **New Resource:** `awscc_elasticloadbalancingv2_load_balancer`
* **New Resource:** `awscc_healthimaging_datastore`
* **New Resource:** `awscc_iot_software_package`
* **New Resource:** `awscc_iot_software_package_version`
* **New Resource:** `awscc_s3_bucket_policy`
* **New Resource:** `awscc_ssm_parameter`

## 0.61.0 (September 21, 2023)

FEATURES:

* **New Data Source:** `awscc_ec2_transit_gateway_route_table`
* **New Data Source:** `awscc_ec2_transit_gateway_route_tables`
* **New Data Source:** `awscc_ec2_vpc_endpoint_connection_notification`
* **New Data Source:** `awscc_ec2_vpc_endpoint_connection_notifications`
* **New Data Source:** `awscc_mediapackagev2_channel`
* **New Data Source:** `awscc_mediapackagev2_channel_group`
* **New Data Source:** `awscc_mediapackagev2_channel_groups`
* **New Data Source:** `awscc_mediapackagev2_channel_policy`
* **New Data Source:** `awscc_mediapackagev2_origin_endpoint`
* **New Data Source:** `awscc_mediapackagev2_origin_endpoint_policy`
* **New Data Source:** `awscc_sagemaker_data_quality_job_definitions`
* **New Data Source:** `awscc_sagemaker_model_bias_job_definitions`
* **New Data Source:** `awscc_sagemaker_model_explainability_job_definitions`
* **New Data Source:** `awscc_sagemaker_model_quality_job_definitions`
* **New Resource:** `awscc_ec2_transit_gateway_route_table`
* **New Resource:** `awscc_ec2_vpc_endpoint_connection_notification`
* **New Resource:** `awscc_mediapackagev2_channel`
* **New Resource:** `awscc_mediapackagev2_channel_group`
* **New Resource:** `awscc_mediapackagev2_channel_policy`
* **New Resource:** `awscc_mediapackagev2_origin_endpoint`
* **New Resource:** `awscc_mediapackagev2_origin_endpoint_policy`

## 0.60.0 (September  7, 2023)

FEATURES:

* **New Data Source:** `awscc_chatbot_microsoft_teams_channel_configurations`
* **New Data Source:** `awscc_cleanrooms_analysis_template`
* **New Data Source:** `awscc_cleanrooms_analysis_templates`
* **New Data Source:** `awscc_cloudformation_hook_versions`
* **New Data Source:** `awscc_connect_view`
* **New Data Source:** `awscc_connect_view_version`
* **New Data Source:** `awscc_ec2_eip_association`
* **New Data Source:** `awscc_ec2_eip_associations`
* **New Data Source:** `awscc_guardduty_detector`
* **New Data Source:** `awscc_guardduty_detectors`
* **New Data Source:** `awscc_iotwireless_task_definitions`
* **New Data Source:** `awscc_pcaconnectorad_connector`
* **New Data Source:** `awscc_pcaconnectorad_connectors`
* **New Data Source:** `awscc_pcaconnectorad_directory_registration`
* **New Data Source:** `awscc_pcaconnectorad_directory_registrations`
* **New Data Source:** `awscc_pcaconnectorad_service_principal_name`
* **New Data Source:** `awscc_pcaconnectorad_template`
* **New Data Source:** `awscc_pcaconnectorad_template_group_access_control_entry`
* **New Resource:** `awscc_cleanrooms_analysis_template`
* **New Resource:** `awscc_connect_view`
* **New Resource:** `awscc_connect_view_version`
* **New Resource:** `awscc_ec2_eip_association`
* **New Resource:** `awscc_guardduty_detector`
* **New Resource:** `awscc_pcaconnectorad_connector`
* **New Resource:** `awscc_pcaconnectorad_directory_registration`
* **New Resource:** `awscc_pcaconnectorad_service_principal_name`
* **New Resource:** `awscc_pcaconnectorad_template`
* **New Resource:** `awscc_pcaconnectorad_template_group_access_control_entry`

## 0.59.0 (August 24, 2023)

NOTES:

* provider: Updates to Go 1.20, the last release that will run on any release of Windows 7, 8, Server 2008 and Server 2012. A future release will update to Go 1.21, and these platforms will no longer be supported
* provider: Updates to Go 1.20, the last release that will run on macOS 10.13 High Sierra or 10.14 Mojave. A future release will update to Go 1.21, and these platforms will no longer be supported
* provider: Updates to Go 1.20. The provider will now notice the `trust-ad` option in `/etc/resolv.conf` and, if set, will set the "authentic data" option in outgoing DNS requests in order to better match the behavior of the GNU libc resolver

## 0.58.0 (August 10, 2023)

FEATURES:

* **New Data Source:** `awscc_cloudwatch_alarm`
* **New Data Source:** `awscc_cloudwatch_alarms`
* **New Data Source:** `awscc_config_config_rule`
* **New Data Source:** `awscc_config_config_rules`
* **New Data Source:** `awscc_connect_traffic_distribution_group`
* **New Data Source:** `awscc_datasync_location_azure_blob`
* **New Data Source:** `awscc_datasync_location_azure_blobs`
* **New Data Source:** `awscc_dms_replication_config`
* **New Data Source:** `awscc_dms_replication_configs`
* **New Data Source:** `awscc_ec2_network_interface_attachment`
* **New Data Source:** `awscc_ec2_network_interface_attachments`
* **New Data Source:** `awscc_ec2_route`
* **New Data Source:** `awscc_lambda_layer_version_permission`
* **New Data Source:** `awscc_mediatailor_channel`
* **New Data Source:** `awscc_mediatailor_channel_policy`
* **New Data Source:** `awscc_mediatailor_channels`
* **New Data Source:** `awscc_mediatailor_live_source`
* **New Data Source:** `awscc_mediatailor_vod_source`
* **New Data Source:** `awscc_pipes_pipes`
* **New Data Source:** `awscc_sqs_queue_inline_policy`
* **New Resource:** `awscc_cloudwatch_alarm`
* **New Resource:** `awscc_config_config_rule`
* **New Resource:** `awscc_connect_traffic_distribution_group`
* **New Resource:** `awscc_datasync_location_azure_blob`
* **New Resource:** `awscc_dms_replication_config`
* **New Resource:** `awscc_ec2_network_interface_attachment`
* **New Resource:** `awscc_ec2_route`
* **New Resource:** `awscc_lambda_layer_version_permission`
* **New Resource:** `awscc_mediatailor_channel`
* **New Resource:** `awscc_mediatailor_channel_policy`
* **New Resource:** `awscc_mediatailor_live_source`
* **New Resource:** `awscc_mediatailor_vod_source`
* **New Resource:** `awscc_sqs_queue_inline_policy`

## 0.57.0 (July 27, 2023)

FEATURES:

* **New Data Source:** `awscc_iam_managed_policies`
* **New Data Source:** `awscc_iam_managed_policy`
* **New Data Source:** `awscc_lambda_permission`
* **New Data Source:** `awscc_sns_topic_inline_policy`
* **New Resource:** `awscc_iam_managed_policy`
* **New Resource:** `awscc_lambda_permission`
* **New Resource:** `awscc_sns_topic_inline_policy`

## 0.56.0 (July 20, 2023)

FEATURES:

* **New Data Source:** `awscc_apigatewayv2_api_mapping`
* **New Data Source:** `awscc_connect_queue`
* **New Data Source:** `awscc_connect_routing_profile`
* **New Data Source:** `awscc_iam_group_policy`
* **New Data Source:** `awscc_iam_role_policy`
* **New Data Source:** `awscc_iam_user_policy`
* **New Data Source:** `awscc_logs_account_policy`
* **New Data Source:** `awscc_sns_topic_policy`
* **New Resource:** `awscc_apigatewayv2_api_mapping`
* **New Resource:** `awscc_connect_queue`
* **New Resource:** `awscc_connect_routing_profile`
* **New Resource:** `awscc_iam_group_policy`
* **New Resource:** `awscc_iam_role_policy`
* **New Resource:** `awscc_iam_user_policy`
* **New Resource:** `awscc_logs_account_policy`
* **New Resource:** `awscc_sns_topic_policy`

## 0.55.0 (June 29, 2023)

FEATURES:

* **New Data Source:** `awscc_apprunner_auto_scaling_configuration`
* **New Data Source:** `awscc_apprunner_auto_scaling_configurations`
* **New Data Source:** `awscc_appstream_app_block_builder`
* **New Data Source:** `awscc_appsync_source_api_association`
* **New Data Source:** `awscc_cleanrooms_configured_table_association`
* **New Data Source:** `awscc_cognito_identity_pool_principal_tag`
* **New Data Source:** `awscc_comprehend_document_classifier`
* **New Data Source:** `awscc_comprehend_document_classifiers`
* **New Data Source:** `awscc_connect_prompt`
* **New Data Source:** `awscc_customerprofiles_calculated_attribute_definition`
* **New Data Source:** `awscc_customerprofiles_calculated_attribute_definitions`
* **New Data Source:** `awscc_customerprofiles_event_stream`
* **New Data Source:** `awscc_ec2_launch_template`
* **New Data Source:** `awscc_ec2_launch_templates`
* **New Data Source:** `awscc_ec2_volume_attachment`
* **New Data Source:** `awscc_ec2_volume_attachments`
* **New Data Source:** `awscc_iam_service_linked_role`
* **New Data Source:** `awscc_organizations_organization`
* **New Data Source:** `awscc_organizations_organizations`
* **New Data Source:** `awscc_stepfunctions_state_machine_alias`
* **New Data Source:** `awscc_stepfunctions_state_machine_version`
* **New Data Source:** `awscc_verifiedpermissions_identity_source`
* **New Data Source:** `awscc_verifiedpermissions_policy`
* **New Data Source:** `awscc_verifiedpermissions_policy_store`
* **New Data Source:** `awscc_verifiedpermissions_policy_stores`
* **New Data Source:** `awscc_verifiedpermissions_policy_template`
* **New Data Source:** `awscc_vpclattice_auth_policy`
* **New Data Source:** `awscc_vpclattice_resource_policy`
* **New Resource:** `awscc_apprunner_auto_scaling_configuration`
* **New Resource:** `awscc_appstream_app_block_builder`
* **New Resource:** `awscc_appsync_source_api_association`
* **New Resource:** `awscc_cleanrooms_configured_table_association`
* **New Resource:** `awscc_cognito_identity_pool_principal_tag`
* **New Resource:** `awscc_comprehend_document_classifier`
* **New Resource:** `awscc_connect_prompt`
* **New Resource:** `awscc_customerprofiles_calculated_attribute_definition`
* **New Resource:** `awscc_customerprofiles_event_stream`
* **New Resource:** `awscc_ec2_launch_template`
* **New Resource:** `awscc_ec2_volume_attachment`
* **New Resource:** `awscc_iam_service_linked_role`
* **New Resource:** `awscc_organizations_organization`
* **New Resource:** `awscc_stepfunctions_state_machine_alias`
* **New Resource:** `awscc_stepfunctions_state_machine_version`
* **New Resource:** `awscc_verifiedpermissions_identity_source`
* **New Resource:** `awscc_verifiedpermissions_policy`
* **New Resource:** `awscc_verifiedpermissions_policy_store`
* **New Resource:** `awscc_verifiedpermissions_policy_template`
* **New Resource:** `awscc_vpclattice_auth_policy`
* **New Resource:** `awscc_vpclattice_resource_policy`

## 0.54.0 (June 15, 2023)

FEATURES:

* **New Data Source:** `awscc_athena_capacity_reservation`
* **New Data Source:** `awscc_athena_capacity_reservations`
* **New Data Source:** `awscc_cleanrooms_collaboration`
* **New Data Source:** `awscc_cleanrooms_collaborations`
* **New Data Source:** `awscc_cleanrooms_configured_table`
* **New Data Source:** `awscc_cleanrooms_configured_tables`
* **New Data Source:** `awscc_cleanrooms_membership`
* **New Data Source:** `awscc_cleanrooms_memberships`
* **New Data Source:** `awscc_mediaconnect_bridge`
* **New Data Source:** `awscc_mediaconnect_bridge_output`
* **New Data Source:** `awscc_mediaconnect_bridge_source`
* **New Data Source:** `awscc_mediaconnect_bridges`
* **New Data Source:** `awscc_mediaconnect_gateway`
* **New Data Source:** `awscc_mediaconnect_gateways`
* **New Data Source:** `awscc_rds_custom_db_engine_version`
* **New Data Source:** `awscc_rds_custom_db_engine_versions`
* **New Data Source:** `awscc_securityhub_standard`
* **New Resource:** `awscc_athena_capacity_reservation`
* **New Resource:** `awscc_cleanrooms_collaboration`
* **New Resource:** `awscc_cleanrooms_configured_table`
* **New Resource:** `awscc_cleanrooms_membership`
* **New Resource:** `awscc_mediaconnect_bridge`
* **New Resource:** `awscc_mediaconnect_bridge_output`
* **New Resource:** `awscc_mediaconnect_bridge_source`
* **New Resource:** `awscc_mediaconnect_gateway`
* **New Resource:** `awscc_rds_custom_db_engine_version`
* **New Resource:** `awscc_securityhub_standard`

## 0.53.0 (June  1, 2023)

FEATURES:

* **New Data Source:** `awscc_detective_organization_admin`
* **New Data Source:** `awscc_ec2_subnet_cidr_block`
* **New Data Source:** `awscc_ec2_subnet_cidr_blocks`
* **New Data Source:** `awscc_quicksight_topic`
* **New Data Source:** `awscc_shield_drt_access`
* **New Data Source:** `awscc_shield_drt_accesses`
* **New Data Source:** `awscc_shield_proactive_engagement`
* **New Data Source:** `awscc_shield_proactive_engagements`
* **New Data Source:** `awscc_shield_protection`
* **New Data Source:** `awscc_shield_protection_group`
* **New Data Source:** `awscc_shield_protection_groups`
* **New Data Source:** `awscc_shield_protections`
* **New Resource:** `awscc_detective_organization_admin`
* **New Resource:** `awscc_ec2_subnet_cidr_block`
* **New Resource:** `awscc_quicksight_topic`
* **New Resource:** `awscc_shield_drt_access`
* **New Resource:** `awscc_shield_proactive_engagement`
* **New Resource:** `awscc_shield_protection`
* **New Resource:** `awscc_shield_protection_group`

## 0.52.0 (May 11, 2023)

FEATURES:

* **New Data Source:** `awscc_apigatewayv2_integration_response`
* **New Data Source:** `awscc_backupgateway_hypervisor`
* **New Data Source:** `awscc_backupgateway_hypervisors`
* **New Data Source:** `awscc_datasync_storage_system`
* **New Data Source:** `awscc_datasync_storage_systems`
* **New Data Source:** `awscc_ec2_verified_access_endpoint`
* **New Data Source:** `awscc_ec2_verified_access_endpoints`
* **New Data Source:** `awscc_ec2_verified_access_group`
* **New Data Source:** `awscc_ec2_verified_access_groups`
* **New Data Source:** `awscc_ec2_verified_access_instance`
* **New Data Source:** `awscc_ec2_verified_access_instances`
* **New Data Source:** `awscc_ec2_verified_access_trust_provider`
* **New Data Source:** `awscc_ec2_verified_access_trust_providers`
* **New Data Source:** `awscc_osis_pipeline`
* **New Data Source:** `awscc_osis_pipelines`
* **New Data Source:** `awscc_proton_environment_template`
* **New Data Source:** `awscc_proton_environment_templates`
* **New Data Source:** `awscc_proton_service_template`
* **New Data Source:** `awscc_proton_service_templates`
* **New Data Source:** `awscc_quicksight_vpc_connection`
* **New Data Source:** `awscc_secretsmanager_secret`
* **New Data Source:** `awscc_secretsmanager_secrets`
* **New Resource:** `awscc_apigatewayv2_integration_response`
* **New Resource:** `awscc_backupgateway_hypervisor`
* **New Resource:** `awscc_datasync_storage_system`
* **New Resource:** `awscc_ec2_verified_access_endpoint`
* **New Resource:** `awscc_ec2_verified_access_group`
* **New Resource:** `awscc_ec2_verified_access_instance`
* **New Resource:** `awscc_ec2_verified_access_trust_provider`
* **New Resource:** `awscc_osis_pipeline`
* **New Resource:** `awscc_proton_environment_template`
* **New Resource:** `awscc_proton_service_template`
* **New Resource:** `awscc_quicksight_vpc_connection`
* **New Resource:** `awscc_secretsmanager_secret`

## 0.51.0 (April 27, 2023)

FEATURES:

* **New Data Source:** `awscc_ec2_vpc_endpoint_service_permissions`
* **New Data Source:** `awscc_ec2_vpc_endpoint_service_permissions_plural`
* **New Data Source:** `awscc_frauddetector_list`
* **New Data Source:** `awscc_frauddetector_lists`
* **New Data Source:** `awscc_msk_cluster_policy`
* **New Data Source:** `awscc_msk_vpc_connection`
* **New Data Source:** `awscc_msk_vpc_connections`
* **New Data Source:** `awscc_ram_permission`
* **New Data Source:** `awscc_ram_permissions`
* **New Data Source:** `awscc_xray_sampling_rules`
* **New Resource:** `awscc_ec2_vpc_endpoint_service_permissions`
* **New Resource:** `awscc_frauddetector_list`
* **New Resource:** `awscc_msk_cluster_policy`
* **New Resource:** `awscc_msk_vpc_connection`
* **New Resource:** `awscc_ram_permission`

## 0.50.0 (April 13, 2023)

NOTES:

* provider: The `skip_medatadata_api_check` argument is being deprecated in favor of `skip_metadata_api_check` ([#908](https://github.com/hashicorp/terraform-provider-awscc/issues/908))

FEATURES:

* **New Data Source:** `awscc_appconfig_extension_association`
* **New Data Source:** `awscc_appconfig_extension_associations`
* **New Data Source:** `awscc_devopsguru_log_anomaly_detection_integration`
* **New Data Source:** `awscc_devopsguru_log_anomaly_detection_integrations`
* **New Data Source:** `awscc_iotwireless_wireless_device_import_task`
* **New Data Source:** `awscc_iotwireless_wireless_device_import_tasks`
* **New Data Source:** `awscc_neptune_db_cluster`
* **New Data Source:** `awscc_neptune_db_clusters`
* **New Data Source:** `awscc_quicksight_refresh_schedule`
* **New Data Source:** `awscc_ssmcontacts_plan`
* **New Data Source:** `awscc_ssmcontacts_rotation`
* **New Resource:** `awscc_appconfig_extension_association`
* **New Resource:** `awscc_devopsguru_log_anomaly_detection_integration`
* **New Resource:** `awscc_iotwireless_wireless_device_import_task`
* **New Resource:** `awscc_neptune_db_cluster`
* **New Resource:** `awscc_quicksight_refresh_schedule`
* **New Resource:** `awscc_ssmcontacts_plan`
* **New Resource:** `awscc_ssmcontacts_rotation`

## 0.49.0 (March 30, 2023)

FEATURES:

* **New Data Source:** `awscc_apigatewayv2_route`
* **New Data Source:** `awscc_backup_backup_plans`
* **New Data Source:** `awscc_backup_backup_selections`
* **New Data Source:** `awscc_backup_backup_vaults`
* **New Data Source:** `awscc_chatbot_microsoft_teams_channel_configuration`
* **New Data Source:** `awscc_comprehend_flywheel`
* **New Data Source:** `awscc_comprehend_flywheels`
* **New Data Source:** `awscc_devopsguru_resource_collections`
* **New Data Source:** `awscc_sagemaker_inference_experiment`
* **New Data Source:** `awscc_sagemaker_inference_experiments`
* **New Resource:** `awscc_apigatewayv2_route`
* **New Resource:** `awscc_chatbot_microsoft_teams_channel_configuration`
* **New Resource:** `awscc_comprehend_flywheel`
* **New Resource:** `awscc_sagemaker_inference_experiment`

## 0.48.0 (March  9, 2023)

FEATURES:

* **New Data Source:** `awscc_docdbelastic_cluster`
* **New Data Source:** `awscc_managedblockchain_accessor`
* **New Data Source:** `awscc_docdbelastic_clusters`
* **New Data Source:** `awscc_managedblockchain_accessors`
* **New Resource:** `awscc_docdbelastic_cluster`
* **New Resource:** `awscc_managedblockchain_accessor`

## 0.47.0 (February 22, 2023)

FEATURES:

* **New Data Source:** `awscc_apigateway_vpc_link`
* **New Data Source:** `awscc_apigateway_vpc_links`
* **New Data Source:** `awscc_ec2_local_gateway_route_table`
* **New Data Source:** `awscc_ec2_local_gateway_route_table_virtual_interface_group_association`
* **New Data Source:** `awscc_ec2_local_gateway_route_table_virtual_interface_group_associations`
* **New Data Source:** `awscc_ec2_local_gateway_route_tables`
* **New Data Source:** `awscc_ec2_vpc_endpoint_service`
* **New Data Source:** `awscc_ec2_vpc_endpoint_services`
* **New Data Source:** `awscc_ec2_vpn_connection_route`
* **New Data Source:** `awscc_ec2_vpn_connection_routes`
* **New Data Source:** `awscc_fms_resource_set`
* **New Data Source:** `awscc_internetmonitor_monitor`
* **New Data Source:** `awscc_internetmonitor_monitors`
* **New Data Source:** `awscc_networkmanager_transit_gateway_peering`
* **New Data Source:** `awscc_networkmanager_transit_gateway_peerings`
* **New Data Source:** `awscc_networkmanager_transit_gateway_route_table_attachment`
* **New Data Source:** `awscc_networkmanager_transit_gateway_route_table_attachments`
* **New Data Source:** `awscc_organizations_resource_policy`
* **New Data Source:** `awscc_systemsmanagersap_application`
* **New Data Source:** `awscc_systemsmanagersap_applications`
* **New Resource:** `awscc_apigateway_vpc_link`
* **New Resource:** `awscc_ec2_local_gateway_route_table`
* **New Resource:** `awscc_ec2_local_gateway_route_table_virtual_interface_group_association`
* **New Resource:** `awscc_ec2_vpc_endpoint_service`
* **New Resource:** `awscc_ec2_vpn_connection_route`
* **New Resource:** `awscc_fms_resource_set`
* **New Resource:** `awscc_internetmonitor_monitor`
* **New Resource:** `awscc_networkmanager_transit_gateway_peering`
* **New Resource:** `awscc_networkmanager_transit_gateway_route_table_attachment`
* **New Resource:** `awscc_organizations_resource_policy`
* **New Resource:** `awscc_systemsmanagersap_application`

## 0.46.0 (February  9, 2023)

FEATURES:

* **New Data Source:** `awscc_cloudtrail_channel`
* **New Data Source:** `awscc_cloudtrail_channels`
* **New Data Source:** `awscc_cloudtrail_resource_policy`
* **New Data Source:** `awscc_connect_approved_origin`
* **New Data Source:** `awscc_connect_integration_association`
* **New Data Source:** `awscc_connect_security_key`
* **New Data Source:** `awscc_ec2_ipam_pool_cidr`
* **New Data Source:** `awscc_ec2_ipam_resource_discoveries`
* **New Data Source:** `awscc_ec2_ipam_resource_discovery`
* **New Data Source:** `awscc_ec2_ipam_resource_discovery_association`
* **New Data Source:** `awscc_ec2_ipam_resource_discovery_associations`
* **New Data Source:** `awscc_omics_reference_store`
* **New Data Source:** `awscc_omics_reference_stores`
* **New Data Source:** `awscc_omics_run_group`
* **New Data Source:** `awscc_omics_run_groups`
* **New Data Source:** `awscc_omics_sequence_store`
* **New Data Source:** `awscc_omics_sequence_stores`
* **New Data Source:** `awscc_omics_workflow`
* **New Data Source:** `awscc_omics_workflows`
* **New Data Source:** `awscc_sagemaker_space`
* **New Data Source:** `awscc_sagemaker_spaces`
* **New Data Source:** `awscc_simspaceweaver_simulation`
* **New Data Source:** `awscc_simspaceweaver_simulations`
* **New Data Source:** `awscc_ssmcontacts_contacts`
* **New Resource:** `awscc_cloudtrail_channel`
* **New Resource:** `awscc_cloudtrail_resource_policy`
* **New Resource:** `awscc_connect_approved_origin`
* **New Resource:** `awscc_connect_integration_association`
* **New Resource:** `awscc_connect_security_key`
* **New Resource:** `awscc_ec2_ipam_pool_cidr`
* **New Resource:** `awscc_ec2_ipam_resource_discovery`
* **New Resource:** `awscc_ec2_ipam_resource_discovery_association`
* **New Resource:** `awscc_omics_reference_store`
* **New Resource:** `awscc_omics_run_group`
* **New Resource:** `awscc_omics_sequence_store`
* **New Resource:** `awscc_omics_workflow`
* **New Resource:** `awscc_sagemaker_space`
* **New Resource:** `awscc_simspaceweaver_simulation`

## 0.45.0 (January 19, 2023)

FEATURES:

* **New Data Source:** `awscc_kendraranking_execution_plan`
* **New Data Source:** `awscc_kendraranking_execution_plans`
* **New Resource:** `awscc_/kendraranking_execution_plan`

## 0.44.0 (January  5, 2023)

FEATURES:

* **New Data Source:** `awscc_directoryservice_simple_ad`
* **New Data Source:** `awscc_directoryservice_simple_ads`
* **New Data Source:** `awscc_elasticbeanstalk_configuration_template`
* **New Data Source:** `awscc_elasticbeanstalk_configuration_templates`
* **New Data Source:** `awscc_elasticloadbalancingv2_target_group`
* **New Data Source:** `awscc_elasticloadbalancingv2_target_groups`
* **New Data Source:** `awscc_gamelift_build`
* **New Data Source:** `awscc_gamelift_builds`
* **New Data Source:** `awscc_stepfunctions_state_machines`
* **New Resource:** `awscc_directoryservice_simple_ad`
* **New Resource:** `awscc_elasticbeanstalk_configuration_template`
* **New Resource:** `awscc_elasticloadbalancingv2_target_group`
* **New Resource:** `awscc_gamelift_build`

## [0.43.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.43.0) (December 15, 2022)

FEATURES:

* **New Data Source:** `awscc_grafana_workspace`
* **New Data Source:** `awscc_grafana_workspaces`
* **New Data Source:** `awscc_iottwinmaker_sync_job`
* **New Data Source:** `awscc_opensearchserverless_access_policies`
* **New Data Source:** `awscc_opensearchserverless_access_policy`
* **New Data Source:** `awscc_opensearchserverless_collection`
* **New Data Source:** `awscc_opensearchserverless_collections`
* **New Data Source:** `awscc_opensearchserverless_security_config`
* **New Data Source:** `awscc_opensearchserverless_security_configs`
* **New Data Source:** `awscc_opensearchserverless_security_policies`
* **New Data Source:** `awscc_opensearchserverless_security_policy`
* **New Data Source:** `awscc_opensearchserverless_vpc_endpoint`
* **New Data Source:** `awscc_opensearchserverless_vpc_endpoints`
* **New Resource:** `awscc_grafana_workspace`
* **New Resource:** `awscc_iottwinmaker_sync_job`
* **New Resource:** `awscc_opensearchserverless_access_policy`
* **New Resource:** `awscc_opensearchserverless_collection`
* **New Resource:** `awscc_opensearchserverless_security_config`
* **New Resource:** `awscc_opensearchserverless_security_policy`
* **New Resource:** `awscc_opensearchserverless_vpc_endpoint`

## [0.42.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.42.0) (December  8, 2022)

FEATURES:

* **New Data Source:** `awscc_apigateway_rest_api`
* **New Data Source:** `awscc_apigateway_rest_apis`
* **New Data Source:** `awscc_appflow_connector`
* **New Data Source:** `awscc_appflow_connectors`
* **New Data Source:** `awscc_ec2_network_performance_metric_subscription`
* **New Data Source:** `awscc_ec2_network_performance_metric_subscriptions`
* **New Data Source:** `awscc_pipes_pipe`
* **New Resource:** `awscc_apigateway_rest_api`
* **New Resource:** `awscc_appflow_connector`
* **New Resource:** `awscc_ec2_network_performance_metric_subscription`
* **New Resource:** `awscc_pipes_pipe`

## [0.41.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.41.0) (December  1, 2022)

FEATURES:

* **New Data Source:** `awscc_gamelift_location`
* **New Data Source:** `awscc_gamelift_locations`
* **New Data Source:** `awscc_oam_link`
* **New Data Source:** `awscc_oam_links`
* **New Data Source:** `awscc_oam_sink`
* **New Data Source:** `awscc_oam_sinks`
* **New Resource:** `awscc_gamelift_location`
* **New Resource:** `awscc_oam_link`
* **New Resource:** `awscc_oam_sink`

## [0.40.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.40.0) (November 23, 2022)

FEATURES:

* **New Data Source:** `awscc_apigateway_documentation_part`
* **New Data Source:** `awscc_apigatewayv2_api`
* **New Data Source:** `awscc_apigatewayv2_apis`
* **New Data Source:** `awscc_apigatewayv2_authorizer`
* **New Data Source:** `awscc_apigatewayv2_deployment`
* **New Data Source:** `awscc_apigatewayv2_model`
* **New Data Source:** `awscc_appstream_image_builder`
* **New Data Source:** `awscc_appstream_image_builders`
* **New Data Source:** `awscc_autoscaling_scaling_policies`
* **New Data Source:** `awscc_autoscaling_scaling_policy`
* **New Data Source:** `awscc_autoscaling_scheduled_action`
* **New Data Source:** `awscc_autoscaling_scheduled_actions`
* **New Data Source:** `awscc_cloudfront_continuous_deployment_policies`
* **New Data Source:** `awscc_cloudfront_continuous_deployment_policy`
* **New Data Source:** `awscc_cloudfront_monitoring_subscription`
* **New Data Source:** `awscc_codedeploy_deployment_config`
* **New Data Source:** `awscc_codedeploy_deployment_configs`
* **New Data Source:** `awscc_datapipeline_pipeline`
* **New Data Source:** `awscc_datapipeline_pipelines`
* **New Data Source:** `awscc_dynamodb_table`
* **New Data Source:** `awscc_dynamodb_tables`
* **New Data Source:** `awscc_ec2_eip`
* **New Data Source:** `awscc_ec2_eips`
* **New Data Source:** `awscc_ec2_volume`
* **New Data Source:** `awscc_ec2_volumes`
* **New Data Source:** `awscc_ec2_vpn_connection`
* **New Data Source:** `awscc_ec2_vpn_connections`
* **New Data Source:** `awscc_elasticache_subnet_group`
* **New Data Source:** `awscc_elasticache_subnet_groups`
* **New Data Source:** `awscc_elasticbeanstalk_application`
* **New Data Source:** `awscc_elasticbeanstalk_application_version`
* **New Data Source:** `awscc_elasticbeanstalk_application_versions`
* **New Data Source:** `awscc_elasticbeanstalk_applications`
* **New Data Source:** `awscc_elasticbeanstalk_environment`
* **New Data Source:** `awscc_elasticbeanstalk_environments`
* **New Data Source:** `awscc_emr_security_configuration`
* **New Data Source:** `awscc_emr_security_configurations`
* **New Data Source:** `awscc_iot_policies`
* **New Data Source:** `awscc_iot_policy`
* **New Data Source:** `awscc_logs_destination`
* **New Data Source:** `awscc_logs_destinations`
* **New Data Source:** `awscc_logs_log_stream`
* **New Data Source:** `awscc_logs_metric_filter`
* **New Data Source:** `awscc_logs_metric_filters`
* **New Data Source:** `awscc_rds_db_cluster`
* **New Data Source:** `awscc_rds_db_cluster_parameter_group`
* **New Data Source:** `awscc_rds_db_cluster_parameter_groups`
* **New Data Source:** `awscc_rds_db_clusters`
* **New Data Source:** `awscc_rds_db_instance`
* **New Data Source:** `awscc_rds_db_instances`
* **New Data Source:** `awscc_rds_option_group`
* **New Data Source:** `awscc_rds_option_groups`
* **New Data Source:** `awscc_redshift_cluster_parameter_group`
* **New Data Source:** `awscc_redshift_cluster_parameter_groups`
* **New Data Source:** `awscc_redshift_cluster_subnet_group`
* **New Data Source:** `awscc_redshift_cluster_subnet_groups`
* **New Resource:** `awscc_apigateway_documentation_part`
* **New Resource:** `awscc_apigatewayv2_api`
* **New Resource:** `awscc_apigatewayv2_authorizer`
* **New Resource:** `awscc_apigatewayv2_deployment`
* **New Resource:** `awscc_apigatewayv2_model`
* **New Resource:** `awscc_appstream_image_builder`
* **New Resource:** `awscc_autoscaling_scaling_policy`
* **New Resource:** `awscc_autoscaling_scheduled_action`
* **New Resource:** `awscc_cloudfront_continuous_deployment_policy`
* **New Resource:** `awscc_cloudfront_monitoring_subscription`
* **New Resource:** `awscc_codedeploy_deployment_config`
* **New Resource:** `awscc_datapipeline_pipeline`
* **New Resource:** `awscc_dynamodb_table`
* **New Resource:** `awscc_ec2_eip`
* **New Resource:** `awscc_ec2_volume`
* **New Resource:** `awscc_ec2_vpn_connection`
* **New Resource:** `awscc_elasticache_subnet_group`
* **New Resource:** `awscc_elasticbeanstalk_application`
* **New Resource:** `awscc_elasticbeanstalk_application_version`
* **New Resource:** `awscc_emr_security_configuration`
* **New Resource:** `awscc_iot_policy`
* **New Resource:** `awscc_logs_destination`
* **New Resource:** `awscc_logs_log_stream`
* **New Resource:** `awscc_logs_metric_filter`
* **New Resource:** `awscc_rds_db_cluster`
* **New Resource:** `awscc_rds_db_cluster_parameter_group`
* **New Resource:** `awscc_rds_db_instance`
* **New Resource:** `awscc_rds_option_group`
* **New Resource:** `awscc_redshift_cluster_parameter_group`
* **New Resource:** `awscc_redshift_cluster_subnet_group`

## [0.39.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.39.0) (November 17, 2022)

FEATURES:

* **New Data Source:** `awscc_organizations_organizational_unit`
* **New Data Source:** `awscc_ssm_resource_policies`
* **New Data Source:** `awscc_ssm_resource_policy`
* **New Resource:** `awscc_organizations_organizational_unit`
* **New Resource:** `awscc_ssm_resource_policy`

## [0.38.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.38.0) (November 10, 2022)

FEATURES:

* **New Data Source:** `awscc_organizations_account`
* **New Data Source:** `awscc_organizations_policy`
* **New Data Source:** `awscc_resourceexplorer2_default_view_association`
* **New Data Source:** `awscc_resourceexplorer2_index`
* **New Data Source:** `awscc_resourceexplorer2_indices`
* **New Data Source:** `awscc_resourceexplorer2_view`
* **New Data Source:** `awscc_resourceexplorer2_views`
* **New Data Source:** `awscc_scheduler_schedule_group`
* **New Data Source:** `awscc_scheduler_schedule_groups`
* **New Data Source:** `awscc_ses_vdm_attributes`
* **New Data Source:** `awscc_xray_resource_policies`
* **New Data Source:** `awscc_xray_resource_policy`
* **New Resource:** `awscc_organizations_account`
* **New Resource:** `awscc_organizations_policy`
* **New Resource:** `awscc_resourceexplorer2_default_view_association`
* **New Resource:** `awscc_resourceexplorer2_index`
* **New Resource:** `awscc_resourceexplorer2_view`
* **New Resource:** `awscc_scheduler_schedule_group`
* **New Resource:** `awscc_ses_vdm_attributes`
* **New Resource:** `awscc_xray_resource_policy`

## [0.37.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.37.0) (November  3, 2022)

FEATURES:

* **New Data Source:** `awscc_apprunner_vpc_ingress_connection`
* **New Data Source:** `awscc_apprunner_vpc_ingress_connections`
* **New Data Source:** `awscc_supportapp_slack_workspace_configuration`
* **New Data Source:** `awscc_supportapp_slack_workspace_configurations`
* **New Resource:** `awscc_apprunner_vpc_ingress_connection`
* **New Resource:** `awscc_supportapp_slack_workspace_configuration`

## [0.36.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.36.0) (October 27, 2022)

BUG FIXES:

* resource/awscc_networkmanager_core_network: Fix `Invalid JSON string` error on resource Create ([#708](https://github.com/hashicorp/terraform-provider-awscc/issues/708))

FEATURES:

* **New Data Source:** `awscc_fsx_data_repository_association`
* **New Data Source:** `awscc_fsx_data_repository_associations`
* **New Resource:** `awscc_fsx_data_repository_association`

## [0.35.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.35.0) (October 20, 2022)

FEATURES:

* **New Data Source:** `awscc_refactorspaces_application`
* **New Resource:** `awscc_refactorspaces_application`

## [0.34.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.34.0) (October 13, 2022)

FEATURES:

* **New Data Source:** `awscc_greengrassv2_deployment`
* **New Data Source:** `awscc_greengrassv2_deployments`
* **New Data Source:** `awscc_transfer_agreement`
* **New Data Source:** `awscc_transfer_certificate`
* **New Data Source:** `awscc_transfer_certificates`
* **New Data Source:** `awscc_transfer_connector`
* **New Data Source:** `awscc_transfer_connectors`
* **New Data Source:** `awscc_transfer_profile`
* **New Data Source:** `awscc_transfer_profiles`
* **New Resource:** `awscc_greengrassv2_deployment`
* **New Resource:** `awscc_transfer_agreement`
* **New Resource:** `awscc_transfer_certificate`
* **New Resource:** `awscc_transfer_connector`
* **New Resource:** `awscc_transfer_profile`

## [0.33.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.33.0) (September 29, 2022)

NOTES:

* Optional attributes without a default value specified in the resource's CloudFormation schema are now handled as _Computed_, indicating that if no value has been configured then Terraform will not perform drift detection on the attribute ([#667](https://github.com/hashicorp/terraform-provider-awscc/issues/667))

FEATURES:

* **New Data Source:** `awscc_identitystore_group`
* **New Data Source:** `awscc_identitystore_group_membership`
* **New Data Source:** `awscc_iotfleetwise_campaign`
* **New Data Source:** `awscc_iotfleetwise_campaigns`
* **New Data Source:** `awscc_iotfleetwise_fleet`
* **New Data Source:** `awscc_iotfleetwise_fleets`
* **New Data Source:** `awscc_iotfleetwise_model_manifest`
* **New Data Source:** `awscc_iotfleetwise_model_manifests`
* **New Data Source:** `awscc_iotfleetwise_signal_catalog`
* **New Data Source:** `awscc_iotfleetwise_signal_catalogs`
* **New Data Source:** `awscc_iotfleetwise_vehicle`
* **New Data Source:** `awscc_iotfleetwise_vehicles`
* **New Data Source:** `awscc_m2_application`
* **New Data Source:** `awscc_m2_applications`
* **New Resource:** `awscc_identitystore_group`
* **New Resource:** `awscc_identitystore_group_membership`
* **New Resource:** `awscc_iotfleetwise_campaign`
* **New Resource:** `awscc_iotfleetwise_fleet`
* **New Resource:** `awscc_iotfleetwise_model_manifest`
* **New Resource:** `awscc_iotfleetwise_signal_catalog`
* **New Resource:** `awscc_iotfleetwise_vehicle`
* **New Resource:** `awscc_m2_application`

## [0.32.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.32.0) (September  8, 2022)

FEATURES:

* **New Data Source:** `awscc_cloudfront_origin_access_control`
* **New Data Source:** `awscc_cloudfront_origin_access_controls`
* **New Data Source:** `awscc_m2_environment`
* **New Data Source:** `awscc_m2_environments`
* **New Resource:** `awscc_cloudfront_origin_access_control`
* **New Resource:** `awscc_m2_environment`

## [0.31.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.31.0) (September  1, 2022)

FEATURES:

* **New Data Source:** `awscc_connect_instance_storage_config`
* **New Data Source:** `awscc_controltower_enabled_control`
* **New Data Source:** `awscc_macie_allow_list`
* **New Data Source:** `awscc_supportapp_account_alias`
* **New Data Source:** `awscc_supportapp_account_aliases`
* **New Data Source:** `awscc_supportapp_slack_channel_configuration`
* **New Data Source:** `awscc_supportapp_slack_channel_configurations`
* **New Resource:** `awscc_connect_instance_storage_config`
* **New Resource:** `awscc_controltower_enabled_control`
* **New Resource:** `awscc_macie_allow_list`
* **New Resource:** `awscc_supportapp_account_alias`
* **New Resource:** `awscc_supportapp_slack_channel_configuration`
* provider: Support `me-central-1` as a valid AWS Region

## [0.30.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.30.0) (August 18, 2022)

FEATURES:

* **New Data Source:** `awscc_msk_serverless_cluster`
* **New Data Source:** `awscc_msk_serverless_clusters`
* **New Resource:** `awscc_msk_serverless_cluster`

## [0.29.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.29.0) (August  4, 2022)

BUG FIXES:

* Documentation: Nested attributes are now correctly grouped in "optional", "required" and "read-only" ([#580](https://github.com/hashicorp/terraform-provider-awscc/issues/580))

## [0.28.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.28.0) (July 21, 2022)

FEATURES:

* **New Data Source:** `awscc_evidently_segment`
* **New Data Source:** `awscc_redshiftserverless_workgroup`
* **New Data Source:** `awscc_redshiftserverless_workgroups`
* **New Data Source:** `awscc_rolesanywhere_crl`
* **New Data Source:** `awscc_rolesanywhere_crls`
* **New Resource:** `awscc_evidently_segment`
* **New Resource:** `awscc_redshiftserverless_workgroup`
* **New Resource:** `awscc_rolesanywhere_crl`

## [0.27.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.27.0) (July  7, 2022)

FEATURES:

* **New Data Source:** `awscc_datasync_location_fsx_ontap`
* **New Data Source:** `awscc_datasync_location_fsx_ontaps`
* **New Data Source:** `awscc_rolesanywhere_profile`
* **New Data Source:** `awscc_rolesanywhere_profiles`
* **New Data Source:** `awscc_rolesanywhere_trust_anchor`
* **New Resource:** `awscc_datasync_location_fsx_ontap`
* **New Resource:** `awscc_rolesanywhere_profile`
* **New Resource:** `awscc_rolesanywhere_trust_anchor`

## [0.26.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.26.0) (June 30, 2022)

BREAKING CHANGES:

* data-source/awscc_rds_db_cluster_parameter_group: Temporarily removed
* data-source/awscc_rds_db_cluster_parameter_groups: Temporarily removed
* resource/awscc_rds_db_cluster_parameter_group: Temporarily removed

FEATURES:

* **New Data Source:** `awscc_appstream_directory_config`
* **New Data Source:** `awscc_appstream_directory_configs`
* **New Data Source:** `awscc_cloudtrail_event_data_store`
* **New Data Source:** `awscc_cloudtrail_event_data_stores`
* **New Data Source:** `awscc_ec2_capacity_reservation`
* **New Data Source:** `awscc_ec2_capacity_reservations`
* **New Data Source:** `awscc_lakeformation_data_cells_filter`
* **New Data Source:** `awscc_lakeformation_data_cells_filters`
* **New Data Source:** `awscc_lakeformation_principal_permissions`
* **New Data Source:** `awscc_lakeformation_tag`
* **New Data Source:** `awscc_lakeformation_tag_association`
* **New Data Source:** `awscc_lakeformation_tags`
* **New Data Source:** `awscc_redshiftserverless_namespaces`
* **New Data Source:** `awscc_ses_dedicated_ip_pool`
* **New Data Source:** `awscc_ses_dedicated_ip_pools`
* **New Data Source:** `awscc_ses_email_identities`
* **New Data Source:** `awscc_ses_email_identity`
* **New Resource:** `awscc_appstream_directory_config`
* **New Resource:** `awscc_cloudtrail_event_data_store`
* **New Resource:** `awscc_ec2_capacity_reservation`
* **New Resource:** `awscc_lakeformation_data_cells_filter`
* **New Resource:** `awscc_lakeformation_principal_permissions`
* **New Resource:** `awscc_lakeformation_tag`
* **New Resource:** `awscc_lakeformation_tag_association`
* **New Resource:** `awscc_ses_dedicated_ip_pool`
* **New Resource:** `awscc_ses_email_identity`

## [0.25.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.25.0) (June 23, 2022)

FEATURES:

* **New Data Source:** `awscc_apigatewayv2_vpc_link`
* **New Data Source:** `awscc_apigatewayv2_vpc_links`
* **New Data Source:** `awscc_connect_task_template`
* **New Data Source:** `awscc_connectcampaigns_campaign`
* **New Data Source:** `awscc_connectcampaigns_campaigns`
* **New Data Source:** `awscc_ec2_customer_gateway`
* **New Data Source:** `awscc_ec2_customer_gateways`
* **New Data Source:** `awscc_ec2_nat_gateway`
* **New Data Source:** `awscc_ec2_nat_gateways`
* **New Data Source:** `awscc_ec2_vpn_gateway`
* **New Data Source:** `awscc_ec2_vpn_gateways`
* **New Data Source:** `awscc_iot_ca_certificate`
* **New Data Source:** `awscc_iot_ca_certificates`
* **New Data Source:** `awscc_rds_db_cluster_parameter_group`
* **New Data Source:** `awscc_rds_db_cluster_parameter_groups`
* **New Data Source:** `awscc_rds_db_parameter_group`
* **New Data Source:** `awscc_rds_db_parameter_groups`
* **New Data Source:** `awscc_rds_db_subnet_group`
* **New Data Source:** `awscc_rds_db_subnet_groups`
* **New Data Source:** `awscc_rds_event_subscription`
* **New Data Source:** `awscc_rds_event_subscriptions`
* **New Data Source:** `awscc_redshiftserverless_namespace`
* **New Data Source:** `awscc_route53_cidr_collection`
* **New Data Source:** `awscc_route53_cidr_collections`
* **New Resource:** `awscc_apigatewayv2_vpc_link`
* **New Resource:** `awscc_connect_task_template`
* **New Resource:** `awscc_connectcampaigns_campaign`
* **New Resource:** `awscc_ec2_customer_gateway`
* **New Resource:** `awscc_ec2_nat_gateway`
* **New Resource:** `awscc_ec2_vpn_gateway`
* **New Resource:** `awscc_iot_ca_certificate`
* **New Resource:** `awscc_rds_db_cluster_parameter_group`
* **New Resource:** `awscc_rds_db_parameter_group`
* **New Resource:** `awscc_rds_db_subnet_group`
* **New Resource:** `awscc_rds_event_subscription`
* **New Resource:** `awscc_redshiftserverless_namespace`
* **New Resource:** `awscc_route53_cidr_collection`

BUG FIXES:

* resource/awscc_networkmanager_core_network: Fix `Model validation failed` error ([#537](https://github.com/hashicorp/terraform-provider-awscc/issues/537))

## [0.24.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.24.0) (June  2, 2022)

FEATURES:

* **New Data Source:** `awscc_emrserverless_applications`
* **New Data Source:** `awscc_kinesisanalyticsv2_application`
* **New Data Source:** `awscc_kinesisanalyticsv2_applications`
* **New Resource:** `awscc_kinesisanalyticsv2_application`

## [0.23.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.23.0) (May 26, 2022)

FEATURES:

* **New Data Source:** `awscc_ec2_placement_group`
* **New Data Source:** `awscc_ec2_placement_groups`
* **New Data Source:** `awscc_ec2_vpc_peering_connection`
* **New Data Source:** `awscc_ec2_vpc_peering_connections`
* **New Data Source:** `awscc_emrserverless_application`
* **New Data Source:** `awscc_iotwireless_network_analyzer_configuration`
* **New Data Source:** `awscc_iotwireless_network_analyzer_configurations`
* **New Resource:** `awscc_ec2_placement_group`
* **New Resource:** `awscc_ec2_vpc_peering_connection`
* **New Resource:** `awscc_emrserverless_application`
* **New Resource:** `awscc_iotwireless_network_analyzer_configuration`

## [0.22.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.22.0) (May 20, 2022)

FEATURES:

* **New Data Source:** `awscc_sagemaker_model_package`
* **New Data Source:** `awscc_sagemaker_model_packages`
* **New Resource:** `awscc_sagemaker_model_package`

BUG FIXES:

* Prevent infinite loop if waiting for any asynchronous operation to complete returns an AWS API error ([#510](https://github.com/hashicorp/terraform-provider-awscc/issues/510))

## [0.21.0](https://github.com/hashicorp/terraform-provider-awscc/releases/tag/v0.21.0) (May 12, 2022)

FEATURES:

* **New Data Source:** `awscc_iot_role_alias`
* **New Data Source:** `awscc_iot_role_aliases`
* **New Data Source:** `awscc_macie_sessions`
* **New Data Source:** `awscc_networkmanager_connect_attachment`
* **New Data Source:** `awscc_networkmanager_connect_attachments`
* **New Data Source:** `awscc_networkmanager_connect_peer`
* **New Data Source:** `awscc_networkmanager_connect_peers`
* **New Data Source:** `awscc_networkmanager_core_network`
* **New Data Source:** `awscc_networkmanager_core_networks`
* **New Data Source:** `awscc_networkmanager_site_to_site_vpn_attachment`
* **New Data Source:** `awscc_networkmanager_site_to_site_vpn_attachments`
* **New Data Source:** `awscc_networkmanager_vpc_attachment`
* **New Data Source:** `awscc_networkmanager_vpc_attachments`
* **New Resource:** `awscc_iot_role_alias`
* **New Resource:** `awscc_networkmanager_connect_attachment`
* **New Resource:** `awscc_networkmanager_connect_peer`
* **New Resource:** `awscc_networkmanager_core_network`
* **New Resource:** `awscc_networkmanager_site_to_site_vpn_attachment`
* **New Resource:** `awscc_networkmanager_vpc_attachment`

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
