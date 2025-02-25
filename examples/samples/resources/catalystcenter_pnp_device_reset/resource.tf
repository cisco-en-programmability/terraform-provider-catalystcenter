
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_pnp_device_reset" "example" {
  provider = catalystcenter

  parameters {
    device_reset_list {

      config_list {

        config_id = "string"
        config_parameters {

          key   = "string"
          value = "string"
        }
      }
      device_id                  = "string"
      license_level              = "string"
      license_type               = "string"
      top_of_stack_serial_number = "string"
      project_id                 = "string"
      workflow_id                = "string"
    }

  }
}
