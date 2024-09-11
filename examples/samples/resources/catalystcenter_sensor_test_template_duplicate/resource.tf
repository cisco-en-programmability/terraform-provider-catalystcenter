terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}


resource "catalystcenter_sensor_test_template_duplicate" "example" {
  provider = catalystcenter

  parameters {
    new_template_name = "string"
    template_name     = "string"
  }
}
