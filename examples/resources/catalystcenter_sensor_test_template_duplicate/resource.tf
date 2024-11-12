
resource "catalystcenter_sensor_test_template_duplicate" "example" {
  provider = meraki
  parameters {

    new_template_name = "string"
    template_name     = "string"
  }
}

output "catalystcenter_sensor_test_template_duplicate_example" {
  value = catalystcenter_sensor_test_template_duplicate.example
}