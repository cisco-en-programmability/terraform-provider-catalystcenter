
data "catalystcenter_sda_port_channels_count" "example" {
  provider              = catalystcenter
  connected_device_type = "string"
  fabric_id             = "string"
  network_device_id     = "string"
  port_channel_name     = "string"
}

output "catalystcenter_sda_port_channels_count_example" {
  value = data.catalystcenter_sda_port_channels_count.example.item
}
