---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_personalize_dataset Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Personalize::Dataset
---

# awscc_personalize_dataset (Data Source)

Data Source schema for AWS::Personalize::Dataset



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `dataset_arn` (String) The ARN of the dataset
- `dataset_group_arn` (String) The Amazon Resource Name (ARN) of the dataset group to add the dataset to
- `dataset_import_job` (Attributes) Initial DatasetImportJob for the created dataset (see [below for nested schema](#nestedatt--dataset_import_job))
- `dataset_type` (String) The type of dataset
- `name` (String) The name for the dataset
- `schema_arn` (String) The ARN of the schema to associate with the dataset. The schema defines the dataset fields.

<a id="nestedatt--dataset_import_job"></a>
### Nested Schema for `dataset_import_job`

Read-Only:

- `data_source` (Attributes) The Amazon S3 bucket that contains the training data to import. (see [below for nested schema](#nestedatt--dataset_import_job--data_source))
- `dataset_arn` (String) The ARN of the dataset that receives the imported data
- `dataset_import_job_arn` (String) The ARN of the dataset import job
- `job_name` (String) The name for the dataset import job.
- `role_arn` (String) The ARN of the IAM role that has permissions to read from the Amazon S3 data source.

<a id="nestedatt--dataset_import_job--data_source"></a>
### Nested Schema for `dataset_import_job.data_source`

Read-Only:

- `data_location` (String) The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.
