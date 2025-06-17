
data "catalystcenter_golden_tag_image" "example" {
  provider                 = catalystcenter
  device_family_identifier = "string"
  device_role              = "string"
  image_id                 = "string"
  site_id                  = "string"
}

output "catalystcenter_golden_tag_image_example" {
  value = data.catalystcenter_golden_tag_image.example.item
}
