
resource "catalystcenter_business_sda_wireless_controller_create" "example" {
  provider = catalystcenter
  parameters = [{

    device_name         = "string"
    site_name_hierarchy = "string"
  }]
}

output "catalystcenter_business_sda_wireless_controller_create_example" {
  value = catalystcenter_business_sda_wireless_controller_create.example
}