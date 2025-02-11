
data "catalystcenter_sda_device" "example" {
    provider = catalystcenter
    device_management_ip_address = "string"
}

output "catalystcenter_sda_device_example" {
    value = data.catalystcenter_sda_device.example.item
}
