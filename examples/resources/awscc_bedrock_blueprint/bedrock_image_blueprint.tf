
resources "awscc_bedrockdataautomation_blueprint" "example" {
    blueprint_name = "EXAMPLE"
    type = "DOCUMENT"
    blueprint_stage = "LIVE"
    schema = 
}


bda_create_blueprint_response = bedrock_data_automation_client.create_blueprint(
    blueprintName='CUSTOM_PAYSLIP_BLUEPRINT',
    type='DOCUMENT',
    blueprintStage='LIVE',
    schema=json.dumps({
        '$schema': 'http://json-schema.org/draft-07/schema#',
        'description': 'default',
        'documentClass': 'default',
        'type': 'object',
        'properties': {
            'gross_pay_this_period': {
                'type': 'number',
                'inferenceType': 'extractive',
                'description': 'The gross pay for this pay period from the Earnings table'
            },
            'net_pay': {
                'type': 'number',
                'inferenceType': 'extractive',
                'description': 'The net pay for this pay period from the bottom of the document'
            }
        }
    }),
)