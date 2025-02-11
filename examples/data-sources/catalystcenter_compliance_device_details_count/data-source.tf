
data "catalystcenter_compliance_device_details_count" "example" {
    provider = catalystcenter
    compliance_status = "string"
    compliance_type = "string"
}

output "catalystcenter_compliance_device_details_count_example" {
    value = data.catalystcenter_compliance_device_details_count.example.item
}
