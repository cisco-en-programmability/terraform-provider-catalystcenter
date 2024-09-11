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

resource "catalystcenter_sda_fabric_authentication_profile" "example" {
  provider = catalystcenter
  parameters {
    payload {
      authenticate_template_name    = "Open Authentication"
      authentication_order          = "dot1x"
      dot1x_to_mab_fallback_timeout = "21"
      number_of_hosts               = "Unlimited"
      site_name_hierarchy           = "Global/San Francisco"
      wake_on_lan                   = "false"
    }
  }
}

output "catalystcenter_sda_fabric_authentication_profile_example" {
  value = catalystcenter_sda_fabric_authentication_profile.example
}
