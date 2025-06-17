
data "catalystcenter_network_device_config_files" "example" {
  provider          = catalystcenter
  file_type         = "string"
  id                = "string"
  limit             = 1
  network_device_id = "string"
  offset            = 1
}

output "catalystcenter_network_device_config_files_example" {
  value = data.catalystcenter_network_device_config_files.example.items
}
