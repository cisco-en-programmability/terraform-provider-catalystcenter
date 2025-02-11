
data "catalystcenter_network_device_wireless_lan" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_network_device_wireless_lan_example" {
    value = data.catalystcenter_network_device_wireless_lan.example.item
}
