
data "catalystcenter_network_devices_network_profiles_for_sites" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
    order = "string"
    sort_by = "string"
    type = "string"
}

output "catalystcenter_network_devices_network_profiles_for_sites_example" {
    value = data.catalystcenter_network_devices_network_profiles_for_sites.example.items
}

data "catalystcenter_network_devices_network_profiles_for_sites" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_network_devices_network_profiles_for_sites_example" {
    value = data.catalystcenter_network_devices_network_profiles_for_sites.example.item
}
