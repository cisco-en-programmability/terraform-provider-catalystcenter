
data "catalystcenter_qos_device_interface" "example" {
  provider          = catalystcenter
  network_device_id = "string"
}

output "catalystcenter_qos_device_interface_example" {
  value = data.catalystcenter_qos_device_interface.example.items
}
