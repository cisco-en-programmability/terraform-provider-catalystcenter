
data "catalystcenter_deploy_template" "example" {
  provider      = catalystcenter
  deployment_id = "string"
}

output "catalystcenter_deploy_template_example" {
  value = data.catalystcenter_deploy_template.example.item
}
