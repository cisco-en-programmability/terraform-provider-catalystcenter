
data "catalystcenter_license_status" "example" {
  provider = catalystcenter
}

output "catalystcenter_license_status_example" {
  value = data.catalystcenter_license_status.example.item
}
