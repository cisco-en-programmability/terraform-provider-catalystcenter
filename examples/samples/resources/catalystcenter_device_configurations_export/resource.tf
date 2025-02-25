
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

resource "catalystcenter_device_configurations_export" "example" {
  provider = catalystcenter

  parameters {
    device_id = ["3923aed0-16e5-4ed0-b430-ff6dcfd9c517"]
    password  = "Hola123*"
  }
}
