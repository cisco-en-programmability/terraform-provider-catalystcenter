
terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

#variable "username" {
#  type = string
#}

#variable "password" {
#  type = string
#}

# Configure provider with your Cisco Catalyst Center SDK credentials


resource "catalystcenter_area" "test_area_1" {
  provider = catalystcenter
  parameters {

    site {

      area {

        name        = "Test Area 1"
        parent_name = "Global"
      }
    }
    type = "area"
  }
}

resource "catalystcenter_building" "test_building_1" {
  provider = catalystcenter
  parameters {

    site {

      building {

        address     = "8 AV JEAN MEDECIN 06000 NICE FRANCE"
        name        = "Test Building 1"
        parent_name = catalystcenter_area.test_area_1.item[0].site_name_hierarchy
      }
    }
    type = "building"
  }
}

resource "catalystcenter_floor" "test_floor_0" {
  provider = catalystcenter
  parameters {

    site {

      floor {

        floor_number = 0
        height       = 10
        length       = 10
        name         = "Test Floor 0"
        parent_name  = "Global/Test Area 1/Test Building 1"
        rf_model     = "Cubes And Walled Offices"
        width        = 10
      }
    }
    type = "floor"
  }
}


resource "catalystcenter_floor" "test_floor_1" {
  provider = catalystcenter
  parameters {

    site {

      floor {

        floor_number = 1
        height       = 10
        length       = 10
        name         = "Test Floor 1"
        parent_name  = catalystcenter_building.test_building_1.item[0].site_name_hierarchy
        rf_model     = "Cubes And Walled Offices"
        width        = 10
      }
    }
    type = "floor"
  }
}

