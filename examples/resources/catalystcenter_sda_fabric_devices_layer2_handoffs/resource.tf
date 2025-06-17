
resource "catalystcenter_sda_fabric_devices_layer2_handoffs" "example" {
  provider = catalystcenter
 
  parameters {

    external_vlan_id  = 1
    fabric_id         = "string"
    id                = "string"
    interface_name    = "string"
    internal_vlan_id  = 1
    network_device_id = "string"
  }
}

output "catalystcenter_sda_fabric_devices_layer2_handoffs_example" {
  value = catalystcenter_sda_fabric_devices_layer2_handoffs.example
}