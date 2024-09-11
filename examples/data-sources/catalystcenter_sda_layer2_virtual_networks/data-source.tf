
data "catalystcenter_sda_layer2_virtual_networks" "example" {
  provider                               = catalystcenter
  associated_layer3_virtual_network_name = "string"
  fabric_id                              = "string"
  id                                     = "string"
  limit                                  = 1
  offset                                 = 1
  traffic_type                           = "string"
  vlan_id                                = 1.0
  vlan_name                              = "string"
}

output "catalystcenter_sda_layer2_virtual_networks_example" {
  value = data.catalystcenter_sda_layer2_virtual_networks.example.items
}
