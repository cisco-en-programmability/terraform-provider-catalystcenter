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

data "catalystcenter_site_design_floormap" "example" {
  provider = catalystcenter
  floor_id = "0f6661c2-ba34-4f4d-ae60-459cf293f689"
}

output "catalystcenter_site_design_floormap_example" {
  value = data.catalystcenter_site_design_floormap.example.item
}
