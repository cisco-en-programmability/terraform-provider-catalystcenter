
data "catalystcenter_configuration_template" "example" {
    provider = catalystcenter
    filter_conflicting_templates = "false"
    product_family = "string"
    product_series = "string"
    product_type = "string"
    project_id = "string"
    project_names = ["string"]
    software_type = "string"
    software_version = "string"
    sort_order = "string"
    tags = ["string"]
    un_committed = "false"
}

output "catalystcenter_configuration_template_example" {
    value = data.catalystcenter_configuration_template.example.items
}

data "catalystcenter_configuration_template" "example" {
    provider = catalystcenter
    latest_version = "false"
    template_id = "string"
}

output "catalystcenter_configuration_template_example" {
    value = data.catalystcenter_configuration_template.example.item
}
