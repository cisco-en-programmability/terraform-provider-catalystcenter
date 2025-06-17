
resource "catalystcenter_maps_import_start" "example" {
  provider = catalystcenter
}

output "catalystcenter_maps_import_start_example" {
  value = catalystcenter_maps_import_start.example
}