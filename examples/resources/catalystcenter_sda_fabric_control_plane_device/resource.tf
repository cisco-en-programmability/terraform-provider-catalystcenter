
resource "catalystcenter_sda_fabric_control_plane_device" "example" {
  provider = catalystcenter

  parameters {

    device_management_ip_address = "string"
    route_distribution_protocol  = "string"
    site_name_hierarchy          = "string"
  }
}

output "catalystcenter_sda_fabric_control_plane_device_example" {
  value = catalystcenter_sda_fabric_control_plane_device.example
}