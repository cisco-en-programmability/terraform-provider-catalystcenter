terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}
resource "catalystcenter_sda_fabric_sites" "example" {
  provider = catalystcenter

  parameters {
    payload {
      authentication_profile_name = "string"
      id                          = "string"
      is_pub_sub_enabled          = "false"
      site_id                     = "string"
    }
  }
}

output "catalystcenter_sda_fabric_sites_example" {
  value = catalystcenter_sda_fabric_sites.example
}
