
data "catalystcenter_device_credential" "example" {
  provider = catalystcenter
  site_id  = "string"
}

output "catalystcenter_device_credential_example" {
  value = data.catalystcenter_device_credential.example.item
}
