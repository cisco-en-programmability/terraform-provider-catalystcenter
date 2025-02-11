
data "catalystcenter_network_devices_assigned_to_site_count" "example" {
    provider = catalystcenter
    site_id = "string"
}

output "catalystcenter_network_devices_assigned_to_site_count_example" {
    value = data.catalystcenter_network_devices_assigned_to_site_count.example.item
}
