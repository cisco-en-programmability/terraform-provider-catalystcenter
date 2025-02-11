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


data "catalystcenter_qos_device_interface" "example" {
  provider          = catalystcenter
  network_device_id = "string"
}

output "catalystcenter_qos_device_interface_example" {
  value = data.catalystcenter_qos_device_interface.example.items
}
