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

data "catalystcenter_reserve_ip_subpool" "example" {
  provider = catalystcenter
  limit    = "2"
  offset   = "1"
  site_id  = "9e860d9e-6499-40d1-9645-4b45bd684219"
}

output "catalystcenter_reserve_ip_subpool_example" {
  value = data.catalystcenter_reserve_ip_subpool.example.items
}
