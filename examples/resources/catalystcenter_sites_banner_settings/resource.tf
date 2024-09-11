
resource "catalystcenter_sites_banner_settings" "example" {
  provider = catalystcenter

  parameters {

    banner {

      message = "string"
      type    = "string"
    }
    id = "string"
  }
}

output "catalystcenter_sites_banner_settings_example" {
  value = catalystcenter_sites_banner_settings.example
}