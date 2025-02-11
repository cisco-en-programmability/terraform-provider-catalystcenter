
data "catalystcenter_network_devices_assigned_to_site_id" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_network_devices_assigned_to_site_id_example" {
    value = data.catalystcenter_network_devices_assigned_to_site_id.example.item
}
