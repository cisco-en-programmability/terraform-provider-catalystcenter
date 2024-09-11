
data "catalystcenter_compliance_device_by_id_detail" "example" {
  provider        = catalystcenter
  category        = "string"
  compliance_type = "string"
  device_uuid     = "string"
  diff_list       = "false"
  key             = "string"
  value           = "string"
}

output "catalystcenter_compliance_device_by_id_detail_example" {
  value = data.catalystcenter_compliance_device_by_id_detail.example.items
}
