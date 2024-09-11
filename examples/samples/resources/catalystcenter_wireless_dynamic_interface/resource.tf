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
  debug = "true"
}

resource "catalystcenter_wireless_dynamic_interface" "example" {
  provider = catalystcenter
  parameters {

    interface_name = "hello"
    //vlan_id        = 1.0
  }
}

output "catalystcenter_wireless_dynamic_interface_example" {
  value = catalystcenter_wireless_dynamic_interface.example
}
