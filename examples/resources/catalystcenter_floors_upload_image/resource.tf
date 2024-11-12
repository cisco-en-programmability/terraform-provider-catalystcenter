
resource "catalystcenter_floors_upload_image" "example" {
  provider = meraki
  id       = "string"
  parameters {

  }
}

output "catalystcenter_floors_upload_image_example" {
  value = catalystcenter_floors_upload_image.example
}