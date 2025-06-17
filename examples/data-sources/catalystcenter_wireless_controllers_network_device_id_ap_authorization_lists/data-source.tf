
data "catalystcenter_wireless_controllers_network_device_id_ap_authorization_lists" "example" {
  provider          = catalystcenter
  network_device_id = "string"
}

output "catalystcenter_wireless_controllers_network_device_id_ap_authorization_lists_example" {
  value = data.catalystcenter_wireless_controllers_network_device_id_ap_authorization_lists.example.item
}
