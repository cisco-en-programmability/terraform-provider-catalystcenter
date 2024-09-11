
data "catalystcenter_images_count" "example" {
  provider                        = catalystcenter
  golden                          = "string"
  has_addon_images                = "false"
  imported                        = "false"
  integrity                       = "string"
  is_addon_images                 = "false"
  name                            = "string"
  product_name_ordinal            = 1.0
  site_id                         = "string"
  supervisor_product_name_ordinal = 1.0
  version                         = "string"
}

output "catalystcenter_images_count_example" {
  value = data.catalystcenter_images_count.example.item
}
