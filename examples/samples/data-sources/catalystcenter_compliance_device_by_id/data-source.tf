
data "catalystcenter_compliance_device_by_id" "example" {
  provider    = catalystcenter
  device_uuid = "string"
}

output "catalystcenter_compliance_device_by_id_example" {
  value = data.catalystcenter_compliance_device_by_id.example.item
}
