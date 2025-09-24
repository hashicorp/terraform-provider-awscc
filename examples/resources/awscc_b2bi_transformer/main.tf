# B2BI Transformer example
resource "awscc_b2bi_transformer" "example" {
  name   = "example-transformer"
  status = "active"

  file_format      = "JSON"
  mapping_template = "{ \"order\": { \"po_number\": {{BEG03}}, \"date\": {{BEG05}} } }"
  edi_type = {
    x12_details = {
      transaction_set = "X12_850"
      version         = "VERSION_4010"
    }
  }
  sample_document = "ISA*00*          *00*          *ZZ*YOUR_ISA     *ZZ*THEIR_ISA    *210120*1200*U*00401*000000001*0*P*>~GS*PO*YOUR_ID*THEIR_ID*20210120*1200*1*X*004010~ST*850*0001~BEG*00*SS*123456789*20210120~CTT*1~SE*4*0001~GE*1*1~IEA*1*000000001~"

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}