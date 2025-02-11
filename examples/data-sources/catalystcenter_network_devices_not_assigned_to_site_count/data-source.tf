
data "catalystcenter_network_devices_not_assigned_to_site_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_network_devices_not_assigned_to_site_count_example" {
    value = data.catalystcenter_network_devices_not_assigned_to_site_count.example.item
}
