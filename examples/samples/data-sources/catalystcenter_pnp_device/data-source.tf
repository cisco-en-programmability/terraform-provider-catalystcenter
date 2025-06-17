terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}

data "catalystcenter_pnp_device" "example" {
  provider           = catalystcenter
  cm_state           = ["string"]
  hostname           = "string"
  id                 = "618c2b7ff6b3b66f71acbaf8"
  last_contact       = "false"
  limit              = 1
  mac_address        = "string"
  name               = ["string"]
  offset             = 1
  onb_state          = ["string"]
  pid                = ["string"]
  project_id         = ["string"]
  project_name       = ["string"]
  serial_number      = ["string"]
  site_name          = "string"
  smart_account_id   = ["string"]
  sort               = ["string"]
  sort_order         = "string"
  source             = ["string"]
  state              = ["string"]
  virtual_account_id = ["string"]
  workflow_id        = ["string"]
  workflow_name      = ["string"]
}

output "catalystcenter_pnp_device_example" {
  value = data.catalystcenter_pnp_device.example.items[0].version
}
/*
data "catalystcenter_pnp_device" "example" {
  provider = catalystcenter
  id       = "618c2b7ff6b3b66f71acbaf8"
}

output "catalystcenter_pnp_device_example" {
  value = data.catalystcenter_pnp_device.example
}
*/
