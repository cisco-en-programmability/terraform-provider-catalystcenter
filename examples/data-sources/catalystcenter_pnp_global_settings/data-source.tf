
data "catalystcenter_pnp_global_settings" "example" {
  provider = catalystcenter
}

output "catalystcenter_pnp_global_settings_example" {
  value = data.catalystcenter_pnp_global_settings.example.item
}
