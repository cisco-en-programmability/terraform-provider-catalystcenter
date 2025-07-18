
resource "catalystcenter_site" "example" {
  provider = catalystcenter
 
  parameters {

    site {

      area {

        name        = "string"
        parent_name = "string"
      }
      building {

        address     = "string"
        country     = "string"
        latitude    = 1.0
        longitude   = 1.0
        name        = "string"
        parent_name = "string"
      }
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

output "catalystcenter_site_example" {
  value = catalystcenter_site.example
}