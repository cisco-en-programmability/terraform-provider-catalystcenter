
data "catalystcenter_wireless_controllers_managed_ap_locations_count" "example" {
    provider = catalystcenter
    network_device_id = "string"
}

output "catalystcenter_wireless_controllers_managed_ap_locations_count_example" {
    value = data.catalystcenter_wireless_controllers_managed_ap_locations_count.example.item
}
