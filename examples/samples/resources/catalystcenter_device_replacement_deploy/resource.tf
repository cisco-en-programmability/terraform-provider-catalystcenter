
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

resource "catalystcenter_device_replacement_deploy" "example" {
  provider = catalystcenter

  parameters {
    faulty_device_serial_number      = "string"
    replacement_device_serial_number = "string"
  }
}