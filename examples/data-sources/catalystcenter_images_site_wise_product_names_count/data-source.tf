
data "catalystcenter_images_site_wise_product_names_count" "example" {
  provider     = catalystcenter
  assigned     = "string"
  image_id     = "string"
  product_id   = "string"
  product_name = "string"
  recommended  = "string"
}

output "catalystcenter_images_site_wise_product_names_count_example" {
  value = data.catalystcenter_images_site_wise_product_names_count.example.item
}
