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

data "catalystcenter_wireless_rf_profile" "example" {
  provider        = catalystcenter
  rf_profile_name = "TYPICAL"
}

output "catalystcenter_wireless_rf_profile_example" {
  value = data.catalystcenter_wireless_rf_profile.example.items
}
