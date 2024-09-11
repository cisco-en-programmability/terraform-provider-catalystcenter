
data "catalystcenter_sda_anycast_gateways_count" "example" {
  provider             = catalystcenter
  fabric_id            = "string"
  ip_pool_name         = "string"
  virtual_network_name = "string"
  vlan_id              = 1.0
  vlan_name            = "string"
}

output "catalystcenter_sda_anycast_gateways_count_example" {
  value = data.catalystcenter_sda_anycast_gateways_count.example.item
}
