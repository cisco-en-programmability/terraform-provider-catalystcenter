
data "catalystcenter_sda_virtual_network" "example" {
  provider             = catalystcenter
  site_name_hierarchy  = "string"
  virtual_network_name = "string"
}

output "catalystcenter_sda_virtual_network_example" {
  value = data.catalystcenter_sda_virtual_network.example.item
}
