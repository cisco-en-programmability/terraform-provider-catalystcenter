
resource "catalystcenter_floor" "example" {
  provider = catalystcenter
  parameters {

    site {

      floor {

        floor_number = 1.0
        height       = 1.0
        length       = 1.0
        name         = "string"
        parent_name  = "string"
        rf_model     = "string"
        width        = 1.0
      }
    }
    site_id = "string"
    type    = "string"
  }
}

output "catalystcenter_floor_example" {
  value = catalystcenter_floor.example
}