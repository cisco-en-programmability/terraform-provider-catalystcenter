
data "catalystcenter_ipam_site_ip_address_pools_count" "example" {
    provider = catalystcenter
    site_id = "string"
}

output "catalystcenter_ipam_site_ip_address_pools_count_example" {
    value = data.catalystcenter_ipam_site_ip_address_pools_count.example.item
}
