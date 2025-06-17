
resource "catalystcenter_fabrics_fabric_id_switch_wireless_setting_reload" "example" {
  provider  = catalystcenter
  fabric_id = "string"
  parameters = [{

    device_id = "string"
  }]
}

output "catalystcenter_fabrics_fabric_id_switch_wireless_setting_reload_example" {
  value = catalystcenter_fabrics_fabric_id_switch_wireless_setting_reload.example
}