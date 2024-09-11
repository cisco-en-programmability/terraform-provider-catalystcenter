
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

resource "catalystcenter_network_device" "example" {
  provider = catalystcenter
  parameters {

    id = "3eb928b8-2414-4121-ac35-1247e5d666a4"
  }
}

output "catalystcenter_network_device_example" {
  value = catalystcenter_network_device.example
}
/*
resource "catalystcenter_compliance" "example" {
  depends_on = [catalystcenter_network_device.example]
  provider   = catalystcenter
  parameters {
    trigger_full = true
    categories   = ["PSIRT"]
    device_uuids = [catalystcenter_network_device.example.item.0.id]
  }
}

output "catalystcenter_compliance_example" {
  value = catalystcenter_compliance.example
}
*/