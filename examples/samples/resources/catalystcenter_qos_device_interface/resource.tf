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

resource "catalystcenter_qos_device_interface" "example" {
  provider = catalystcenter
  parameters {
    payload {
    
    #id                  = "string"
    name              = "test"
    network_device_id = "3eb928b8-2414-4121-ac35-1247e5d666a4"
    qos_device_interface_info {

      #dmvpn_remote_sites_bw = [1]
      #instance_id           = 1
      #interface_id          = "string"
      interface_name = "a"
      #label                 = "string"
      role = "DMVPN_SPOKE"
      #upload_bw             = 1
    }
    }
  }
}

output "catalystcenter_qos_device_interface_example" {
  value = catalystcenter_qos_device_interface.example
}
