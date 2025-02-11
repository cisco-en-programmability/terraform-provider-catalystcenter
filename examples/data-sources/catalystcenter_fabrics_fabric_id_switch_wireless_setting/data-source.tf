
data "catalystcenter_fabrics_fabric_id_switch_wireless_setting" "example" {
    provider = catalystcenter
    fabric_id = "string"
}

output "catalystcenter_fabrics_fabric_id_switch_wireless_setting_example" {
    value = data.catalystcenter_fabrics_fabric_id_switch_wireless_setting.example.items
}
