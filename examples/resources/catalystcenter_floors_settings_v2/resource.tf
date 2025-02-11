
resource "catalystcenter_floors_settings_v2" "example" {
    provider = catalystcenter
    parameters {

      units_of_measure = "string"
    }
}

output "catalystcenter_floors_settings_v2_example" {
    value = catalystcenter_floors_settings_v2.example
}
