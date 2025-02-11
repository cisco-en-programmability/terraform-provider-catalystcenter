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

data "catalystcenter_sda_virtual_network_v2" "example" {
  provider             = catalystcenter
  virtual_network_name = "GUEST_VN"
}

output "catalystcenter_sda_virtual_network_v2_example" {
  value = data.catalystcenter_sda_virtual_network_v2.example.item
}
