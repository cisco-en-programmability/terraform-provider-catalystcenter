terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}

data "catalystcenter_sda_virtual_network_ip_pool" "example" {
  provider             = catalystcenter
  ip_pool_name         = "string"
  virtual_network_name = "string"
}

output "catalystcenter_sda_virtual_network_ip_pool_example" {
  value = data.catalystcenter_sda_virtual_network_ip_pool.example.item
}
