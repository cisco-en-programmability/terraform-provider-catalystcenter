terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}
data "catalystcenter_site" "example" {
  provider = catalystcenter
  # limit    = 1
  # name     = "string"
  # offset   = 1
  # site_id  = "string"
  # type     = "floor"
}

output "catalystcenter_site_example" {
  value = data.catalystcenter_site.example.items
}
