
data "catalystcenter_templates_template_id_network_profiles_for_sites" "example" {
    provider = catalystcenter
    template_id = "string"
}

output "catalystcenter_templates_template_id_network_profiles_for_sites_example" {
    value = data.catalystcenter_templates_template_id_network_profiles_for_sites.example.item
}
