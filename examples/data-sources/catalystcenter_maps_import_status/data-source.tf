
data "catalystcenter_maps_import_status" "example" {
  provider            = catalystcenter
  import_context_uuid = "string"
}

output "catalystcenter_maps_import_status_example" {
  value = data.catalystcenter_maps_import_status.example.item
}
