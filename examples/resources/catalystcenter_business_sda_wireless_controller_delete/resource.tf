
resource "catalystcenter_business_sda_wireless_controller_delete" "example" {
  provider         = catalystcenter
  device_ipaddress = "string"
}

output "catalystcenter_business_sda_wireless_controller_delete_example" {
  value = catalystcenter_business_sda_wireless_controller_delete.example
}