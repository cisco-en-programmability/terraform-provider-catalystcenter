
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

resource "catalystcenter_wireless_provision_ssid_create_provision" "example" {
  provider = catalystcenter

  parameters {
    enable_fabric = "false"
    flex_connect {

      enable_flex_connect = "false"
      local_to_vlan       = 1
    }
    managed_aplocations = ["string"]
    ssid_details {

      enable_broadcast_ssi_d = "false"
      enable_fast_lane       = "false"
      enable_mac_filtering   = "false"
      fast_transition        = "string"
      name                   = "string"
      passphrase             = "string"
      radio_policy           = "string"
      security_level         = "string"
      traffic_type           = "string"
      web_auth_url           = "string"
    }
    ssid_type = "string"
  }
}
