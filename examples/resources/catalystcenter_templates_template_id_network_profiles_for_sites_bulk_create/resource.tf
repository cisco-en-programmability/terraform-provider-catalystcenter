
resource "catalystcenter_templates_template_id_network_profiles_for_sites_bulk_create" "example" {
  provider    = catalystcenter
  template_id = "string"
  parameters = [{

    items = []
  }]
}

output "catalystcenter_templates_template_id_network_profiles_for_sites_bulk_create_example" {
  value = catalystcenter_templates_template_id_network_profiles_for_sites_bulk_create.example
}