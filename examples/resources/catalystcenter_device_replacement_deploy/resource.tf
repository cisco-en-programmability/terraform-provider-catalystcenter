
resource "catalystcenter_device_replacement_deploy" "example" {
  provider = catalystcenter
  parameters {

    faulty_device_serial_number      = "string"
    replacement_device_serial_number = "string"
  }
}

output "catalystcenter_device_replacement_deploy_example" {
  value = catalystcenter_device_replacement_deploy.example
}