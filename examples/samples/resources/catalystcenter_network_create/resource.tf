terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_network_create" "global" {
  provider = catalystcenter

  parameters {

    settings {
      dhcp_server = ["1.1.1.1"]

      dns_server {
        domain_name        = "hola"
        primary_ip_address = "1.1.1.1"
      }
      message_of_theday {
        banner_message = "Have a good day"
      }
    }

    site_id = "771662ca-cb7e-47ff-9eaf-9b8c85e8e389"
  }
}
output "catalystcenter_network_create_example" {
  value = catalystcenter_network_create.global
}
