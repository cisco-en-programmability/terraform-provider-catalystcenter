terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}

data "catalystcenter_wireless_dynamic_interface" "example" {
  provider       = catalystcenter
  interface_name = "management"
}

output "catalystcenter_wireless_dynamic_interface_example" {
  value = data.catalystcenter_wireless_dynamic_interface.example.items
}