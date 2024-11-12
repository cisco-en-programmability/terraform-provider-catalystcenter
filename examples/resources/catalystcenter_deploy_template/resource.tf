
resource "catalystcenter_deploy_template" "example" {
  provider = meraki
  parameters {

    force_push_template             = "false"
    is_composite                    = "false"
    main_template_id                = "string"
    member_template_deployment_info = ["string"]
    target_info {

      host_name             = "string"
      id                    = "string"
      params                = "------"
      resource_params       = ["string"]
      type                  = "string"
      versioned_template_id = "string"
    }
    template_id = "string"
  }
}

output "catalystcenter_deploy_template_example" {
  value = catalystcenter_deploy_template.example
}