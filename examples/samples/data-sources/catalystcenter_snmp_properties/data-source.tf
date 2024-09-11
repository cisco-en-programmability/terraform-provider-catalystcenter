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

data "catalystcenter_snmp_properties" "example" {
  provider = catalystcenter
}

output "catalystcenter_snmp_properties_example" {
  value = data.catalystcenter_snmp_properties.example.items
}
