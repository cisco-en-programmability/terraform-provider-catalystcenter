
data "catalystcenter_site_wise_product_names_count" "example" {
    provider = catalystcenter
    product_name = "string"
    site_id = "string"
}

output "catalystcenter_site_wise_product_names_count_example" {
    value = data.catalystcenter_site_wise_product_names_count.example.item
}
