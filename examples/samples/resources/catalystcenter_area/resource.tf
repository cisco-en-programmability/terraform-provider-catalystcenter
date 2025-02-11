
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_area" "example" {
  provider = catalystcenter
  parameters {
    site {

      area {
        name        = "test_tf_area13"
        parent_name = "Global"
      }

    }
    type = "area"
    # site_id ="70c232d5-141e-4a03-918e-5a81acda6f38"
  }
}

output "catalystcenter_area_example" {
  value = catalystcenter_area.example
}
