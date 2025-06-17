
data "catalystcenter_ipam_global_ip_address_pools_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_ipam_global_ip_address_pools_id_example" {
  value = data.catalystcenter_ipam_global_ip_address_pools_id.example.item
}
