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

data "catalystcenter_network_device_config_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_device_config_count_example" {
  value = data.catalystcenter_network_device_config_count.example.item
}
