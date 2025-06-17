
data "catalystcenter_sda_multicast_virtual_networks" "example" {
  provider             = catalystcenter
  fabric_id            = "string"
  limit                = 1
  offset               = 1
  virtual_network_name = "string"
}

output "catalystcenter_sda_multicast_virtual_networks_example" {
  value = data.catalystcenter_sda_multicast_virtual_networks.example.items
}
