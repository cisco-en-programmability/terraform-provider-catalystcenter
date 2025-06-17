
data "catalystcenter_sda_layer2_virtual_networks_count" "example" {
  provider                               = catalystcenter
  associated_layer3_virtual_network_name = "string"
  fabric_id                              = "string"
  traffic_type                           = "string"
  vlan_id                                = 1.0
  vlan_name                              = "string"
}

output "catalystcenter_sda_layer2_virtual_networks_count_example" {
  value = data.catalystcenter_sda_layer2_virtual_networks_count.example.item
}
