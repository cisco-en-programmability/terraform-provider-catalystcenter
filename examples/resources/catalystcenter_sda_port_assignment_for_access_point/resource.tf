
resource "catalystcenter_sda_port_assignment_for_access_point" "example" {
  provider = catalystcenter

  parameters {

    authenticate_template_name   = "string"
    data_ip_address_pool_name    = "string"
    device_management_ip_address = "string"
    interface_description        = "string"
    interface_name               = "string"
    site_name_hierarchy          = "string"
  }
}

output "catalystcenter_sda_port_assignment_for_access_point_example" {
  value = catalystcenter_sda_port_assignment_for_access_point.example
}