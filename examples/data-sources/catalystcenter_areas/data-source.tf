
data "catalystcenter_areas" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_areas_example" {
    value = data.catalystcenter_areas.example.item
}
