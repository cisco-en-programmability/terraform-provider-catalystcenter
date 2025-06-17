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

resource "catalystcenter_command_runner_run_command" "example" {
  provider = catalystcenter

  parameters {
    commands     = ["string"]
    description  = "string"
    device_uuids = ["string"]
    name         = "string"
    timeout      = 1
  }
}
