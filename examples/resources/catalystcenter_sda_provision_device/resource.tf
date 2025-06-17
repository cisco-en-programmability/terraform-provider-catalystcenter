
resource "catalystcenter_sda_provision_device" "example" {
  provider = catalystcenter
 
  parameters {

    device_management_ip_address = "string"
    site_name_hierarchy          = "string"
  }
}

output "catalystcenter_sda_provision_device_example" {
  value = catalystcenter_sda_provision_device.example
}