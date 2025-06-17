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
}

data "catalystcenter_sda_fabric_authentication_profile" "example" {
  provider = catalystcenter
  # authenticate_template_name = "Test"
  site_name_hierarchy = "Global/Pennsylvania"
}

output "catalystcenter_sda_fabric_authentication_profile_example" {
  value = data.catalystcenter_sda_fabric_authentication_profile.example.item
}
