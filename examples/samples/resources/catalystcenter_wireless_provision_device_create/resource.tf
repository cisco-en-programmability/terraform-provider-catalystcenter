terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_wireless_provision_device_create" "example" {
  provider = catalystcenter

  parameters {
    payload {
      device_name = "string"
      dynamic_interfaces {

        interface_gateway          = "string"
        interface_ipaddress        = "string"
        interface_name             = "string"
        interface_netmask_in_cid_r = 1
        lag_or_port_number         = 1
        vlan_id                    = 1
      }
      managed_aplocations = ["string"]
      site                = "string"
    }
  }
}
