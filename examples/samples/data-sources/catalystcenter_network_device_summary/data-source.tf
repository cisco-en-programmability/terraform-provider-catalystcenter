
data "catalystcenter_network_device_summary" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_network_device_summary_example" {
  value = data.catalystcenter_network_device_summary.example.item
}
