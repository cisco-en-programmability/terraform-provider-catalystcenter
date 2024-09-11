
terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_deploy_template_v1" "example" {
  provider = catalystcenter

  parameters {
    force_push_template = "false"
    //is_composite                    = "true"
    //member_template_deployment_info = []
    target_info {
      //host_name             = "C9K-Branch-SFO"
      id = "3923aed0-16e5-4ed0-b430-ff6dcfd9c517"
      params {
        int  = "g1/01/12"
        desc = "auto"
      }
      //resource_params       = []
      type = "MANAGED_DEVICE_IP"
      //versioned_template_id = "c49c6179-0631-4b2c-8f94-3216929af040"
    }
    template_id = "fe2bd8b9-2cf0-4b73-b7dc-755ff0f26363"
  }
}

output "catalystcenter_deploy_template_v1_example" {
  value = catalystcenter_deploy_template_v1.example
}

