
resource "catalystcenter_sites_ntp_settings" "example" {
  provider = catalystcenter
 
  parameters {

    id = "string"
    ntp {

      servers = ["string"]
    }
  }
}

output "catalystcenter_sites_ntp_settings_example" {
  value = catalystcenter_sites_ntp_settings.example
}