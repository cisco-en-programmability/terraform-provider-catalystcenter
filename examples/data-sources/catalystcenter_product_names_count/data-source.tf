
data "catalystcenter_product_names_count" "example" {
  provider     = catalystcenter
  product_id   = "string"
  product_name = "string"
}

output "catalystcenter_product_names_count_example" {
  value = data.catalystcenter_product_names_count.example.item
}
