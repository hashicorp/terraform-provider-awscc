## 0.7.0 (Unreleased)

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
