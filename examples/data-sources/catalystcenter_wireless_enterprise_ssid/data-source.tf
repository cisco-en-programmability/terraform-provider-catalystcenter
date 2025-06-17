
data "catalystcenter_wireless_enterprise_ssid" "example" {
  provider  = catalystcenter
  ssid_name = "string"
}

output "catalystcenter_wireless_enterprise_ssid_example" {
  value = data.catalystcenter_wireless_enterprise_ssid.example.items
}
