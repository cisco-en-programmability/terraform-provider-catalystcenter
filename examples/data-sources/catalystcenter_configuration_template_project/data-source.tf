
data "catalystcenter_configuration_template_project" "example" {
    provider = catalystcenter
    name = "string"
    sort_order = "string"
}

output "catalystcenter_configuration_template_project_example" {
    value = data.catalystcenter_configuration_template_project.example.items
}

data "catalystcenter_configuration_template_project" "example" {
    provider = catalystcenter
    project_id = "string"
}

output "catalystcenter_configuration_template_project_example" {
    value = data.catalystcenter_configuration_template_project.example.item
}
