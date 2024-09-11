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

data "catalystcenter_site_count" "example" {
  provider = catalystcenter
  site_id  = "3e0db2cd-cf3a-4dbd-bfb9-739271ffc20b"
}

output "catalystcenter_site_count_example" {
  value = data.catalystcenter_site_count.example.item
}
