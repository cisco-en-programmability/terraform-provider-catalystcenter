
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

resource "catalystcenter_pnp_device_import" "example" {
  provider = catalystcenter

  parameters {
    payload {
      device_info {
        serial_number = "FOCTEST5AB"
      }
    }
    #   payload {
    #     device_info {
    #       serial_number = "FOCTEST511"
    #     }
    #   }
  }
}

output "catalystcenter_pnp_device_import_example" {
  value = catalystcenter_pnp_device_import.example
}
