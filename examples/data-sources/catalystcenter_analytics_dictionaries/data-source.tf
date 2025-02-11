
data "catalystcenter_analytics_dictionaries" "example" {
    provider = catalystcenter
    include_attributes = "false"
}

output "catalystcenter_analytics_dictionaries_example" {
    value = data.catalystcenter_analytics_dictionaries.example.items
}
