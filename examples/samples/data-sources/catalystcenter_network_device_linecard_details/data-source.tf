
data "catalystcenter_network_device_linecard_details" "example" {
  provider    = catalystcenter
  device_uuid = "string"
}

output "catalystcenter_network_device_linecard_details_example" {
  value = data.catalystcenter_network_device_linecard_details.example.items
}
