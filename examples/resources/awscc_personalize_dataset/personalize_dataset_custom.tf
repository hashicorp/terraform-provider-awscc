resource "awscc_personalize_dataset_group" "personalize_dataset_group" {
  name = "TestPresonalizeDatasetGroup"
}

resource "awscc_personalize_schema" "interactions_schema" {
  name   = "interactions_dataset_schema"
  schema = "{\n \"type\": \"record\",\n \"namespace\": \"com.amazonaws.personalize.schema\",\n \"name\": \"Interactions\",\n \"fields\": [\n { \"name\": \"USER_ID\", \"type\": \"string\" },\n { \"name\": \"ITEM_ID\", \"type\": \"string\" },\n { \"name\": \"TIMESTAMP\", \"type\": \"long\" },\n { \"name\": \"EVENT_TYPE\", \"type\": \"string\" },\n { \"name\": \"EVENT_VALUE\", \"type\": [\"float\",\"null\"]},\n { \"name\": \"IMPRESSION\", \"type\": \"string\" },\n { \"name\": \"DEVICE\", \"type\": [\"string\",\"null\"]}]}"
}

resource "awscc_personalize_dataset" "dataset_interactions" {
  dataset_group_arn = awscc_personalize_dataset_group.personalize_dataset_group.dataset_group_arn
  dataset_type      = "Interactions"
  name              = "Interactions_Dataset"
  schema_arn        = awscc_personalize_schema.interactions_schema.schema_arn
  depends_on        = [awscc_personalize_dataset_group.personalize_dataset_group, awscc_personalize_schema.interactions_schema]
}

resource "awscc_personalize_schema" "users_schema" {
  name   = "users_dataset_schema"
  schema = "{\n \"type\": \"record\",\n \"namespace\": \"com.amazonaws.personalize.schema\",\n \"name\": \"Users\",\n \"fields\": [\n { \"name\": \"USER_ID\", \"type\": \"string\" },\n { \"name\": \"AGE\", \"type\": \"int\" },\n { \"name\": \"GENDER\", \"type\": \"string\",\"categorical\": true }\n ]\n }"
}

resource "awscc_personalize_dataset" "dataset_users" {
  dataset_group_arn = awscc_personalize_dataset_group.personalize_dataset_group.dataset_group_arn
  dataset_type      = "Users"
  name              = "Users_Dataset"
  schema_arn        = awscc_personalize_schema.users_schema.schema_arn
  depends_on        = [awscc_personalize_dataset_group.personalize_dataset_group, awscc_personalize_schema.users_schema]
}

resource "awscc_personalize_schema" "items_schema" {
  name   = "items_dataset_schema"
  schema = "{\n \"type\": \"record\",\n \"namespace\": \"com.amazonaws.personalize.schema\",\n \"name\": \"Items\",\n \"fields\": [\n { \"name\": \"ITEM_ID\", \"type\": \"string\" },\n { \"name\": \"GENRES\", \"type\": [\"null\", \"string\" ], \"categorical\": true},\n { \"name\": \"DESCRIPTION\", \"type\": [\"null\", \"string\" ], \"textual\": true },\n { \"name\": \"CREATION_TIMESTAMP\", \"type\": \"long\"}]\n }"
}

resource "awscc_personalize_dataset" "dataset_items" {
  dataset_group_arn = awscc_personalize_dataset_group.personalize_dataset_group.dataset_group_arn
  dataset_type      = "Items"
  name              = "Items_Dataset"
  schema_arn        = awscc_personalize_schema.items_schema.schema_arn
  depends_on        = [awscc_personalize_dataset_group.personalize_dataset_group, awscc_personalize_schema.items_schema]
}