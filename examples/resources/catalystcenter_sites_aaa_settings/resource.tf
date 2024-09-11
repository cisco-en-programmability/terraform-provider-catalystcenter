
resource "catalystcenter_sites_aaa_settings" "example" {
  provider = catalystcenter

  parameters {

    aaa_client {

      pan                 = "string"
      primary_server_ip   = "string"
      protocol            = "string"
      secondary_server_ip = "string"
      server_type         = "string"
      shared_secret       = "string"
    }
    aaa_network {

      pan                 = "string"
      primary_server_ip   = "string"
      protocol            = "string"
      secondary_server_ip = "string"
      server_type         = "string"
      shared_secret       = "string"
    }
    id = "string"
  }
}

output "catalystcenter_sites_aaa_settings_example" {
  value = catalystcenter_sites_aaa_settings.example
}