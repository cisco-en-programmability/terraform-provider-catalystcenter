
resource "catalystcenter_network_device_export" "example" {
  provider = catalystcenter
  parameters {

    device_uuids   = ["string"]
    operation_enum = "string"
    parameters     = ["string"]
    password       = "******"
  }
}

output "catalystcenter_network_device_export_example" {
  value = catalystcenter_network_device_export.example
}