
resource "catalystcenter_security_rogue_wireless_containment_start" "example" {
  provider = catalystcenter
  parameters = [{

    mac_address = "string"
    type        = 1
  }]
}

output "catalystcenter_security_rogue_wireless_containment_start_example" {
  value = catalystcenter_security_rogue_wireless_containment_start.example
}