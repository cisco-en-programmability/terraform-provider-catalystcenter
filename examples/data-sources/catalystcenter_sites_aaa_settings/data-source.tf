
data "catalystcenter_sites_aaa_settings" "example" {
  provider  = catalystcenter
  id        = "string"
  inherited = "false"
}

output "catalystcenter_sites_aaa_settings_example" {
  value = data.catalystcenter_sites_aaa_settings.example.item
}
