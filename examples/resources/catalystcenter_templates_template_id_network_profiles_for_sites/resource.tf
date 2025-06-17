
resource "catalystcenter_templates_template_id_network_profiles_for_sites" "example" {
  provider = catalystcenter
 
  parameters {

    profile_id  = "string"
    template_id = "string"
  }
}

output "catalystcenter_templates_template_id_network_profiles_for_sites_example" {
  value = catalystcenter_templates_template_id_network_profiles_for_sites.example
}