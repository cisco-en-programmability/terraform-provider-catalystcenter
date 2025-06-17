
data "catalystcenter_compliance_device_details" "example" {
  provider          = catalystcenter
  compliance_status = "string"
  compliance_type   = "string"
  device_uuid       = "string"
  limit             = 1
  offset            = 1
}

output "catalystcenter_compliance_device_details_example" {
  value = data.catalystcenter_compliance_device_details.example.items
}
