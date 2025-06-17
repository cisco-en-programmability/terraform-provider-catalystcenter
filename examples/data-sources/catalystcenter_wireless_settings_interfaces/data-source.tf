
data "catalystcenter_wireless_settings_interfaces" "example" {
  provider       = catalystcenter
  interface_name = "string"
  limit          = 1
  offset         = 1
  vlan_id        = 1.0
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
