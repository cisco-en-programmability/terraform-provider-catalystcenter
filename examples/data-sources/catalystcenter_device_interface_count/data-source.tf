
data "catalystcenter_device_interface_count" "example" {
    provider = catalystcenter
}

output "catalystcenter_device_interface_count_example" {
    value = data.catalystcenter_device_interface_count.example.item
}
