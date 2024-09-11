
data "catalystcenter_sites_image_distribution_settings" "example" {
  provider  = catalystcenter
  id        = "string"
  inherited = "false"
}

output "catalystcenter_sites_image_distribution_settings_example" {
  value = data.catalystcenter_sites_image_distribution_settings.example.item
}
