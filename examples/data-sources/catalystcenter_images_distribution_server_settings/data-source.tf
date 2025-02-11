
data "catalystcenter_images_distribution_server_settings" "example" {
    provider = catalystcenter
}

output "catalystcenter_images_distribution_server_settings_example" {
    value = data.catalystcenter_images_distribution_server_settings.example.items
}
