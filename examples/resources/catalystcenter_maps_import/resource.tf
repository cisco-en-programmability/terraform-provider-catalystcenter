
resource "catalystcenter_maps_import" "example" {
  provider            = catalystcenter
  import_context_uuid = "string"
  parameters {

  }
}

output "catalystcenter_maps_import_example" {
  value = catalystcenter_maps_import.example
}