
resource "catalystcenter_wireless_dynamic_interface" "example" {
  provider = catalystcenter

  parameters {

    interface_name = "string"
    vlan_id        = 1.0
  }
}

output "catalystcenter_wireless_dynamic_interface_example" {
  value = catalystcenter_wireless_dynamic_interface.example
}