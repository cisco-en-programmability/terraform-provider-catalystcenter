
data "catalystcenter_eox_status_device" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
}

output "catalystcenter_eox_status_device_example" {
    value = data.catalystcenter_eox_status_device.example.items
}

data "catalystcenter_eox_status_device" "example" {
    provider = catalystcenter
    device_id = "string"
}

output "catalystcenter_eox_status_device_example" {
    value = data.catalystcenter_eox_status_device.example.item
}
