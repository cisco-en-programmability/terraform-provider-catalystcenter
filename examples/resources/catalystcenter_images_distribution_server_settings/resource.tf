
resource "catalystcenter_images_distribution_server_settings" "example" {
    provider = catalystcenter

    parameters {

      password = "******"
      port_number = 1.0
      root_location = "string"
      server_address = "string"
      username = "string"
    }
}

output "catalystcenter_images_distribution_server_settings_example" {
    value = catalystcenter_images_distribution_server_settings.example
}
