
data "catalystcenter_license_term_details" "example" {
  provider             = catalystcenter
  device_type          = "string"
  smart_account_id     = "string"
  virtual_account_name = "string"
}

output "catalystcenter_license_term_details_example" {
  value = data.catalystcenter_license_term_details.example.items
}
