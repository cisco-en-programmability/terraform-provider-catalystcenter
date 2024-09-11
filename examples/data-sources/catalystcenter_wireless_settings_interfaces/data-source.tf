
data "catalystcenter_wireless_settings_interfaces" "example" {
  provider = catalystcenter
  limit    = 1
  offset   = 1
}

output "catalystcenter_wireless_settings_interfaces_example" {
  value = data.catalystcenter_wireless_settings_interfaces.example.items
}

data "catalystcenter_wireless_settings_interfaces" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_wireless_settings_interfaces_example" {
  value = data.catalystcenter_wireless_settings_interfaces.example.item
}
