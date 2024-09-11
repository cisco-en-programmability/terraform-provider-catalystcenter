
resource "catalystcenter_floors_settings" "example" {
  provider = catalystcenter

  parameters {

    units_of_measure = "string"
  }
}

output "catalystcenter_floors_settings_example" {
  value = catalystcenter_floors_settings.example
}