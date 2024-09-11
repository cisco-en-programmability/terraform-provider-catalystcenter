terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_wireless_provision_access_point" "example" {
  provider = catalystcenter

  parameters {
    payload {
      custom_ap_group_name   = "string"
      custom_flex_group_name = ["string"]
      device_name            = "string"
      rf_profile             = "string"
      site_id                = "string"
      site_name_hierarchy    = "string"
      type                   = "string"
    }
  }
}
