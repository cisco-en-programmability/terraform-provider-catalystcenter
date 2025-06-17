
resource "catalystcenter_golden_tag_image" "example" {
  provider = catalystcenter
 
  parameters {

    device_family_identifier = "string"
    device_role              = "string"
    image_id                 = "string"
    site_id                  = "string"
  }
}

output "catalystcenter_golden_tag_image_example" {
  value = catalystcenter_golden_tag_image.example
}