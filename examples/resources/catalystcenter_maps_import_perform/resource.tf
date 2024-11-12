
resource "catalystcenter_maps_import_perform" "example" {
  provider            = meraki
  import_context_uuid = "string"
  parameters {

  }
}

output "catalystcenter_maps_import_perform_example" {
  value = catalystcenter_maps_import_perform.example
}