
resource "catalystcenter_security_rogue_wireless_containment_stop" "example" {
  provider = catalystcenter
  parameters = [{

    mac_address = "string"
    type        = 1
    wlc_ip      = "string"
  }]
}

output "catalystcenter_security_rogue_wireless_containment_stop_example" {
  value = catalystcenter_security_rogue_wireless_containment_stop.example
}