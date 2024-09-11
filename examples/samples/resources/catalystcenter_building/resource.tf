
terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_building" "example" {
  count    = 100
  provider = catalystcenter
  parameters {
    site {
      building {
        name        = "Terraform-Test-${count.index}"
        address     = "255 China Basin Street, San Francisco, California 94158, United States 1"
        parent_name = "Global"
        latitude    = 37.77178651716143
        longitude   = -122.39062051589885
      }
    }
    type = "building"
    # site_id ="70c232d5-141e-4a03-918e-5a81acda6f38"
  }
}


output "catalystcenter_building_example" {
  value = catalystcenter_building.example
}
