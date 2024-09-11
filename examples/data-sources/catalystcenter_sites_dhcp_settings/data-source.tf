
data "catalystcenter_sites_dhcp_settings" "example" {
  provider  = catalystcenter
  id        = "string"
  inherited = "false"
}

output "catalystcenter_sites_dhcp_settings_example" {
  value = data.catalystcenter_sites_dhcp_settings.example.item
}
