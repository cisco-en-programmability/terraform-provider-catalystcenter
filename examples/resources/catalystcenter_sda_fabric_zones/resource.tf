
resource "catalystcenter_sda_fabric_zones" "example" {
  provider = catalystcenter
 
  parameters {

    authentication_profile_name = "string"
    id                          = "string"
    site_id                     = "string"
  }
}

output "catalystcenter_sda_fabric_zones_example" {
  value = catalystcenter_sda_fabric_zones.example
}