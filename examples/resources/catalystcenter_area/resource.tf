
resource "catalystcenter_area" "example" {
  provider = catalystcenter
  parameters {

    site {

      area {

        name        = "string"
        parent_name = "string"
      }

    }
    site_id = "string"
    type    = "string"
  }
}

output "catalystcenter_area_example" {
  value = catalystcenter_area.example
}