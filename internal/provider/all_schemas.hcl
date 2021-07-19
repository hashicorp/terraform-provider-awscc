meta_schema {
  source {
    url = "https://raw.githubusercontent.com/aws-cloudformation/aws-cloudformation-resource-schema/v2.0.4/src/main/resources/schema/provider.definition.schema.v1.json"
  }

  local = "../service/cloudformation/meta-schemas/provider.definition.schema.v1.json"
}

resource_schema "aws_logs_log_group" {
  source {
    url = "https://raw.githubusercontent.com/aws-cloudformation/aws-cloudformation-resource-providers-logs/8b9229f78b832800b7fb6c1165bcb3893f44b856/aws-logs-loggroup/aws-logs-loggroup.json"
  }

  local = "../service/cloudformation/schemas/us-west-2/aws-logs-loggroup.json"
}

resource_schema "aws_appmesh_virtual_service" {
  source {
    url = "???"
  }

  local = "../service/cloudformation/schemas/us-west-2/aws-appmesh-virtualservice.json"
}

resource_schema "aws_synthetics_canary" {
  source {
    url = "???"
  }

  local = "../service/cloudformation/schemas/us-west-2/aws-synthetics-canary.json"
}

resource_schema "aws_backup_backup_plan" {
  source {
    url = "???"
  }

  local = "../service/cloudformation/schemas/us-west-2/aws-backup-backupplan.json"
}

#resource_schema "aws_sagemaker_data_quality_job_definition" {
#  source {
#    url = "???"
#  }
#
#  local = "../service/cloudformation/schemas/us-west-2/aws-sagemaker-dataqualityjobdefinition.json"
#}

#resource_schema "aws_stepfunctions_state_machine" {
#  source {
#    url = "???"
#  }
#
#  local = "../service/cloudformation/schemas/us-west-2/aws-stepfunctions-statemachine.json"
#}

#resource_schema "aws_xray_sampling_rule" {
#  source {
#    url = "???"
#  }
#
#  local = "../service/cloudformation/schemas/us-west-2/aws-xray-samplingrule.json"
#}

#resource_schema "aws_athena_workgroup" {
#  source {
#    url = "https://raw.githubusercontent.com/aws-cloudformation/aws-cloudformation-resource-providers-athena/a2b8f560891e1256fd1ba965e9ffe3dcddde1855/workgroup/aws-athena-workgroup.json"
#  }
#
#  local = "aws_athena_workgroup.cf-resource-schema.json"
#}

#resource_schema "aws_ecs_task_definition" {
#  source {
#    url = "???"
#  }
#
# local = "../service/cloudformation/schemas/us-west-2/aws-ecs-taskdefinition.json"
#}