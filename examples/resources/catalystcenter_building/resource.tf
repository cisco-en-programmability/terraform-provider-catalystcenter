
resource "catalystcenter_building" "example" {
  provider = catalystcenter
  parameters {

    site {

      building {

        address     = "string"
        country     = "string"
        latitude    = 1.0
        longitude   = 1.0
        name        = "string"
        parent_name = "string"
      }
    }
    site_id = "string"
    type    = "string"
  }
}

output "catalystcenter_building_example" {
  value = catalystcenter_building.example
}