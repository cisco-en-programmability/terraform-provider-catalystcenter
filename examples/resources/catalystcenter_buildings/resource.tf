
resource "catalystcenter_buildings" "example" {
  provider = catalystcenter

  parameters {

    address   = "string"
    country   = "string"
    id        = "string"
    latitude  = 1.0
    longitude = 1.0
    name      = "string"
    parent_id = "string"
  }
}

output "catalystcenter_buildings_example" {
  value = catalystcenter_buildings.example
}