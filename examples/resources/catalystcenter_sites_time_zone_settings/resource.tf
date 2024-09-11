
resource "catalystcenter_sites_time_zone_settings" "example" {
  provider = catalystcenter

  parameters {

    id = "string"
    time_zone {

      identifier = "string"
    }
  }
}

output "catalystcenter_sites_time_zone_settings_example" {
  value = catalystcenter_sites_time_zone_settings.example
}