
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_lan_automation_create" "example" {
  provider = catalystcenter
  parameters {

    payload {
      discovered_device_site_name_hierarchy = "string"
      host_name_file_id                     = "string"
      host_name_prefix                      = "string"
      ip_pools {

        ip_pool_name = "string"
        ip_pool_role = "string"
      }
      isis_domain_pwd                    = "string"
      peer_device_managment_ipaddress    = "string"
      primary_device_interface_names     = ["string"]
      primary_device_managment_ipaddress = "string"
      redistribute_isis_to_bgp           = "false"
    }
  }
}

output "catalystcenter_lan_automation_create_example" {
  value = catalystcenter_lan_automation_create.example
}
