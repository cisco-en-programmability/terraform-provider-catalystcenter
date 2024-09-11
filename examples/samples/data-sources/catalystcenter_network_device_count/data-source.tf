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

data "catalystcenter_network_device_count" "example" {
  provider = catalystcenter
  #device_id = "3923aed0-16e5-4ed0-b430-ff6dcfd9c517"
}

output "catalystcenter_network_device_count_example" {
  value = data.catalystcenter_network_device_count.example.item
}
