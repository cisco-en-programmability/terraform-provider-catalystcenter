
resource "catalystcenter_images_download" "example" {
  provider = meraki
  id       = "string"
  parameters {

  }
}

output "catalystcenter_images_download_example" {
  value = catalystcenter_images_download.example
}