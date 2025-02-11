
data "catalystcenter_license_device_license_details" "example" {
    provider = catalystcenter
    device_uuid = "string"
}

output "catalystcenter_license_device_license_details_example" {
    value = data.catalystcenter_license_device_license_details.example.item
}
