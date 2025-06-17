
data "catalystcenter_templates_template_id_network_profiles_for_sites_count" "example" {
  provider    = catalystcenter
  template_id = "string"
}

output "catalystcenter_templates_template_id_network_profiles_for_sites_count_example" {
  value = data.catalystcenter_templates_template_id_network_profiles_for_sites_count.example.item
}
