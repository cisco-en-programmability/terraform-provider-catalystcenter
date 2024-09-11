
data "catalystcenter_sda_anycast_gateways" "example" {
  provider             = catalystcenter
  fabric_id            = "string"
  id                   = "string"
  ip_pool_name         = "string"
  limit                = 1
  offset               = 1
  virtual_network_name = "string"
  vlan_id              = 1.0
  vlan_name            = "string"
}

output "catalystcenter_sda_anycast_gateways_example" {
  value = data.catalystcenter_sda_anycast_gateways.example.items
}
