
data "catalystcenter_templates_template_id_versions_count" "example" {
    provider = catalystcenter
    latest_version = "false"
    template_id = "string"
    version_number = 1
}

output "catalystcenter_templates_template_id_versions_count_example" {
    value = data.catalystcenter_templates_template_id_versions_count.example.item
}
