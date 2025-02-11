
resource "catalystcenter_templates_template_id_network_profiles_for_sites_bulk_delete" "example" {
    provider = meraki
    profile_id = "string"
    template_id = "string"
    parameters = [{
      
    }]
}

output "catalystcenter_templates_template_id_network_profiles_for_sites_bulk_delete_example" {
    value = catalystcenter_templates_template_id_network_profiles_for_sites_bulk_delete.example
}