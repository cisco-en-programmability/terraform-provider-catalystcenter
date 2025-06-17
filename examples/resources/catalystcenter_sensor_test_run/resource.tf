
resource "catalystcenter_sensor_test_run" "example" {
  provider = catalystcenter
  parameters = [{

    template_name = "string"
  }]
}

output "catalystcenter_sensor_test_run_example" {
  value = catalystcenter_sensor_test_run.example
}