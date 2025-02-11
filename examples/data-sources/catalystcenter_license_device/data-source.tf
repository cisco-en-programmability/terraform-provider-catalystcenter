
data "catalystcenter_license_device" "example" {
    provider = catalystcenter
}

output "catalystcenter_license_device_example" {
    value = data.catalystcenter_license_device.example.items
}
