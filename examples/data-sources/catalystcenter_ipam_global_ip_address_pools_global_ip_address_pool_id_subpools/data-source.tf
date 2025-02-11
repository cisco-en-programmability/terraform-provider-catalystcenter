
data "catalystcenter_ipam_global_ip_address_pools_global_ip_address_pool_id_subpools" "example" {
    provider = catalystcenter
    global_ip_address_pool_id = "string"
    limit = 1
    offset = 1
}

output "catalystcenter_ipam_global_ip_address_pools_global_ip_address_pool_id_subpools_example" {
    value = data.catalystcenter_ipam_global_ip_address_pools_global_ip_address_pool_id_subpools.example.items
}
