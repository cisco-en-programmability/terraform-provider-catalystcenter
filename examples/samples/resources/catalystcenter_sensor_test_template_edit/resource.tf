

terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
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
    schedule {

      frequency {

        unit  = "string"
        value = 1
      }
      schedule_range {

        day = "string"
        time_range {

          frequency {

            unit  = "string"
            value = 1
          }
          from = "string"
          to   = "string"
        }
      }
      test_schedule_mode = "string"
    }
    template_name = "string"
  }
}
