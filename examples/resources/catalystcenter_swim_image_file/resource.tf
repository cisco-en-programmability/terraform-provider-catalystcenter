
resource "catalystcenter_swim_image_file" "example" {
  provider                     = meraki
  is_third_party               = "false"
  third_party_application_type = "string"
  third_party_image_family     = "string"
  third_party_vendor           = "string"
}

output "catalystcenter_swim_image_file_example" {
  value = catalystcenter_swim_image_file.example
}