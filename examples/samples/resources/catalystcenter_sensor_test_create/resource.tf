
terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}


resource "catalystcenter_sensor_test_create" "example" {
  provider = catalystcenter

  parameters {
    ap_coverage {

      bands                 = "string"
      number_of_aps_to_test = "string"
      rssi_threshold        = "string"
    }
    model_version = 1
    name          = "string"
    ssids {

      auth_type    = "string"
      profile_name = "string"
      psk          = "string"
      qos_policy   = "string"
      ssid         = "string"
      tests {

        name   = "string"
      }
      third_party {

        selected = "false"
      }
    }
  }
}
