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

resource "catalystcenter_sda_fabric_site" "example" {
  provider = catalystcenter
  parameters {
    fabric_name         = "Default LAN Fabric 2"
    site_name_hierarchy = "Global/San Francisco"
  }
}

output "catalystcenter_sda_fabric_site_example" {
  value = catalystcenter_sda_fabric_site.example
}
