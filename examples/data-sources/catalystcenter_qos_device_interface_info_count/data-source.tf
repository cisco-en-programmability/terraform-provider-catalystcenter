
data "catalystcenter_qos_device_interface_info_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_qos_device_interface_info_count_example" {
  value = data.catalystcenter_qos_device_interface_info_count.example.item
}
