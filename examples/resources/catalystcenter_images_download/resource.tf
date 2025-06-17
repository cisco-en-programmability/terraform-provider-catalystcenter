
resource "catalystcenter_images_download" "example" {
  provider = catalystcenter
  id       = "string"
  parameters = [{

  }]
}

output "catalystcenter_images_download_example" {
  value = catalystcenter_images_download.example
}