
data "catalystcenter_sites_device_credentials" "example" {
  provider  = catalystcenter
  id        = "string"
  inherited = "false"
}

output "catalystcenter_sites_device_credentials_example" {
  value = data.catalystcenter_sites_device_credentials.example.item
}
