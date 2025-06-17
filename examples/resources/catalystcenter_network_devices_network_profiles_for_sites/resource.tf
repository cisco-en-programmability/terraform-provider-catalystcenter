
resource "catalystcenter_network_devices_network_profiles_for_sites" "example" {
  provider = catalystcenter
 
  parameters {

    id = "string"
  }
}

output "catalystcenter_network_devices_network_profiles_for_sites_example" {
  value = catalystcenter_network_devices_network_profiles_for_sites.example
}