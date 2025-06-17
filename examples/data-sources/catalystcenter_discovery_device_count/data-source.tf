
data "catalystcenter_discovery_device_count" "example" {
  provider = catalystcenter
  id       = "string"
  task_id  = "string"
}

output "catalystcenter_discovery_device_count_example" {
  value = data.catalystcenter_discovery_device_count.example.item
}
