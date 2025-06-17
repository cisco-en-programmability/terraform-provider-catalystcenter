
data "catalystcenter_wireless_controllers_ssid_details" "example" {
  provider          = catalystcenter
  admin_status      = "false"
  limit             = 1
  managed           = "false"
  network_device_id = "string"
  offset            = 1
  ssid_name         = "string"
}

output "catalystcenter_wireless_controllers_ssid_details_example" {
  value = data.catalystcenter_wireless_controllers_ssid_details.example.items
}
