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

data "catalystcenter_global_pool" "example" {
  provider = catalystcenter
  limit    = "1"
  offset   = "1"
}

output "catalystcenter_global_pool_example" {
  value = data.catalystcenter_global_pool.example.items
}
