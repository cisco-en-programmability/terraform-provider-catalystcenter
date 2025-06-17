
resource "catalystcenter_sda_fabric_sites" "example" {
  provider = catalystcenter
 
  parameters {

    authentication_profile_name = "string"
    id                          = "string"
    is_pub_sub_enabled          = "false"
    site_id                     = "string"
  }
}

output "catalystcenter_sda_fabric_sites_example" {
  value = catalystcenter_sda_fabric_sites.example
}