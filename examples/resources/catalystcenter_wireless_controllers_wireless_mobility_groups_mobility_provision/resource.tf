
resource "catalystcenter_wireless_controllers_wireless_mobility_groups_mobility_provision" "example" {
  provider = catalystcenter
  parameters {

    data_link_encryption = "false"
    dtls_high_cipher     = "false"
    mac_address          = "string"
    management_ip        = "string"
    mobility_group_name  = "string"
    mobility_peers {

      device_series          = "string"
      hash_key               = "string"
      member_mac_address     = "string"
      mobility_group_name    = "string"
      peer_device_name       = "string"
      peer_ip                = "string"
      peer_network_device_id = "string"
      private_ip_address     = "string"
    }
    network_device_id = "string"
  }
}

output "catalystcenter_wireless_controllers_wireless_mobility_groups_mobility_provision_example" {
  value = catalystcenter_wireless_controllers_wireless_mobility_groups_mobility_provision.example
}