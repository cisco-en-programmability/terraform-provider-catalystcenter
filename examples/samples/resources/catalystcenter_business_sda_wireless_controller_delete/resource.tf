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

resource "catalystcenter_business_sda_wireless_controller_delete" "example" {
  provider = catalystcenter

  parameters {
    device_ipaddress  = "string"
    persistbapioutput = "true"
  }
}
