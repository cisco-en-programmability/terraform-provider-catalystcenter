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

data "catalystcenter_nfv_profile" "example" {
  provider = catalystcenter
  id       = "string"
  limit    = 1
  name     = "string"
  offset   = 1
}

output "catalystcenter_nfv_profile_example" {
  value = data.catalystcenter_nfv_profile.example.items
}
