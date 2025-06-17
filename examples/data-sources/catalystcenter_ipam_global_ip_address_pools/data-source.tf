
data "catalystcenter_ipam_global_ip_address_pools" "example" {
  provider = catalystcenter
  limit    = 1
  offset   = 1
  order    = "string"
  sort_by  = "string"
}

output "catalystcenter_ipam_global_ip_address_pools_example" {
  value = data.catalystcenter_ipam_global_ip_address_pools.example.items
}
