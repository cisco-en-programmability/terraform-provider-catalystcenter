
resource "catalystcenter_disassociate_site_to_network_profile" "example" {
  provider           = catalystcenter
  network_profile_id = "string"
  site_id            = "string"
  parameters {

  }
}

output "catalystcenter_disassociate_site_to_network_profile_example" {
  value = catalystcenter_disassociate_site_to_network_profile.example
}