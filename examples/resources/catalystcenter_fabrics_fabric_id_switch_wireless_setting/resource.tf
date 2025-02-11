
resource "catalystcenter_fabrics_fabric_id_switch_wireless_setting" "example" {
    provider = catalystcenter

    parameters {

      enable_wireless = "false"
      fabric_id = "string"
      id = "string"
      rolling_ap_upgrade {

        ap_reboot_percentage = 1
        enable_rolling_ap_upgrade = "false"
      }
    }
}

output "catalystcenter_fabrics_fabric_id_switch_wireless_setting_example" {
    value = catalystcenter_fabrics_fabric_id_switch_wireless_setting.example
}
