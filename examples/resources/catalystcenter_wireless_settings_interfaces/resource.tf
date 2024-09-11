
resource "catalystcenter_wireless_settings_interfaces" "example" {
  provider = catalystcenter

  parameters {

    id             = "string"
    interface_name = "string"
    vlan_id        = 1
  }
}

output "catalystcenter_wireless_settings_interfaces_example" {
  value = catalystcenter_wireless_settings_interfaces.example
}