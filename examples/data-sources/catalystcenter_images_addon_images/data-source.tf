
data "catalystcenter_images_addon_images" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_images_addon_images_example" {
  value = data.catalystcenter_images_addon_images.example.items
}
