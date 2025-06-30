---
page_title: "awscc_personalize_dataset Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Personalize::Dataset.
---

# awscc_personalize_dataset (Resource)

Resource schema for AWS::Personalize::Dataset.

## Example Usage

### Personalize dataset with required files
Creation of interactions, items and users datasets with required fields
```terraform
resource "awscc_personalize_dataset_group" "personalize_dataset_group" {
    name = "TestPresonalizeDatasetGroup"
}

resource "awscc_personalize_schema" "interactions_schema"{
    name = "interactions_dataset_schema"
    schema = "{\n \"type\": \"record\",\n \"namespace\": \"com.amazonaws.personalize.schema\",\n \"name\": \"Interactions\",\n \"fields\": [\n { \"name\": \"USER_ID\", \"type\": \"string\" },\n { \"name\": \"ITEM_ID\", \"type\": \"string\" },\n { \"name\": \"TIMESTAMP\", \"type\": \"long\" }\n ]\n }"
}

resource "awscc_personalize_dataset" "dataset_interactions"{
    dataset_group_arn = awscc_personalize_dataset_group.personalize_dataset_group.dataset_group_arn
    dataset_type = "Interactions"
    name = "Interactions_Dataset"
    schema_arn = awscc_personalize_schema.interactions_schema.schema_arn
    depends_on = [awscc_personalize_dataset_group.personalize_dataset_group, awscc_personalize_schema.interactions_schema]
}

resource "awscc_personalize_schema" "users_schema"{
    name = "users_dataset_schema"
    schema = "{\n \"type\": \"record\",\n \"namespace\": \"com.amazonaws.personalize.schema\",\n \"name\": \"Users\",\n \"fields\": [\n { \"name\": \"USER_ID\", \"type\": \"string\" },\n { \"name\": \"AGE\", \"type\": \"int\" },\n { \"name\": \"GENDER\", \"type\": \"string\",\"categorical\": true }\n ]\n }"
}

resource "awscc_personalize_dataset" "dataset_users"{
    dataset_group_arn = awscc_personalize_dataset_group.personalize_dataset_group.dataset_group_arn
    dataset_type = "Users"
    name = "Users_Dataset"
    schema_arn = awscc_personalize_schema.users_schema.schema_arn
    depends_on = [awscc_personalize_dataset_group.personalize_dataset_group, awscc_personalize_schema.users_schema]
}

resource "awscc_personalize_schema" "items_schema"{
    name = "items_dataset_schema"
    schema = "{\n \"type\": \"record\",\n \"namespace\": \"com.amazonaws.personalize.schema\",\n \"name\": \"Items\",\n \"fields\": [\n { \"name\": \"ITEM_ID\", \"type\": \"string\" },\n { \"name\": \"GENRES\", \"type\": [\"null\", \"string\" ], \"categorical\": true},\n { \"name\": \"DESCRIPTION\", \"type\": [\"null\", \"string\" ], \"textual\": true }\n ]\n }"
}

resource "awscc_personalize_dataset" "dataset_items"{
    dataset_group_arn = awscc_personalize_dataset_group.personalize_dataset_group.dataset_group_arn
    dataset_type = "Items"
    name = "Items_Dataset"
    schema_arn = awscc_personalize_schema.items_schema.schema_arn
    depends_on = [awscc_personalize_dataset_group.personalize_dataset_group, awscc_personalize_schema.items_schema]
}
```

### Personalize dataset with all possible fields
Creation of interactions, items and users datasets with more than the required fields
```terraform
resource "awscc_personalize_dataset_group" "personalize_dataset_group" {
    name = "TestPresonalizeDatasetGroup"
}

resource "awscc_personalize_schema" "interactions_schema"{
    name = "interactions_dataset_schema"
    schema = "{\n \"type\": \"record\",\n \"namespace\": \"com.amazonaws.personalize.schema\",\n \"name\": \"Interactions\",\n \"fields\": [\n { \"name\": \"USER_ID\", \"type\": \"string\" },\n { \"name\": \"ITEM_ID\", \"type\": \"string\" },\n { \"name\": \"TIMESTAMP\", \"type\": \"long\" },\n { \"name\": \"EVENT_TYPE\", \"type\": \"string\" },\n { \"name\": \"EVENT_VALUE\", \"type\": [\"float\",\"null\"]},\n { \"name\": \"IMPRESSION\", \"type\": \"string\" },\n { \"name\": \"DEVICE\", \"type\": [\"string\",\"null\"]}]}"
}

resource "awscc_personalize_dataset" "dataset_interactions"{
    dataset_group_arn = awscc_personalize_dataset_group.personalize_dataset_group.dataset_group_arn
    dataset_type = "Interactions"
    name = "Interactions_Dataset"
    schema_arn = awscc_personalize_schema.interactions_schema.schema_arn
    depends_on = [awscc_personalize_dataset_group.personalize_dataset_group, awscc_personalize_schema.interactions_schema]
}

resource "awscc_personalize_schema" "users_schema"{
    name = "users_dataset_schema"
    schema = "{\n \"type\": \"record\",\n \"namespace\": \"com.amazonaws.personalize.schema\",\n \"name\": \"Users\",\n \"fields\": [\n { \"name\": \"USER_ID\", \"type\": \"string\" },\n { \"name\": \"AGE\", \"type\": \"int\" },\n { \"name\": \"GENDER\", \"type\": \"string\",\"categorical\": true }\n ]\n }"
}

resource "awscc_personalize_dataset" "dataset_users"{
    dataset_group_arn = awscc_personalize_dataset_group.personalize_dataset_group.dataset_group_arn
    dataset_type = "Users"
    name = "Users_Dataset"
    schema_arn = awscc_personalize_schema.users_schema.schema_arn
    depends_on = [awscc_personalize_dataset_group.personalize_dataset_group, awscc_personalize_schema.users_schema]
}

resource "awscc_personalize_schema" "items_schema"{
    name = "items_dataset_schema"
    schema = "{\n \"type\": \"record\",\n \"namespace\": \"com.amazonaws.personalize.schema\",\n \"name\": \"Items\",\n \"fields\": [\n { \"name\": \"ITEM_ID\", \"type\": \"string\" },\n { \"name\": \"GENRES\", \"type\": [\"null\", \"string\" ], \"categorical\": true},\n { \"name\": \"DESCRIPTION\", \"type\": [\"null\", \"string\" ], \"textual\": true },\n { \"name\": \"CREATION_TIMESTAMP\", \"type\": \"long\"}]\n }"
}

resource "awscc_personalize_dataset" "dataset_items"{
    dataset_group_arn = awscc_personalize_dataset_group.personalize_dataset_group.dataset_group_arn
    dataset_type = "Items"
    name = "Items_Dataset"
    schema_arn = awscc_personalize_schema.items_schema.schema_arn
    depends_on = [awscc_personalize_dataset_group.personalize_dataset_group, awscc_personalize_schema.items_schema]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `dataset_group_arn` (String) The Amazon Resource Name (ARN) of the dataset group to add the dataset to
- `dataset_type` (String) The type of dataset
- `name` (String) The name for the dataset
- `schema_arn` (String) The ARN of the schema to associate with the dataset. The schema defines the dataset fields.

### Optional

- `dataset_import_job` (Attributes) Initial DatasetImportJob for the created dataset (see [below for nested schema](#nestedatt--dataset_import_job))

### Read-Only

- `dataset_arn` (String) The ARN of the dataset
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--dataset_import_job"></a>
### Nested Schema for `dataset_import_job`

Optional:

- `data_source` (Attributes) The Amazon S3 bucket that contains the training data to import. (see [below for nested schema](#nestedatt--dataset_import_job--data_source))
- `dataset_arn` (String) The ARN of the dataset that receives the imported data
- `dataset_import_job_arn` (String) The ARN of the dataset import job
- `job_name` (String) The name for the dataset import job.
- `role_arn` (String) The ARN of the IAM role that has permissions to read from the Amazon S3 data source.

<a id="nestedatt--dataset_import_job--data_source"></a>
### Nested Schema for `dataset_import_job.data_source`

Optional:

- `data_location` (String) The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_personalize_dataset.example
  id = "dataset_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_personalize_dataset.example "dataset_arn"
```