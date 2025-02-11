
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

resource "catalystcenter_wireless_provision_ssid_delete_reprovision" "example" {
  provider = catalystcenter

  parameters {
    managed_aplocations = "string"
    ssid_name           = "string"
  }
}
