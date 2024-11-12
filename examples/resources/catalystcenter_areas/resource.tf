
resource "catalystcenter_areas" "example" {
  provider = catalystcenter

  parameters {

    id        = "string"
    name      = "string"
    parent_id = "string"
  }
}

output "catalystcenter_areas_example" {
  value = catalystcenter_areas.example
}