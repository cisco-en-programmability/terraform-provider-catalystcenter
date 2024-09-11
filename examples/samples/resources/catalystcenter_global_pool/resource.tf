
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.32-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_global_pool" "example" {
  provider = catalystcenter
  parameters {

    #id = "string"
    settings {
      ippool {
        #ip_address_space = "string"
        dhcp_server_ips = ["100.100.100.100"]
        dns_server_ips  = ["101.101.101.101"]
        gateway         = "13s.0.0.1"
        #id               = "string"
        ip_pool_cidr = "14.0.0.0/8"
        ip_pool_name = "13Network"
        type         = "generic"
      }
    }
  }
}

output "catalystcenter_global_pool_example" {
  value = catalystcenter_global_pool.example
}
