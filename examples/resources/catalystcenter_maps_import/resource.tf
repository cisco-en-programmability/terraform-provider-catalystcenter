
resource "catalystcenter_maps_import" "example" {
  provider            = meraki
  import_context_uuid = "string"
  parameters {

  }
}

output "catalystcenter_maps_import_example" {
  value = catalystcenter_maps_import.example
}