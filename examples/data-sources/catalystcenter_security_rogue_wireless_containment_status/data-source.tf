
data "catalystcenter_security_rogue_wireless_containment_status" "example" {
  provider    = catalystcenter
  mac_address = "string"
}

output "catalystcenter_security_rogue_wireless_containment_status_example" {
  value = data.catalystcenter_security_rogue_wireless_containment_status.example.items
}
