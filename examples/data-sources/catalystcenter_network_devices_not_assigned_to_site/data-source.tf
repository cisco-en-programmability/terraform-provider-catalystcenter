
data "catalystcenter_network_devices_not_assigned_to_site" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
}

output "catalystcenter_network_devices_not_assigned_to_site_example" {
    value = data.catalystcenter_network_devices_not_assigned_to_site.example.item
}
