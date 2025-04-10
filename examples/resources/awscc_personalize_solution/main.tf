# First, create a dataset group
resource "awscc_personalize_dataset_group" "example" {
  name = "example-dataset-group"
}

# Create a schema for the dataset
resource "awscc_personalize_schema" "example" {
  name = "example-schema"
  schema = jsonencode({
    type      = "record"
    name      = "Interactions"
    namespace = "com.amazonaws.personalize.schema"
    fields = [
      {
        name = "USER_ID"
        type = "string"
      },
      {
        name = "ITEM_ID"
        type = "string"
      },
      {
        name = "TIMESTAMP"
        type = "long"
      }
    ]
    version = "1.0"
  })
}

# Create a dataset
resource "awscc_personalize_dataset" "example" {
  name              = "example-dataset"
  dataset_group_arn = awscc_personalize_dataset_group.example.dataset_group_arn
  dataset_type      = "Interactions"
  schema_arn        = awscc_personalize_schema.example.schema_arn
}

# Create the solution
resource "awscc_personalize_solution" "example" {
  name              = "example-solution"
  dataset_group_arn = awscc_personalize_dataset_group.example.dataset_group_arn
  recipe_arn        = "arn:aws:personalize:::recipe/aws-user-personalization"

  solution_config = {
    algorithm_hyper_parameters = {
      "hidden_dimension" = "64"
    }
  }

  depends_on = [awscc_personalize_dataset.example]
}