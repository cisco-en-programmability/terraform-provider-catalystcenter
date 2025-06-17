
data "catalystcenter_network_device_config_files_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_network_device_config_files_id_example" {
  value = data.catalystcenter_network_device_config_files_id.example.item
}
