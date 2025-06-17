
data "catalystcenter_compliance_device" "example" {
  provider          = catalystcenter
  compliance_status = "string"
  device_uuid       = "string"
  limit             = 1
  offset            = 1
}

output "catalystcenter_compliance_device_example" {
  value = data.catalystcenter_compliance_device.example.items
}
