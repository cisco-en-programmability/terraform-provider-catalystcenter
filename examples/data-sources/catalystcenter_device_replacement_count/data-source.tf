
data "catalystcenter_device_replacement_count" "example" {
    provider = catalystcenter
    replacement_status = ["string"]
}

output "catalystcenter_device_replacement_count_example" {
    value = data.catalystcenter_device_replacement_count.example.item
}
