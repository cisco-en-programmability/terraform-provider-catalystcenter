
data "catalystcenter_network_device_config_" "example" {
  provider     = catalystcenter
  created_by   = "string"
  created_time = "string"
  device_id    = "string"
  file_type    = "string"
  limit        = 1
  offset       = 1
}

output "catalystcenter_network_device_config__example" {
  value = data.catalystcenter_network_device_config_.example.items
}
