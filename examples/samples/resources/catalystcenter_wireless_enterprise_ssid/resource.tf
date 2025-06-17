
terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}


resource "catalystcenter_wireless_enterprise_ssid" "example" {
  provider = catalystcenter
  parameters {
    name                       = "TestPersonal2"
    ssid_name                  = "TestPersonal2"
    security_level             = "WPA2_ENTERPRISE"
    traffic_type               = "voicedata"
    radio_policy               = "Triple band operation (2.6GHz, 5GHz and 6GHz)"
    fast_transition            = "Adaptive"
    mfp_client_protection      = "Optional"
    protected_management_frame = "Optional"
    # multi_psk_settings {
    #   passphraseType = "ASCII"
    # }

  }
}



output "catalystcenter_wireless_enterprise_ssid_example" {
  value = catalystcenter_wireless_enterprise_ssid.example
}
