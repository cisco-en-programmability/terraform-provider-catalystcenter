
resource "catalystcenter_device_configurations_export" "example" {
  provider = meraki
  parameters {

    device_id = "string"
    password  = "******"
  }
}

output "catalystcenter_device_configurations_export_example" {
  value = catalystcenter_device_configurations_export.example
}