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

data "catalystcenter_application_sets" "example" {
  provider = catalystcenter
  limit    = 1
  name     = "consumer-misc"
  offset   = 1
}

output "catalystcenter_application_sets_example" {
  value = data.catalystcenter_application_sets.example.items
}
