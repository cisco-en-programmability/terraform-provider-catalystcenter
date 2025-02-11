terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}

data "catalystcenter_wireless_psk_override" "example" {
  provider = catalystcenter
  payload {
    pass_phrase = "create"
    site        = "Global/San Francisco"
    ssid        = "test999_pop"
  }
}
