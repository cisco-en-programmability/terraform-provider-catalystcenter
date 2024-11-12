
resource "catalystcenter_wireless_controllers_wireless_mobility_groups_mobility_reset" "example" {
  provider = meraki
  parameters {

    network_device_id = "string"
  }
}

output "catalystcenter_wireless_controllers_wireless_mobility_groups_mobility_reset_example" {
  value = catalystcenter_wireless_controllers_wireless_mobility_groups_mobility_reset.example
}