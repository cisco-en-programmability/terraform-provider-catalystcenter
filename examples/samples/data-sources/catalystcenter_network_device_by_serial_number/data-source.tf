
terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

data "catalystcenter_network_device_by_serial_number" "example" {
  provider      = catalystcenter
  serial_number = "FOC2214Z084"
}

output "catalystcenter_network_device_by_serial_number_example" {
  value = data.catalystcenter_network_device_by_serial_number.example.item
}
