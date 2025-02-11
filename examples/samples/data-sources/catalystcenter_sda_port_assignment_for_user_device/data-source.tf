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

data "catalystcenter_sda_port_assignment_for_user_device" "example" {
  provider                     = catalystcenter
  device_management_ip_address = "string"
  interface_name               = "string"
}

output "catalystcenter_sda_port_assignment_for_user_device_example" {
  value = data.catalystcenter_sda_port_assignment_for_user_device.example.item
}
