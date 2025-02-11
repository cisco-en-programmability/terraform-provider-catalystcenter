
data "catalystcenter_sites_count" "example" {
    provider = catalystcenter
    name = "string"
}

output "catalystcenter_sites_count_example" {
    value = data.catalystcenter_sites_count.example.items
}
