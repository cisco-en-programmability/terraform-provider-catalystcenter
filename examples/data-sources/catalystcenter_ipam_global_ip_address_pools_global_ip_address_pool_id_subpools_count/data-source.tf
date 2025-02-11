
data "catalystcenter_ipam_global_ip_address_pools_global_ip_address_pool_id_subpools_count" "example" {
    provider = catalystcenter
    global_ip_address_pool_id = "string"
}

output "catalystcenter_ipam_global_ip_address_pools_global_ip_address_pool_id_subpools_count_example" {
    value = data.catalystcenter_ipam_global_ip_address_pools_global_ip_address_pool_id_subpools_count.example.item
}
