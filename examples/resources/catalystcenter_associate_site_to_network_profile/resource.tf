
resource "catalystcenter_associate_site_to_network_profile" "example" {
  provider           = meraki
  network_profile_id = "string"
  site_id            = "string"
  parameters {

  }
}

output "catalystcenter_associate_site_to_network_profile_example" {
  value = catalystcenter_associate_site_to_network_profile.example
}