
terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_floor" "example" {
  provider = catalystcenter
  parameters {
    site {
      floor {
        floor_number = 5
        height       = 5
        length       = 100
        name         = "floor-1"
        parent_name  = "Global"
        rf_model     = "Cubes And Walled Offices"
        width        = 100
      }
    }
    type = "floor"
    # site_id ="70c232d5-141e-4a03-918e-5a81acda6f38"
  }
}

output "catalystcenter_floor_example" {
  value = catalystcenter_floor.example
}
