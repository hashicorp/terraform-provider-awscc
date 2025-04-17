# Create an IoT Wireless Device Profile
resource "awscc_iotwireless_device_profile" "example" {
  name = "example-device-profile"

  lo_ra_wan = {
    mac_version           = "1.0.3"
    reg_params_revision   = "RP002-1.0.1"
    rf_region             = "US915"
    supports_class_b      = false
    supports_class_c      = true
    supports_join         = true
    max_eirp              = 10
    max_duty_cycle        = 1
    supports_32_bit_f_cnt = true
    rx_delay_1            = 1
    rx_dr_offset_1        = 0
    rx_data_rate_2        = 8
    rx_freq_2             = 9233000 # Frequency in kHz (923.3 MHz)
  }

  tags = [{
    key   = "Modified By"
    value = "AWSCC"
  }]
}