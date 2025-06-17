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

data "catalystcenter_sda_virtual_network" "example" {
  provider             = catalystcenter
  site_name_hierarchy  = "Global/San Francisco"
  virtual_network_name = "Test_terraform"
}

output "catalystcenter_sda_virtual_network_example" {
  value = data.catalystcenter_sda_virtual_network.example.item
}
