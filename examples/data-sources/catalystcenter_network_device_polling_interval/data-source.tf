
data "catalystcenter_network_device_polling_interval" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_network_device_polling_interval_example" {
    value = data.catalystcenter_network_device_polling_interval.example.item
}
