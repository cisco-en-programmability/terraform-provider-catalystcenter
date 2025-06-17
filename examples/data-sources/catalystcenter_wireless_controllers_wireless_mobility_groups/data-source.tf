
data "catalystcenter_wireless_controllers_wireless_mobility_groups" "example" {
  provider          = catalystcenter
  network_device_id = "string"
}

output "catalystcenter_wireless_controllers_wireless_mobility_groups_example" {
  value = data.catalystcenter_wireless_controllers_wireless_mobility_groups.example.items
}
