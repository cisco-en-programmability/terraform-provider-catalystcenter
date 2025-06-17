
data "catalystcenter_network_devices_id" "example" {
  provider = catalystcenter
  id       = "string"
  views    = "string"
}

output "catalystcenter_network_devices_id_example" {
  value = data.catalystcenter_network_devices_id.example.item
}
