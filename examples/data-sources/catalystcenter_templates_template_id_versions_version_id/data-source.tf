
data "catalystcenter_templates_template_id_versions_version_id" "example" {
    provider = catalystcenter
    template_id = "string"
    version_id = "string"
}

output "catalystcenter_templates_template_id_versions_version_id_example" {
    value = data.catalystcenter_templates_template_id_versions_version_id.example.item
}
