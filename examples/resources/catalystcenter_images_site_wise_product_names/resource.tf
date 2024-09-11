
resource "catalystcenter_images_site_wise_product_names" "example" {
  provider = catalystcenter

  parameters {

    image_id             = "string"
    product_name_ordinal = 1.0
    site_ids             = ["string"]
  }
}

output "catalystcenter_images_site_wise_product_names_example" {
  value = catalystcenter_images_site_wise_product_names.example
}