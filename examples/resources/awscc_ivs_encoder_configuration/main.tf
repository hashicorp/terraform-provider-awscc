# Create IVS Encoder Configuration
resource "awscc_ivs_encoder_configuration" "example" {
  name = "example-encoder-config"
  video = {
    bitrate   = 3000000 # 3 Mbps
    framerate = 60      # 60 fps
    height    = 1080    # 1080p
    width     = 1920    # 1920px
  }
  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}