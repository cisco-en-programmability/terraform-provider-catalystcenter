
data "catalystcenter_network_device_config" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_device_config_example" {
  value = data.catalystcenter_network_device_config.example.items
}

data "catalystcenter_network_device_config" "example" {
  provider          = catalystcenter
  network_device_id = "string"
}

output "catalystcenter_network_device_config_example" {
  value = data.catalystcenter_network_device_config.example.item
}
