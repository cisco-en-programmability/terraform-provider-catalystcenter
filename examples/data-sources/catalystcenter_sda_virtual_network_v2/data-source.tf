
data "catalystcenter_sda_virtual_network_v2" "example" {
  provider             = catalystcenter
  virtual_network_name = "string"
}

output "catalystcenter_sda_virtual_network_v2_example" {
  value = data.catalystcenter_sda_virtual_network_v2.example.item
}
