
data "catalystcenter_wireless_controllers_anchor_capable_devices" "example" {
    provider = catalystcenter
}

output "catalystcenter_wireless_controllers_anchor_capable_devices_example" {
    value = data.catalystcenter_wireless_controllers_anchor_capable_devices.example.item
}
