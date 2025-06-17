
data "catalystcenter_network_device_functional_capability" "example" {
  provider      = catalystcenter
  device_id     = "string"
  function_name = ["string"]
}

output "catalystcenter_network_device_functional_capability_example" {
  value = data.catalystcenter_network_device_functional_capability.example.items
}

data "catalystcenter_network_device_functional_capability" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_network_device_functional_capability_example" {
  value = data.catalystcenter_network_device_functional_capability.example.item
}
