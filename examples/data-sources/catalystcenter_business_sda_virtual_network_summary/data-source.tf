
data "catalystcenter_business_sda_virtual_network_summary" "example" {
    provider = catalystcenter
    site_name_hierarchy = "string"
}

output "catalystcenter_business_sda_virtual_network_summary_example" {
    value = data.catalystcenter_business_sda_virtual_network_summary.example.item
}
