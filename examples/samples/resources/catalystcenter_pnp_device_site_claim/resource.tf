
terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_pnp_device_site_claim" "this" {
  parameters {
    site_id   = "f1188cff-6110-45a3-946c-4831d62a3cd4"
    device_id = "645b9623e0a25d5fbb063d9b"

    type = "Default" # Default means switch. Available values: Default, StackSwitch, AccessPoint, Sensor, CatalystWLC, MobilityExpress
    image_info {
      skip = true # TODO: don't skip this
    }
    config_info {
      config_id = "d6456fd2-a0b8-4c7c-9fea-f00b12f69b64"
      config_parameters {
        key   = "HOSTNAME"
        value = "labhf-osl-bn-2"
      }
      config_parameters {
        key   = "LOOPBACK0"
        value = "100.68.12.1"
      }
      config_parameters {
        key   = "BORDER-AS"
        value = "65002"
      }
    }
  }
}

# output "catalystcenter_pnp_device_site_claim_example" {
#   value = catalystcenter_pnp_device_site_claim.example
# }
