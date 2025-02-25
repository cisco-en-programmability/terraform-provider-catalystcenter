
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_pnp_device" "example" {
  provider = catalystcenter
  # lifecycle {
  #   create_before_destroy = true
  # }
  parameters {

    # id = "61f1c861264f342e4fa1a78e"
    device_info {

      serial_number = "FLM2213W05S"
      stack         = "false"
      # state= "Unclaimed"



      sudi_required = "false"
      hostname      = "FLM2213W05W"


    }
  }
}

output "catalystcenter_pnp_device_example" {
  value = catalystcenter_pnp_device.example
}
