

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

resource "catalystcenter_sensor_test_template_edit" "example" {
  provider = catalystcenter

  parameters {
    location_info_list {

      all_sensors    = "false"
      location_id    = "string"
      location_type  = "string"
      site_hierarchy = "string"
    }
    template_name = "string"
  }
}
