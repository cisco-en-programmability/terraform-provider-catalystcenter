
data "catalystcenter_discovery_device" "example" {
  provider = catalystcenter
  id       = "string"
  task_id  = "string"
}

output "catalystcenter_discovery_device_example" {
  value = data.catalystcenter_discovery_device.example.items
}
