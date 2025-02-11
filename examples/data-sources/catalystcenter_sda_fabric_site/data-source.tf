
data "catalystcenter_sda_fabric_site" "example" {
    provider = catalystcenter
    site_name_hierarchy = "string"
}

output "catalystcenter_sda_fabric_site_example" {
    value = data.catalystcenter_sda_fabric_site.example.item
}
