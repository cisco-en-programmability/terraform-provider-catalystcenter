
data "catalystcenter_buildings" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_buildings_example" {
    value = data.catalystcenter_buildings.example.item
}
