
data "catalystcenter_sda_virtual_network_ip_pool" "example" {
  provider             = catalystcenter
  ip_pool_name         = "string"
  site_name_hierarchy  = "string"
  virtual_network_name = "string"
}

output "catalystcenter_sda_virtual_network_ip_pool_example" {
  value = data.catalystcenter_sda_virtual_network_ip_pool.example.item
}
