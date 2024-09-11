
resource "catalystcenter_floors_upload_image" "example" {
  provider = catalystcenter
  id       = "string"
  parameters {

  }
}

output "catalystcenter_floors_upload_image_example" {
  value = catalystcenter_floors_upload_image.example
}