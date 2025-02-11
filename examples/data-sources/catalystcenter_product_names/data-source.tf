
data "catalystcenter_product_names" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
    product_id = "string"
    product_name = "string"
}

output "catalystcenter_product_names_example" {
    value = data.catalystcenter_product_names.example.items
}

data "catalystcenter_product_names" "example" {
    provider = catalystcenter
    product_name_ordinal = 1.0
}

output "catalystcenter_product_names_example" {
    value = data.catalystcenter_product_names.example.item
}
