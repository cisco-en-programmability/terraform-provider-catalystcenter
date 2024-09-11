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

data "catalystcenter_buildings_planned_access_points" "example" {
  provider    = catalystcenter
  building_id = "2397da83-4e12-4d04-9bd3-a57b2ad91652"
  //limit = 1
  //offset = 1
  //radios = "false"
}

output "catalystcenter_buildings_planned_access_points_example" {
  value = data.catalystcenter_buildings_planned_access_points.example.items
}
