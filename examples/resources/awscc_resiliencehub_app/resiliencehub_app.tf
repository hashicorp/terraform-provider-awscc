resource "awscc_resiliencehub_app" "example" {
    name                  = "example-app"
    description           = "This is a test application"
    resiliency_policy_arn = awscc_resiliencehub_resiliency_policy.test.arn
    app_template_body     = jsondecode({
        resources = [{
                type = "AWS::EKS::Deployment",
                logicalResourceId = {
                    eksSourceName  = "EKS-Cluster/Deployment-Store",
                    identifier = "MyEKSCluster-store"
                }
                name = "eksdeployment",
                parameters = {
                     tags = {
                     env = "dev"
                    }
                }
            }],
        appComponent = [{
                id   = "AppComponent-EKSDeployment-Store",
                name = "AppComponent-EKSDeployment-Store",
                resourceNames = [
                    "eksdeployment"
                ],
                type = "AWS::ResilienceHub::ComputeAppComponent"
            }],
        excludedResources = {
            logicalResourceIds = []
        },
        version = 2.0
})    
    event_subscriptions = [{
                eventType = "DriftDetected"
                name = "test-app"
                snsTopicArn = aws_sns_topic.test.arn
            }]
    app_assessment_schedule = "Daily"
    permission_model = {
        type              = "RoleBased"
        invoker_role_name = "resilience-hub-test-app-role"
        }
    resource_mappings       = {
            eks_source_name       = "EKS-Cluster/Deployment-Store"
            mapping_type          = "EKS"
            logical_resource_id   = "MyEKSCluster-store"
            physical_resource_id  = {  
                  aws_account_id  = "11222333344"
                  aws_region      = "us-west-2"
                  identifier      = "arn:aws:eks:us-west-2:11222333344:cluster/EKS-Cluster/Deployment-Store"
                  type            = "Arn"
        }
    }
    tags  = {
          key   = "Modified By"
          value = "AWSCC"
        }
  }
