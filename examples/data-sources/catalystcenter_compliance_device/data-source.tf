
data "catalystcenter_compliance_device" "example" {
  provider          = catalystcenter
  compliance_status = "string"
  device_uuid       = "string"
}

output "catalystcenter_compliance_device_example" {
  value = data.catalystcenter_compliance_device.example.items
}
