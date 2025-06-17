
resource "catalystcenter_maps_export" "example" {
  provider            = catalystcenter
  site_hierarchy_uuid = "string"
  parameters = [{

  }]
}

output "catalystcenter_maps_export_example" {
  value = catalystcenter_maps_export.example
}