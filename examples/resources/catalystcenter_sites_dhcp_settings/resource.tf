
resource "catalystcenter_sites_dhcp_settings" "example" {
  provider = catalystcenter
 
  parameters {

    dhcp {

      servers = ["string"]
    }
    id = "string"
  }
}

output "catalystcenter_sites_dhcp_settings_example" {
  value = catalystcenter_sites_dhcp_settings.example
}