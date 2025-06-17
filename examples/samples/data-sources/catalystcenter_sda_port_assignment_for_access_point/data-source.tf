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

data "catalystcenter_sda_port_assignment_for_access_point" "example" {
  provider                     = catalystcenter
  device_management_ip_address = "10.121.1.1"
  interface_name               = "string"
}

output "catalystcenter_sda_port_assignment_for_access_point_example" {
  value = data.catalystcenter_sda_port_assignment_for_access_point.example.item
}
