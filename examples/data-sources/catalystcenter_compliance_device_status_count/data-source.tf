
data "catalystcenter_compliance_device_status_count" "example" {
  provider          = catalystcenter
  compliance_status = "string"
}

output "catalystcenter_compliance_device_status_count_example" {
  value = data.catalystcenter_compliance_device_status_count.example.item
}
