
data "catalystcenter_sda_port_assignment_for_user_device" "example" {
  provider                     = catalystcenter
  device_management_ip_address = "string"
  interface_name               = "string"
}

output "catalystcenter_sda_port_assignment_for_user_device_example" {
  value = data.catalystcenter_sda_port_assignment_for_user_device.example.item
}
