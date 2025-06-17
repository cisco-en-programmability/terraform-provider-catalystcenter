
resource "catalystcenter_license_renew" "example" {
  provider = catalystcenter
}

output "catalystcenter_license_renew_example" {
  value = catalystcenter_license_renew.example
}