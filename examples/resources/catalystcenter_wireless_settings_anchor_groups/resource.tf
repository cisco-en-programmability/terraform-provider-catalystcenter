
resource "catalystcenter_wireless_settings_anchor_groups" "example" {
    provider = catalystcenter

    parameters {

      anchor_group_name = "string"
      mobility_anchors {

        anchor_priority = "string"
        device_name = "string"
        ip_address = "string"
        mac_address = "string"
        managed_anchor_wlc = "false"
        mobility_group_name = "string"
        peer_device_type = "string"
        private_ip = "string"
      }
    }
}

output "catalystcenter_wireless_settings_anchor_groups_example" {
    value = catalystcenter_wireless_settings_anchor_groups.example
}
