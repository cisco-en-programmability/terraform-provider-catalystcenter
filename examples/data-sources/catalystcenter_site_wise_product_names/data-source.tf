
data "catalystcenter_site_wise_product_names" "example" {
  provider     = catalystcenter
  limit        = 1
  offset       = 1
  product_name = "string"
  site_id      = "string"
}

output "catalystcenter_site_wise_product_names_example" {
  value = data.catalystcenter_site_wise_product_names.example.item
}
