
resource "catalystcenter_sites_device_credentials_apply" "example" {
  provider = catalystcenter
  parameters = [{

    device_credential_id = "string"
    site_id              = "string"
  }]
}

output "catalystcenter_sites_device_credentials_apply_example" {
  value = catalystcenter_sites_device_credentials_apply.example
}