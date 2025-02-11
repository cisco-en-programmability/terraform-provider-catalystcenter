
data "catalystcenter_dhcp_services_count" "example" {
    provider = catalystcenter
    device_id = "string"
    device_name = "string"
    device_site_hierarchy = "string"
    device_site_hierarchy_id = "string"
    device_site_id = "string"
    end_time = 1609459200
    server_ip = "string"
    start_time = 1609459200
    xca_lle_rid = "string"
}

output "catalystcenter_dhcp_services_count_example" {
    value = data.catalystcenter_dhcp_services_count.example.item
}
