
data "catalystcenter_configuration_template_deploy_status" "example" {
    provider = catalystcenter
    deployment_id = "string"
}

output "catalystcenter_configuration_template_deploy_status_example" {
    value = data.catalystcenter_configuration_template_deploy_status.example.item
}
