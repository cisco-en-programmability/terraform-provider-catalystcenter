
data "catalystcenter_images_addon_images_count" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_images_addon_images_count_example" {
  value = data.catalystcenter_images_addon_images_count.example.item
}
