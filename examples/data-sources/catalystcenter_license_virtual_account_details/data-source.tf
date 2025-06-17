
data "catalystcenter_license_virtual_account_details" "example" {
  provider         = catalystcenter
  smart_account_id = "string"
}

output "catalystcenter_license_virtual_account_details_example" {
  value = data.catalystcenter_license_virtual_account_details.example.item
}
