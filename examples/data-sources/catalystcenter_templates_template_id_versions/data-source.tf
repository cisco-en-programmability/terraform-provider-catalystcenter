
data "catalystcenter_templates_template_id_versions" "example" {
    provider = catalystcenter
    latest_version = "false"
    limit = 1
    offset = 1
    order = "string"
    template_id = "string"
    version_number = 1
}

output "catalystcenter_templates_template_id_versions_example" {
    value = data.catalystcenter_templates_template_id_versions.example.items
}
