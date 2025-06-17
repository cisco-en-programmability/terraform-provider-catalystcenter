
data "catalystcenter_sda_port_assignment_for_access_point" "example" {
  provider                     = catalystcenter
  device_management_ip_address = "string"
  interface_name               = "string"
}

output "catalystcenter_sda_port_assignment_for_access_point_example" {
  value = data.catalystcenter_sda_port_assignment_for_access_point.example.item
}
