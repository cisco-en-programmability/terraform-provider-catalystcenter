
data "catalystcenter_network_devices_network_profiles_for_sites_count" "example" {
  provider = catalystcenter
  type     = "string"
}

output "catalystcenter_network_devices_network_profiles_for_sites_count_example" {
  value = data.catalystcenter_network_devices_network_profiles_for_sites_count.example.item
}
