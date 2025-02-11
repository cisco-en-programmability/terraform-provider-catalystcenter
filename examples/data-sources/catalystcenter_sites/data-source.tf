
data "catalystcenter_sites" "example" {
    provider = catalystcenter
    limit = 1
    name = "string"
    name_hierarchy = "string"
    offset = 1
    type = "string"
    units_of_measure = "string"
}

output "catalystcenter_sites_example" {
    value = data.catalystcenter_sites.example.items
}
