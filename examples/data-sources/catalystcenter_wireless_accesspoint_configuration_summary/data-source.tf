
data "catalystcenter_wireless_accesspoint_configuration_summary" "example" {
  provider = catalystcenter
  key      = "string"
}

output "catalystcenter_wireless_accesspoint_configuration_summary_example" {
  value = data.catalystcenter_wireless_accesspoint_configuration_summary.example.item
}
