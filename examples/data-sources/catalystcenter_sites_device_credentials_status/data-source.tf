
data "catalystcenter_sites_device_credentials_status" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_sites_device_credentials_status_example" {
  value = data.catalystcenter_sites_device_credentials_status.example.item
}
