
data "catalystcenter_license_last_operation_status" "example" {
    provider = catalystcenter
}

output "catalystcenter_license_last_operation_status_example" {
    value = data.catalystcenter_license_last_operation_status.example.item
}
