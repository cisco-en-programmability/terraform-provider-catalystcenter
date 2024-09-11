
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

resource "catalystcenter_network_device_export" "example" {
  provider = catalystcenter

  parameters {
    device_uuids   = ["string"]
    id             = "string"
    operation_enum = "string"
    parameters     = ["string"]
    password       = "******"
  }
}
