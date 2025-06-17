terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

data "catalystcenter_interface" "example" {
  provider       = catalystcenter
  interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486809"
}

output "catalystcenter_interface_example" {
  value = data.catalystcenter_interface.example.item
}
