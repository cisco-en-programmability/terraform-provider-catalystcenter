
data "catalystcenter_network_device_replacements_id" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_network_device_replacements_id_example" {
    value = data.catalystcenter_network_device_replacements_id.example.item
}
