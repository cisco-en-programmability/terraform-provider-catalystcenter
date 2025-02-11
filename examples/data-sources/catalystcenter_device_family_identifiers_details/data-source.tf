
data "catalystcenter_device_family_identifiers_details" "example" {
    provider = catalystcenter
}

output "catalystcenter_device_family_identifiers_details_example" {
    value = data.catalystcenter_device_family_identifiers_details.example.items
}
