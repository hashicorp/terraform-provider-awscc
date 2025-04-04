# Example ActiveMQ configuration with basic settings
resource "awscc_amazonmq_configuration" "example" {
  name           = "example-mq-config"
  engine_type    = "ACTIVEMQ"
  engine_version = "5.17.6"
  description    = "Example Amazon MQ configuration managed by Terraform"
  data = base64encode(<<-EOT
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<broker xmlns="http://activemq.apache.org/schema/core">
  <destinationPolicy>
    <policyMap>
      <policyEntries>
        <policyEntry topic=">" producerFlowControl="true">
          <pendingMessageLimitStrategy>
            <constantPendingMessageLimitStrategy limit="1000"/>
          </pendingMessageLimitStrategy>
        </policyEntry>
        <policyEntry queue=">" producerFlowControl="true" memoryLimit="1mb"/>
      </policyEntries>
    </policyMap>
  </destinationPolicy>
</broker>
EOT
  )

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}