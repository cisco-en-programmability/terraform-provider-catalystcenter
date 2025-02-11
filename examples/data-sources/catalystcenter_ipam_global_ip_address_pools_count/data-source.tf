
data "catalystcenter_ipam_global_ip_address_pools_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_ipam_global_ip_address_pools_count_example" {
    value = data.catalystcenter_ipam_global_ip_address_pools_count.example.item
}
