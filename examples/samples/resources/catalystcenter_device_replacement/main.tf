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
  debug = "true"
}

resource "catalystcenter_device_replacement" "example" {
  provider = catalystcenter
  parameters {
    payload {
      creation_time                    = 1
      family                           = "string"
      faulty_device_id                 = "string"
      faulty_device_name               = "string"
      faulty_device_platform           = "string"
      faulty_device_serial_number      = "string"
      id                               = "string"
      neighbour_device_id              = "string"
      network_readiness_task_id        = "string"
      replacement_device_platform      = "string"
      replacement_device_serial_number = "string"
      replacement_status               = "string"
      replacement_time                 = 1
      workflow_id                      = "string"
    }
  }
}

output "catalystcenter_device_replacement_example" {
  value = catalystcenter_device_replacement.example
}
