
resource "catalystcenter_network_device_config_write_memory" "example" {
  provider = catalystcenter
  parameters {

    device_id = ["string"]
  }
}

output "catalystcenter_network_device_config_write_memory_example" {
  value = catalystcenter_network_device_config_write_memory.example
}