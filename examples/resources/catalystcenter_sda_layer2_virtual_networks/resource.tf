
resource "catalystcenter_sda_layer2_virtual_networks" "example" {
  provider = catalystcenter
 
  parameters {

    associated_layer3_virtual_network_name = "string"
    fabric_id                              = "string"
    id                                     = "string"
    is_fabric_enabled_wireless             = "false"
    is_multiple_ip_to_mac_addresses        = "false"
    is_resource_guard_enabled              = "false"
    is_wireless_flooding_enabled           = "false"
    layer2_flooding_address                = "string"
    layer2_flooding_address_assignment     = "string"
    traffic_type                           = "string"
    vlan_id                                = 1
    vlan_name                              = "string"
  }
}

output "catalystcenter_sda_layer2_virtual_networks_example" {
  value = catalystcenter_sda_layer2_virtual_networks.example
}