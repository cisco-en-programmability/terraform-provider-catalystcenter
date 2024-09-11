
data "catalystcenter_license_smart_account_details" "example" {
  provider = catalystcenter
}

output "catalystcenter_license_smart_account_details_example" {
  value = data.catalystcenter_license_smart_account_details.example.items
}
