
resource "catalystcenter_pnp_device_config_preview" "example" {
  provider = meraki
  parameters {

    device_id = "string"
    site_id   = "string"
    type      = "string"
  }
}

output "catalystcenter_pnp_device_config_preview_example" {
  value = catalystcenter_pnp_device_config_preview.example
}