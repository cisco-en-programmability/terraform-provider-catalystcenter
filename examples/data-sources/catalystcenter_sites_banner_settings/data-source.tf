
data "catalystcenter_sites_banner_settings" "example" {
  provider  = catalystcenter
  id        = "string"
  inherited = "false"
}

output "catalystcenter_sites_banner_settings_example" {
  value = data.catalystcenter_sites_banner_settings.example.item
}
