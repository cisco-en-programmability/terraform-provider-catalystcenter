
resource "catalystcenter_sites_image_distribution_settings" "example" {
  provider = catalystcenter

  parameters {

    id = "string"
    image_distribution {

      servers = ["string"]
    }
  }
}

output "catalystcenter_sites_image_distribution_settings_example" {
  value = catalystcenter_sites_image_distribution_settings.example
}