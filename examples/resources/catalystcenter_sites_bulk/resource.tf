
resource "catalystcenter_sites_bulk" "example" {
    provider = meraki
    parameters = [{
      
      address = "string"
      country = "string"
      floor_number = 1
      height = 1.0
      latitude = 1.0
      length = 1.0
      longitude = 1.0
      name = "string"
      parent_name_hierarchy = "string"
      rf_model = "string"
      type = "string"
      units_of_measure = "string"
      width = 1.0
    }]
}

output "catalystcenter_sites_bulk_example" {
    value = catalystcenter_sites_bulk.example
}