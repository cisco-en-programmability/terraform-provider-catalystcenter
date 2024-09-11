
data "catalystcenter_wireless_settings_interfaces_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_wireless_settings_interfaces_count_example" {
  value = data.catalystcenter_wireless_settings_interfaces_count.example.item
}
