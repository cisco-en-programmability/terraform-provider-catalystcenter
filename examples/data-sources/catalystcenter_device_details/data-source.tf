
data "catalystcenter_device_details" "example" {
    provider = catalystcenter
    identifier = "string"
    search_by = "string"
    timestamp = 1.0
}

output "catalystcenter_device_details_example" {
    value = data.catalystcenter_device_details.example.item
}
