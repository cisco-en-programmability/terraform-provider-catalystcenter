
data "catalystcenter_images" "example" {
    provider = catalystcenter
    golden = "false"
    has_addon_images = "false"
    imported = "false"
    integrity = "string"
    is_addon_images = "false"
    limit = 1
    name = "string"
    offset = 1
    product_name_ordinal = 1.0
    site_id = "string"
    supervisor_product_name_ordinal = 1.0
    version = "string"
}

output "catalystcenter_images_example" {
    value = data.catalystcenter_images.example.items
}
