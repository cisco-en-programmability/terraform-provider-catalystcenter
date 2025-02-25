terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}


data "catalystcenter_network_device_interface_neighbor" "example" {
  provider       = catalystcenter
  device_uuid    = "3923aed0-16e5-4ed0-b430-ff6dcfd9c517"
  interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486809"
}

output "catalystcenter_network_device_interface_neighbor_example" {
  value = data.catalystcenter_network_device_interface_neighbor.example.item
}
