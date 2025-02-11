
data "catalystcenter_maps_supported_access_points" "example" {
    provider = catalystcenter
}

output "catalystcenter_maps_supported_access_points_example" {
    value = data.catalystcenter_maps_supported_access_points.example.items
}
