
resource "catalystcenter_images_id_sites_site_id_untag_golden" "example" {
    provider = meraki
    id = "string"
    site_id = "string"
    parameters = [{
      
      device_roles = ["string"]
      device_tags = ["string"]
      product_name_ordinal = 1.0
      supervisor_product_name_ordinal = 1.0
    }]
}

output "catalystcenter_images_id_sites_site_id_untag_golden_example" {
    value = catalystcenter_images_id_sites_site_id_untag_golden.example
}