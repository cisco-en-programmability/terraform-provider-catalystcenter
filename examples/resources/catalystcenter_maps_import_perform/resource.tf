
resource "catalystcenter_maps_import_perform" "example" {
  provider            = catalystcenter
  import_context_uuid = "string"
  parameters = [{

  }]
}

output "catalystcenter_maps_import_perform_example" {
  value = catalystcenter_maps_import_perform.example
}