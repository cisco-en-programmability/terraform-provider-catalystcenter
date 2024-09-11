
data "catalystcenter_network_device_supervisor_card_details" "example" {
  provider    = catalystcenter
  device_uuid = "string"
}

output "catalystcenter_network_device_supervisor_card_details_example" {
  value = data.catalystcenter_network_device_supervisor_card_details.example.items
}
