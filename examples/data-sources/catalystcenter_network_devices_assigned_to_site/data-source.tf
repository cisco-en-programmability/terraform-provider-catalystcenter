
data "catalystcenter_network_devices_assigned_to_site" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
    site_id = "string"
}

output "catalystcenter_network_devices_assigned_to_site_example" {
    value = data.catalystcenter_network_devices_assigned_to_site.example.items
}
