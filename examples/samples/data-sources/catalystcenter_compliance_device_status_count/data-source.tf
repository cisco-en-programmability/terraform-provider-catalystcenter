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

data "catalystcenter_compliance_device_status_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_compliance_device_status_count_example" {
  value = data.catalystcenter_compliance_device_status_count.example.item
}
