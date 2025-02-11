
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_deploy_template" "example" {
  provider = catalystcenter

  parameters {
    force_push_template = "false"
    //is_composite                    = "true"
    //member_template_deployment_info = []
    target_info {
      //host_name             = "C9K-Branch-SFO"
      id = "3923aed0-16e5-4ed0-b430-ff6dcfd9c517"
      # params {
      #   int  = "g1/01/12"
      #   desc = "auto"
      # }
      //resource_params       = []
      type = "MANAGED_DEVICE_IP"
      //versioned_template_id = "c49c6179-0631-4b2c-8f94-3216929af040"
    }
    template_id = "fcfd4d19-99e2-494e-9c6f-0d85cf3094e5"
  }
}

output "catalystcenter_deploy_template_example" {
  value = catalystcenter_deploy_template.example
}

# data "catalystcenter_task" "example" {
#   depends_on = [catalystcenter_deploy_template.example]
#   provider   = catalystcenter
#   task_id    = catalystcenter_deploy_template.example.item.0.task_id
# }

# output "catalystcenter_task_example" {
#   value = data.catalystcenter_task.example.item
# }
