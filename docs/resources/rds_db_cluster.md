---
page_title: "awscc_rds_db_cluster Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::RDS::DBCluster resource creates an Amazon Aurora DB cluster or Multi-AZ DB cluster.
  For more information about creating an Aurora DB cluster, see Creating an Amazon Aurora DB cluster https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.CreateInstance.html in the Amazon Aurora User Guide.
  For more information about creating a Multi-AZ DB cluster, see Creating a Multi-AZ DB cluster https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/create-multi-az-db-cluster.html in the Amazon RDS User Guide.
  You can only create this resource in AWS Regions where Amazon Aurora or Multi-AZ DB clusters are supported.
  Updating DB clusters
  When properties labeled "Update requires:Replacement https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement" are updated, AWS CloudFormation first creates a replacement DB cluster, then changes references from other dependent resources to point to the replacement DB cluster, and finally deletes the old DB cluster.
  We highly recommend that you take a snapshot of the database before updating the stack. If you don't, you lose the data when AWS CloudFormation replaces your DB cluster. To preserve your data, perform the following procedure:
  Deactivate any applications that are using the DB cluster so that there's no activity on the DB instance.Create a snapshot of the DB cluster. For more information, see Creating a DB cluster snapshot https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CreateSnapshotCluster.html.If you want to restore your DB cluster using a DB cluster snapshot, modify the updated template with your DB cluster changes and add the SnapshotIdentifier property with the ID of the DB cluster snapshot that you want to use.
  After you restore a DB cluster with a SnapshotIdentifier property, you must specify the same SnapshotIdentifier property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the DB cluster snapshot again, and the data in the database is not changed. However, if you don't specify the SnapshotIdentifier property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB cluster is restored from the specified SnapshotIdentifier property, and the original DB cluster is deleted.Update the stack.
  Currently, when you are updating the stack for an Aurora Serverless DB cluster, you can't include changes to any other properties when you specify one of the following properties: PreferredBackupWindow, PreferredMaintenanceWindow, and Port. This limitation doesn't apply to provisioned DB clusters.
  For more information about updating other properties of this resource, see ModifyDBCluster. For more information about updating stacks, see CloudFormation Stacks Updates https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html.
  Deleting DB clusters
  The default DeletionPolicy for AWS::RDS::DBCluster resources is Snapshot. For more information about how AWS CloudFormation deletes resources, see DeletionPolicy Attribute https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html.
---

# awscc_rds_db_cluster (Resource)

The ``AWS::RDS::DBCluster`` resource creates an Amazon Aurora DB cluster or Multi-AZ DB cluster.
 For more information about creating an Aurora DB cluster, see [Creating an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.CreateInstance.html) in the *Amazon Aurora User Guide*.
 For more information about creating a Multi-AZ DB cluster, see [Creating a Multi-AZ DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/create-multi-az-db-cluster.html) in the *Amazon RDS User Guide*.
  You can only create this resource in AWS Regions where Amazon Aurora or Multi-AZ DB clusters are supported.
   *Updating DB clusters* 
 When properties labeled "*Update requires:*[Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)" are updated, AWS CloudFormation first creates a replacement DB cluster, then changes references from other dependent resources to point to the replacement DB cluster, and finally deletes the old DB cluster.
  We highly recommend that you take a snapshot of the database before updating the stack. If you don't, you lose the data when AWS CloudFormation replaces your DB cluster. To preserve your data, perform the following procedure:
  1.  Deactivate any applications that are using the DB cluster so that there's no activity on the DB instance.
  1.  Create a snapshot of the DB cluster. For more information, see [Creating a DB cluster snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_CreateSnapshotCluster.html).
  1.  If you want to restore your DB cluster using a DB cluster snapshot, modify the updated template with your DB cluster changes and add the ``SnapshotIdentifier`` property with the ID of the DB cluster snapshot that you want to use.
 After you restore a DB cluster with a ``SnapshotIdentifier`` property, you must specify the same ``SnapshotIdentifier`` property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the DB cluster snapshot again, and the data in the database is not changed. However, if you don't specify the ``SnapshotIdentifier`` property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB cluster is restored from the specified ``SnapshotIdentifier`` property, and the original DB cluster is deleted.
  1.  Update the stack.
  
  Currently, when you are updating the stack for an Aurora Serverless DB cluster, you can't include changes to any other properties when you specify one of the following properties: ``PreferredBackupWindow``, ``PreferredMaintenanceWindow``, and ``Port``. This limitation doesn't apply to provisioned DB clusters.
 For more information about updating other properties of this resource, see ``ModifyDBCluster``. For more information about updating stacks, see [CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html).
  *Deleting DB clusters* 
 The default ``DeletionPolicy`` for ``AWS::RDS::DBCluster`` resources is ``Snapshot``. For more information about how AWS CloudFormation deletes resources, see [DeletionPolicy Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html).

## Example Usage

### Managed Master Passwords via Secrets Manager, specify Availability Zones example
You can specify the manage_master_user_password attribute to enable managing the master password with Secrets Manager. You can also update an existing cluster to use Secrets Manager by specify the manage_master_user_password attribute and removing the password attribute (removal is required).
```terraform
resource "awscc_rds_db_cluster" "example_db_cluster" {
  availability_zones          = ["us-east-1b", "us-east-1c"]
  engine                      = "aurora-mysql"
  db_cluster_identifier       = "example-dbcluster"
  manage_master_user_password = true
  master_username             = "foo"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `allocated_storage` (Number) The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.
 Valid for Cluster Type: Multi-AZ DB clusters only
 This setting is required to create a Multi-AZ DB cluster.
- `associated_roles` (Attributes List) Provides a list of the AWS Identity and Access Management (IAM) roles that are associated with the DB cluster. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other Amazon Web Services on your behalf.
 Valid for: Aurora DB clusters and Multi-AZ DB clusters (see [below for nested schema](#nestedatt--associated_roles))
- `auto_minor_version_upgrade` (Boolean) Specifies whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window. By default, minor engine upgrades are applied automatically.
 Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB cluster.
 For more information about automatic minor version upgrades, see [Automatically upgrading the minor engine version](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Upgrading.html#USER_UpgradeDBInstance.Upgrading.AutoMinorVersionUpgrades).
- `availability_zones` (List of String) A list of Availability Zones (AZs) where instances in the DB cluster can be created. For information on AWS Regions and Availability Zones, see [Choosing the Regions and Availability Zones](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html) in the *Amazon Aurora User Guide*. 
 Valid for: Aurora DB clusters only
- `backtrack_window` (Number) The target backtrack window, in seconds. To disable backtracking, set this value to ``0``.
 Valid for Cluster Type: Aurora MySQL DB clusters only
 Default: ``0``
 Constraints:
  +  If specified, this value must be set to a number from 0 to 259,200 (72 hours).
- `backup_retention_period` (Number) The number of days for which automated backups are retained.
 Default: 1
 Constraints:
  +  Must be a value from 1 to 35
  
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `cluster_scalability_type` (String) Specifies the scalability mode of the Aurora DB cluster. When set to ``limitless``, the cluster operates as an Aurora Limitless Database, allowing you to create a DB shard group for horizontal scaling (sharding) capabilities. When set to ``standard`` (the default), the cluster uses normal DB instance creation.
- `copy_tags_to_snapshot` (Boolean) A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default is not to copy them.
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `database_insights_mode` (String) The mode of Database Insights to enable for the DB cluster.
 If you set this value to ``advanced``, you must also set the ``PerformanceInsightsEnabled`` parameter to ``true`` and the ``PerformanceInsightsRetentionPeriod`` parameter to 465.
 Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
- `database_name` (String) The name of your database. If you don't provide a name, then Amazon RDS won't create a database in this DB cluster. For naming constraints, see [Naming Constraints](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_Limits.html#RDS_Limits.Constraints) in the *Amazon Aurora User Guide*. 
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `db_cluster_identifier` (String) The DB cluster identifier. This parameter is stored as a lowercase string.
 Constraints:
  +  Must contain from 1 to 63 letters, numbers, or hyphens.
  +  First character must be a letter.
  +  Can't end with a hyphen or contain two consecutive hyphens.
  
 Example: ``my-cluster1``
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `db_cluster_instance_class` (String) The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example ``db.m6gd.xlarge``. Not all DB instance classes are available in all AWS-Regions, or for all database engines.
 For the full list of DB instance classes and availability for your engine, see [DB instance class](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the *Amazon RDS User Guide*.
 This setting is required to create a Multi-AZ DB cluster.
 Valid for Cluster Type: Multi-AZ DB clusters only
- `db_cluster_parameter_group_name` (String) The name of the DB cluster parameter group to associate with this DB cluster.
  If you apply a parameter group to an existing DB cluster, then its DB instances might need to reboot. This can result in an outage while the DB instances are rebooting.
 If you apply a change to parameter group associated with a stopped DB cluster, then the update stack waits until the DB cluster is started.
  To list all of the available DB cluster parameter group names, use the following command:
  ``aws rds describe-db-cluster-parameter-groups --query "DBClusterParameterGroups[].DBClusterParameterGroupName" --output text`` 
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `db_instance_parameter_group_name` (String) The name of the DB parameter group to apply to all instances of the DB cluster.
  When you apply a parameter group using the ``DBInstanceParameterGroupName`` parameter, the DB cluster isn't rebooted automatically. Also, parameter changes are applied immediately rather than during the next maintenance window.
  Valid for Cluster Type: Aurora DB clusters only
 Default: The existing name setting
 Constraints:
  +  The DB parameter group must be in the same DB parameter group family as this DB cluster.
  +  The ``DBInstanceParameterGroupName`` parameter is valid in combination with the ``AllowMajorVersionUpgrade`` parameter for a major version upgrade only.
- `db_subnet_group_name` (String) A DB subnet group that you want to associate with this DB cluster. 
 If you are restoring a DB cluster to a point in time with ``RestoreType`` set to ``copy-on-write``, and don't specify a DB subnet group name, then the DB cluster is restored with a default DB subnet group.
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `db_system_id` (String) Reserved for future use.
- `deletion_protection` (Boolean) A value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection is disabled.
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `domain` (String) Indicates the directory ID of the Active Directory to create the DB cluster.
 For Amazon Aurora DB clusters, Amazon RDS can use Kerberos authentication to authenticate users that connect to the DB cluster.
 For more information, see [Kerberos authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html) in the *Amazon Aurora User Guide*.
 Valid for: Aurora DB clusters only
- `domain_iam_role_name` (String) Specifies the name of the IAM role to use when making API calls to the Directory Service.
 Valid for: Aurora DB clusters only
- `enable_cloudwatch_logs_exports` (List of String) The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine being used. For more information, see [Publishing Database Logs to Amazon CloudWatch Logs](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the *Amazon Aurora User Guide*.
  *Aurora MySQL* 
 Valid values: ``audit``, ``error``, ``general``, ``slowquery``
  *Aurora PostgreSQL* 
 Valid values: ``postgresql``
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `enable_global_write_forwarding` (Boolean) Specifies whether to enable this DB cluster to forward write operations to the primary cluster of a global cluster (Aurora global database). By default, write operations are not allowed on Aurora DB clusters that are secondary clusters in an Aurora global database.
 You can set this value only on Aurora DB clusters that are members of an Aurora global database. With this parameter enabled, a secondary cluster can forward writes to the current primary cluster, and the resulting changes are replicated back to this cluster. For the primary DB cluster of an Aurora global database, this value is used immediately if the primary is demoted by a global cluster API operation, but it does nothing until then.
 Valid for Cluster Type: Aurora DB clusters only
- `enable_http_endpoint` (Boolean) Specifies whether to enable the HTTP endpoint for the DB cluster. By default, the HTTP endpoint isn't enabled.
 When enabled, the HTTP endpoint provides a connectionless web service API (RDS Data API) for running SQL queries on the DB cluster. You can also query your database from inside the RDS console with the RDS query editor.
 For more information, see [Using RDS Data API](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in the *Amazon Aurora User Guide*.
 Valid for Cluster Type: Aurora DB clusters only
- `enable_iam_database_authentication` (Boolean) A value that indicates whether to enable mapping of AWS Identity and Access Management (IAM) accounts to database accounts. By default, mapping is disabled.
 For more information, see [IAM Database Authentication](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) in the *Amazon Aurora User Guide.*
 Valid for: Aurora DB clusters only
- `enable_local_write_forwarding` (Boolean) Specifies whether read replicas can forward write operations to the writer DB instance in the DB cluster. By default, write operations aren't allowed on reader DB instances.
 Valid for: Aurora DB clusters only
- `engine` (String) The name of the database engine to be used for this DB cluster.
 Valid Values:
  +   ``aurora-mysql`` 
  +   ``aurora-postgresql`` 
  +   ``mysql`` 
  +   ``postgres`` 
  
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `engine_lifecycle_support` (String) The life cycle type for this DB cluster.
  By default, this value is set to ``open-source-rds-extended-support``, which enrolls your DB cluster into Amazon RDS Extended Support. At the end of standard support, you can avoid charges for Extended Support by setting the value to ``open-source-rds-extended-support-disabled``. In this case, creating the DB cluster will fail if the DB major version is past its end of standard support date.
  You can use this setting to enroll your DB cluster into Amazon RDS Extended Support. With RDS Extended Support, you can run the selected major engine version on your DB cluster past the end of standard support for that engine version. For more information, see the following sections:
  +  Amazon Aurora - [Amazon RDS Extended Support with Amazon Aurora](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/extended-support.html) in the *Amazon Aurora User Guide*
  +  Amazon RDS - [Amazon RDS Extended Support with Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/extended-support.html) in the *Amazon RDS User Guide*
  
 Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
 Valid Values: ``open-source-rds-extended-support | open-source-rds-extended-support-disabled``
 Default: ``open-source-rds-extended-support``
- `engine_mode` (String) The DB engine mode of the DB cluster, either ``provisioned`` or ``serverless``.
 The ``serverless`` engine mode only applies for Aurora Serverless v1 DB clusters. Aurora Serverless v2 DB clusters use the ``provisioned`` engine mode.
 For information about limitations and requirements for Serverless DB clusters, see the following sections in the *Amazon Aurora User Guide*:
  +   [Limitations of Aurora Serverless v1](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html#aurora-serverless.limitations) 
  +   [Requirements for Aurora Serverless v2](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.requirements.html) 
  
 Valid for Cluster Type: Aurora DB clusters only
- `engine_version` (String) The version number of the database engine to use.
 To list all of the available engine versions for Aurora MySQL version 2 (5.7-compatible) and version 3 (8.0-compatible), use the following command:
  ``aws rds describe-db-engine-versions --engine aurora-mysql --query "DBEngineVersions[].EngineVersion"`` 
 You can supply either ``5.7`` or ``8.0`` to use the default engine version for Aurora MySQL version 2 or version 3, respectively.
 To list all of the available engine versions for Aurora PostgreSQL, use the following command:
  ``aws rds describe-db-engine-versions --engine aurora-postgresql --query "DBEngineVersions[].EngineVersion"`` 
 To list all of the available engine versions for RDS for MySQL, use the following command:
  ``aws rds describe-db-engine-versions --engine mysql --query "DBEngineVersions[].EngineVersion"`` 
 To list all of the available engine versions for RDS for PostgreSQL, use the following command:
  ``aws rds describe-db-engine-versions --engine postgres --query "DBEngineVersions[].EngineVersion"`` 
  *Aurora MySQL* 
 For information, see [Database engine updates for Amazon Aurora MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html) in the *Amazon Aurora User Guide*.
  *Aurora PostgreSQL* 
 For information, see [Amazon Aurora PostgreSQL releases and engine versions](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.20180305.html) in the *Amazon Aurora User Guide*.
  *MySQL* 
 For information, see [Amazon RDS for MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the *Amazon RDS User Guide*.
  *PostgreSQL* 
 For information, see [Amazon RDS for PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts) in the *Amazon RDS User Guide*.
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `global_cluster_identifier` (String) If you are configuring an Aurora global database cluster and want your Aurora DB cluster to be a secondary member in the global database cluster, specify the global cluster ID of the global database cluster. To define the primary database cluster of the global cluster, use the [AWS::RDS::GlobalCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-globalcluster.html) resource. 
  If you aren't configuring a global database cluster, don't specify this property. 
  To remove the DB cluster from a global database cluster, specify an empty value for the ``GlobalClusterIdentifier`` property.
  For information about Aurora global databases, see [Working with Amazon Aurora Global Databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html) in the *Amazon Aurora User Guide*.
 Valid for: Aurora DB clusters only
- `iops` (Number) The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.
 For information about valid IOPS values, see [Provisioned IOPS storage](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS) in the *Amazon RDS User Guide*.
 This setting is required to create a Multi-AZ DB cluster.
 Valid for Cluster Type: Multi-AZ DB clusters only
 Constraints:
  +  Must be a multiple between .5 and 50 of the storage amount for the DB cluster.
- `kms_key_id` (String) The Amazon Resource Name (ARN) of the AWS KMS key that is used to encrypt the database instances in the DB cluster, such as ``arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef``. If you enable the ``StorageEncrypted`` property but don't specify this property, the default KMS key is used. If you specify this property, you must set the ``StorageEncrypted`` property to ``true``.
 If you specify the ``SnapshotIdentifier`` property, the ``StorageEncrypted`` property value is inherited from the snapshot, and if the DB cluster is encrypted, the specified ``KmsKeyId`` property is used.
 If you create a read replica of an encrypted DB cluster in another AWS Region, make sure to set ``KmsKeyId`` to a KMS key identifier that is valid in the destination AWS Region. This KMS key is used to encrypt the read replica in that AWS Region.
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `manage_master_user_password` (Boolean) Specifies whether to manage the master user password with AWS Secrets Manager.
 For more information, see [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide* and [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html) in the *Amazon Aurora User Guide.*
 Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
 Constraints:
  +  Can't manage the master user password with AWS Secrets Manager if ``MasterUserPassword`` is specified.
- `master_user_password` (String) The master password for the DB instance.
  If you specify the ``SourceDBClusterIdentifier``, ``SnapshotIdentifier``, or ``GlobalClusterIdentifier`` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.
  Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `master_user_secret` (Attributes) The secret managed by RDS in AWS Secrets Manager for the master user password.
  When you restore a DB cluster from a snapshot, Amazon RDS generates a new secret instead of reusing the secret specified in the ``SecretArn`` property. This ensures that the restored DB cluster is securely managed with a dedicated secret. To maintain consistent integration with your application, you might need to update resource configurations to reference the newly created secret.
  For more information, see [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the *Amazon RDS User Guide* and [Password management with Secrets Manager](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html) in the *Amazon Aurora User Guide.* (see [below for nested schema](#nestedatt--master_user_secret))
- `master_username` (String) The name of the master user for the DB cluster.
  If you specify the ``SourceDBClusterIdentifier``, ``SnapshotIdentifier``, or ``GlobalClusterIdentifier`` property, don't specify this property. The value is inherited from the source DB cluster, the snapshot, or the primary DB cluster for the global database cluster, respectively.
  Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `monitoring_interval` (Number) The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster. To turn off collecting Enhanced Monitoring metrics, specify ``0``.
 If ``MonitoringRoleArn`` is specified, also set ``MonitoringInterval`` to a value other than ``0``.
 Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
 Valid Values: ``0 | 1 | 5 | 10 | 15 | 30 | 60``
 Default: ``0``
- `monitoring_role_arn` (String) The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs. An example is ``arn:aws:iam:123456789012:role/emaccess``. For information on creating a monitoring role, see [Setting up and enabling Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the *Amazon RDS User Guide*.
 If ``MonitoringInterval`` is set to a value other than ``0``, supply a ``MonitoringRoleArn`` value.
 Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
- `network_type` (String) The network type of the DB cluster.
 Valid values:
  +   ``IPV4`` 
  +   ``DUAL`` 
  
 The network type is determined by the ``DBSubnetGroup`` specified for the DB cluster. A ``DBSubnetGroup`` can support only the IPv4 protocol or the IPv4 and IPv6 protocols (``DUAL``).
 For more information, see [Working with a DB instance in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the *Amazon Aurora User Guide.*
 Valid for: Aurora DB clusters only
- `performance_insights_enabled` (Boolean) Specifies whether to turn on Performance Insights for the DB cluster.
 For more information, see [Using Amazon Performance Insights](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the *Amazon RDS User Guide*.
 Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
- `performance_insights_kms_key_id` (String) The AWS KMS key identifier for encryption of Performance Insights data.
 The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
 If you don't specify a value for ``PerformanceInsightsKMSKeyId``, then Amazon RDS uses your default KMS key. There is a default KMS key for your AWS-account. Your AWS-account has a different default KMS key for each AWS-Region.
 Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
- `performance_insights_retention_period` (Number) The number of days to retain Performance Insights data. When creating a DB cluster without enabling Performance Insights, you can't specify the parameter ``PerformanceInsightsRetentionPeriod``.
 Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
 Valid Values:
  +   ``7`` 
  +  *month* * 31, where *month* is a number of months from 1-23. Examples: ``93`` (3 months * 31), ``341`` (11 months * 31), ``589`` (19 months * 31)
  +   ``731`` 
  
 Default: ``7`` days
 If you specify a retention period that isn't valid, such as ``94``, Amazon RDS issues an error.
- `port` (Number) The port number on which the DB instances in the DB cluster accept connections.
 Default:
  +  When ``EngineMode`` is ``provisioned``, ``3306`` (for both Aurora MySQL and Aurora PostgreSQL)
  +  When ``EngineMode`` is ``serverless``:
  +  ``3306`` when ``Engine`` is ``aurora`` or ``aurora-mysql``
  +  ``5432`` when ``Engine`` is ``aurora-postgresql``
  
  
  The ``No interruption`` on update behavior only applies to DB clusters. If you are updating a DB instance, see [Port](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-database-instance.html#cfn-rds-dbinstance-port) for the AWS::RDS::DBInstance resource.
  Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `preferred_backup_window` (String) The daily time range during which automated backups are created. For more information, see [Backup Window](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow) in the *Amazon Aurora User Guide.*
 Constraints:
  +  Must be in the format ``hh24:mi-hh24:mi``.
  +  Must be in Universal Coordinated Time (UTC).
  +  Must not conflict with the preferred maintenance window.
  +  Must be at least 30 minutes.
  
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `preferred_maintenance_window` (String) The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).
 Format: ``ddd:hh24:mi-ddd:hh24:mi``
 The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region, occurring on a random day of the week. To see the time blocks available, see [Maintaining an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora) in the *Amazon Aurora User Guide.*
 Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
 Constraints: Minimum 30-minute window.
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `publicly_accessible` (Boolean) Specifies whether the DB cluster is publicly accessible.
 When the DB cluster is publicly accessible and you connect from outside of the DB cluster's virtual private cloud (VPC), its Domain Name System (DNS) endpoint resolves to the public IP address. When you connect from within the same VPC as the DB cluster, the endpoint resolves to the private IP address. Access to the DB cluster is ultimately controlled by the security group it uses. That public access isn't permitted if the security group assigned to the DB cluster doesn't permit it.
 When the DB cluster isn't publicly accessible, it is an internal DB cluster with a DNS name that resolves to a private IP address.
 Valid for Cluster Type: Multi-AZ DB clusters only
 Default: The default behavior varies depending on whether ``DBSubnetGroupName`` is specified.
 If ``DBSubnetGroupName`` isn't specified, and ``PubliclyAccessible`` isn't specified, the following applies:
  +  If the default VPC in the target Region doesn?t have an internet gateway attached to it, the DB cluster is private.
  +  If the default VPC in the target Region has an internet gateway attached to it, the DB cluster is public.
  
 If ``DBSubnetGroupName`` is specified, and ``PubliclyAccessible`` isn't specified, the following applies:
  +  If the subnets are part of a VPC that doesn?t have an internet gateway attached to it, the DB cluster is private.
  +  If the subnets are part of a VPC that has an internet gateway attached to it, the DB cluster is public.
- `replication_source_identifier` (String) The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica.
 Valid for: Aurora DB clusters only
- `restore_to_time` (String) The date and time to restore the DB cluster to.
 Valid Values: Value must be a time in Universal Coordinated Time (UTC) format
 Constraints:
  +  Must be before the latest restorable time for the DB instance
  +  Must be specified if ``UseLatestRestorableTime`` parameter isn't provided
  +  Can't be specified if the ``UseLatestRestorableTime`` parameter is enabled
  +  Can't be specified if the ``RestoreType`` parameter is ``copy-on-write``
  
 This property must be used with ``SourceDBClusterIdentifier`` property. The resulting cluster will have the identifier that matches the value of the ``DBclusterIdentifier`` property.
 Example: ``2015-03-07T23:45:00Z``
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `restore_type` (String) The type of restore to be performed. You can specify one of the following values:
  +  ``full-copy`` - The new DB cluster is restored as a full copy of the source DB cluster.
  +  ``copy-on-write`` - The new DB cluster is restored as a clone of the source DB cluster.
  
 If you don't specify a ``RestoreType`` value, then the new DB cluster is restored as a full copy of the source DB cluster.
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `scaling_configuration` (Attributes) The scaling configuration of an Aurora Serverless v1 DB cluster.
 This property is only supported for Aurora Serverless v1. For Aurora Serverless v2, Use the ``ServerlessV2ScalingConfiguration`` property.
 Valid for: Aurora Serverless v1 DB clusters only (see [below for nested schema](#nestedatt--scaling_configuration))
- `serverless_v2_scaling_configuration` (Attributes) The scaling configuration of an Aurora Serverless V2 DB cluster. 
 This property is only supported for Aurora Serverless v2. For Aurora Serverless v1, Use the ``ScalingConfiguration`` property.
 Valid for: Aurora Serverless v2 DB clusters only (see [below for nested schema](#nestedatt--serverless_v2_scaling_configuration))
- `snapshot_identifier` (String) The identifier for the DB snapshot or DB cluster snapshot to restore from.
 You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot. However, you can use only the ARN to specify a DB snapshot.
 After you restore a DB cluster with a ``SnapshotIdentifier`` property, you must specify the same ``SnapshotIdentifier`` property for any future updates to the DB cluster. When you specify this property for an update, the DB cluster is not restored from the snapshot again, and the data in the database is not changed. However, if you don't specify the ``SnapshotIdentifier`` property, an empty DB cluster is created, and the original DB cluster is deleted. If you specify a property that is different from the previous snapshot restore property, a new DB cluster is restored from the specified ``SnapshotIdentifier`` property, and the original DB cluster is deleted.
 If you specify the ``SnapshotIdentifier`` property to restore a DB cluster (as opposed to specifying it for DB cluster updates), then don't specify the following properties:
  +   ``GlobalClusterIdentifier`` 
  +   ``MasterUsername`` 
  +   ``MasterUserPassword`` 
  +   ``ReplicationSourceIdentifier`` 
  +   ``RestoreType`` 
  +   ``SourceDBClusterIdentifier`` 
  +   ``SourceRegion`` 
  +  ``StorageEncrypted`` (for an encrypted snapshot)
  +   ``UseLatestRestorableTime`` 
  
 Constraints:
  +  Must match the identifier of an existing Snapshot.
  
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `source_db_cluster_identifier` (String) When restoring a DB cluster to a point in time, the identifier of the source DB cluster from which to restore.
 Constraints:
  +  Must match the identifier of an existing DBCluster.
  
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `source_region` (String) The AWS Region which contains the source DB cluster when replicating a DB cluster. For example, ``us-east-1``. 
 Valid for: Aurora DB clusters only
- `storage_encrypted` (Boolean) Indicates whether the DB cluster is encrypted.
 If you specify the ``KmsKeyId`` property, then you must enable encryption.
 If you specify the ``SourceDBClusterIdentifier`` property, don't specify this property. The value is inherited from the source DB cluster, and if the DB cluster is encrypted, the specified ``KmsKeyId`` property is used.
 If you specify the ``SnapshotIdentifier`` and the specified snapshot is encrypted, don't specify this property. The value is inherited from the snapshot, and the specified ``KmsKeyId`` property is used.
 If you specify the ``SnapshotIdentifier`` and the specified snapshot isn't encrypted, you can use this property to specify that the restored DB cluster is encrypted. Specify the ``KmsKeyId`` property for the KMS key to use for encryption. If you don't want the restored DB cluster to be encrypted, then don't set this property or set it to ``false``.
  If you specify both the ``StorageEncrypted`` and ``SnapshotIdentifier`` properties without specifying the ``KmsKeyId`` property, then the restored DB cluster inherits the encryption settings from the DB snapshot that provide.
  Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `storage_type` (String) The storage type to associate with the DB cluster.
 For information on storage types for Aurora DB clusters, see [Storage configurations for Amazon Aurora DB clusters](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.StorageReliability.html#aurora-storage-type). For information on storage types for Multi-AZ DB clusters, see [Settings for creating Multi-AZ DB clusters](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/create-multi-az-db-cluster.html#create-multi-az-db-cluster-settings).
 This setting is required to create a Multi-AZ DB cluster.
 When specified for a Multi-AZ DB cluster, a value for the ``Iops`` parameter is required.
 Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters
 Valid Values:
  +  Aurora DB clusters - ``aurora | aurora-iopt1``
  +  Multi-AZ DB clusters - ``io1 | io2 | gp3``
  
 Default:
  +  Aurora DB clusters - ``aurora``
  +  Multi-AZ DB clusters - ``io1``
  
  When you create an Aurora DB cluster with the storage type set to ``aurora-iopt1``, the storage type is returned in the response. The storage type isn't returned when you set it to ``aurora``.
- `tags` (Attributes Set) Tags to assign to the DB cluster.
 Valid for Cluster Type: Aurora DB clusters and Multi-AZ DB clusters (see [below for nested schema](#nestedatt--tags))
- `use_latest_restorable_time` (Boolean) A value that indicates whether to restore the DB cluster to the latest restorable backup time. By default, the DB cluster is not restored to the latest restorable backup time. 
 Valid for: Aurora DB clusters and Multi-AZ DB clusters
- `vpc_security_group_ids` (List of String) A list of EC2 VPC security groups to associate with this DB cluster.
 If you plan to update the resource, don't specify VPC security groups in a shared VPC.
 Valid for: Aurora DB clusters and Multi-AZ DB clusters

### Read-Only

- `db_cluster_arn` (String)
- `db_cluster_resource_id` (String)
- `endpoint` (Attributes) The ``Endpoint`` return value specifies the connection endpoint for the primary instance of the DB cluster. (see [below for nested schema](#nestedatt--endpoint))
- `id` (String) Uniquely identifies the resource.
- `read_endpoint` (Attributes) The ``ReadEndpoint`` return value specifies the reader endpoint for the DB cluster.
 The reader endpoint for a DB cluster load-balances connections across the Aurora Replicas that are available in a DB cluster. As clients request new connections to the reader endpoint, Aurora distributes the connection requests among the Aurora Replicas in the DB cluster. This functionality can help balance your read workload across multiple Aurora Replicas in your DB cluster.
 If a failover occurs, and the Aurora Replica that you are connected to is promoted to be the primary instance, your connection is dropped. To continue sending your read workload to other Aurora Replicas in the cluster, you can then reconnect to the reader endpoint.
 For more information about Aurora endpoints, see [Amazon Aurora connection management](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Overview.Endpoints.html) in the *Amazon Aurora User Guide*. (see [below for nested schema](#nestedatt--read_endpoint))
- `storage_throughput` (Number)

<a id="nestedatt--associated_roles"></a>
### Nested Schema for `associated_roles`

Optional:

- `feature_name` (String) The name of the feature associated with the AWS Identity and Access Management (IAM) role. IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf. For the list of supported feature names, see the ``SupportedFeatureNames`` description in [DBEngineVersion](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html) in the *Amazon RDS API Reference*.
- `role_arn` (String) The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.


<a id="nestedatt--master_user_secret"></a>
### Nested Schema for `master_user_secret`

Optional:

- `kms_key_id` (String) The AWS KMS key identifier that is used to encrypt the secret.

Read-Only:

- `secret_arn` (String) The Amazon Resource Name (ARN) of the secret. This parameter is a return value that you can retrieve using the ``Fn::GetAtt`` intrinsic function. For more information, see [Return values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbcluster.html#aws-resource-rds-dbcluster-return-values).


<a id="nestedatt--scaling_configuration"></a>
### Nested Schema for `scaling_configuration`

Optional:

- `auto_pause` (Boolean) Indicates whether to allow or disallow automatic pause for an Aurora DB cluster in ``serverless`` DB engine mode. A DB cluster can be paused only when it's idle (it has no connections).
  If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot. In this case, the DB cluster is restored when there is a request to connect to it.
- `max_capacity` (Number) The maximum capacity for an Aurora DB cluster in ``serverless`` DB engine mode.
 For Aurora MySQL, valid capacity values are ``1``, ``2``, ``4``, ``8``, ``16``, ``32``, ``64``, ``128``, and ``256``.
 For Aurora PostgreSQL, valid capacity values are ``2``, ``4``, ``8``, ``16``, ``32``, ``64``, ``192``, and ``384``.
 The maximum capacity must be greater than or equal to the minimum capacity.
- `min_capacity` (Number) The minimum capacity for an Aurora DB cluster in ``serverless`` DB engine mode.
 For Aurora MySQL, valid capacity values are ``1``, ``2``, ``4``, ``8``, ``16``, ``32``, ``64``, ``128``, and ``256``.
 For Aurora PostgreSQL, valid capacity values are ``2``, ``4``, ``8``, ``16``, ``32``, ``64``, ``192``, and ``384``.
 The minimum capacity must be less than or equal to the maximum capacity.
- `seconds_before_timeout` (Number) The amount of time, in seconds, that Aurora Serverless v1 tries to find a scaling point to perform seamless scaling before enforcing the timeout action. The default is 300.
 Specify a value between 60 and 600 seconds.
- `seconds_until_auto_pause` (Number) The time, in seconds, before an Aurora DB cluster in ``serverless`` mode is paused.
 Specify a value between 300 and 86,400 seconds.
- `timeout_action` (String) The action to take when the timeout is reached, either ``ForceApplyCapacityChange`` or ``RollbackCapacityChange``.
 ``ForceApplyCapacityChange`` sets the capacity to the specified value as soon as possible.
 ``RollbackCapacityChange``, the default, ignores the capacity change if a scaling point isn't found in the timeout period.
  If you specify ``ForceApplyCapacityChange``, connections that prevent Aurora Serverless v1 from finding a scaling point might be dropped.
  For more information, see [Autoscaling for Aurora Serverless v1](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.auto-scaling) in the *Amazon Aurora User Guide*.


<a id="nestedatt--serverless_v2_scaling_configuration"></a>
### Nested Schema for `serverless_v2_scaling_configuration`

Optional:

- `max_capacity` (Number) The maximum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster. You can specify ACU values in half-step increments, such as 40, 40.5, 41, and so on. The largest value that you can use is 128.
 The maximum capacity must be higher than 0.5 ACUs. For more information, see [Choosing the maximum Aurora Serverless v2 capacity setting for a cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.setting-capacity.html#aurora-serverless-v2.max_capacity_considerations) in the *Amazon Aurora User Guide*.
 Aurora automatically sets certain parameters for Aurora Serverless V2 DB instances to values that depend on the maximum ACU value in the capacity range. When you update the maximum capacity value, the ``ParameterApplyStatus`` value for the DB instance changes to ``pending-reboot``. You can update the parameter values by rebooting the DB instance after changing the capacity range.
- `min_capacity` (Number) The minimum number of Aurora capacity units (ACUs) for a DB instance in an Aurora Serverless v2 cluster. You can specify ACU values in half-step increments, such as 8, 8.5, 9, and so on. For Aurora versions that support the Aurora Serverless v2 auto-pause feature, the smallest value that you can use is 0. For versions that don't support Aurora Serverless v2 auto-pause, the smallest value that you can use is 0.5.
- `seconds_until_auto_pause` (Number) Specifies the number of seconds an Aurora Serverless v2 DB instance must be idle before Aurora attempts to automatically pause it. 
 Specify a value between 300 seconds (five minutes) and 86,400 seconds (one day). The default is 300 seconds.


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Optional:

- `key` (String) A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").
- `value` (String) A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$").


<a id="nestedatt--endpoint"></a>
### Nested Schema for `endpoint`

Read-Only:

- `address` (String) Specifies the connection endpoint for the primary instance of the DB cluster.
- `port` (String) Specifies the port that the database engine is listening on.


<a id="nestedatt--read_endpoint"></a>
### Nested Schema for `read_endpoint`

Read-Only:

- `address` (String) The host address of the reader endpoint.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_rds_db_cluster.example "db_cluster_identifier"
```