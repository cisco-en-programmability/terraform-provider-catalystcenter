
data "catalystcenter_wireless_controllers_ssid_details_count" "example" {
  provider          = catalystcenter
  admin_status      = "false"
  managed           = "false"
  network_device_id = "string"
}

output "catalystcenter_wireless_controllers_ssid_details_count_example" {
  value = data.catalystcenter_wireless_controllers_ssid_details_count.example.item
}
