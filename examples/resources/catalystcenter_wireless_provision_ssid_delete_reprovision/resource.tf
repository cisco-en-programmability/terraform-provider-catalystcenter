
resource "catalystcenter_wireless_provision_ssid_delete_reprovision" "example" {
  provider            = catalystcenter
  managed_aplocations = "string"
  ssid_name           = "string"
  parameters = [{

  }]
}

output "catalystcenter_wireless_provision_ssid_delete_reprovision_example" {
  value = catalystcenter_wireless_provision_ssid_delete_reprovision.example
}