
data "catalystcenter_network_device_register_for_wsa" "example" {
  provider      = catalystcenter
  macaddress    = "string"
  serial_number = "string"
}

output "catalystcenter_network_device_register_for_wsa_example" {
  value = data.catalystcenter_network_device_register_for_wsa.example.item
}
