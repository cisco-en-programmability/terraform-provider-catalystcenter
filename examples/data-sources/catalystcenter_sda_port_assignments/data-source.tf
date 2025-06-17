
data "catalystcenter_sda_port_assignments" "example" {
  provider          = catalystcenter
  data_vlan_name    = "string"
  fabric_id         = "string"
  interface_name    = "string"
  limit             = 1
  native_vlan_id    = 1.0
  network_device_id = "string"
  offset            = 1
  voice_vlan_name   = "string"
}

output "catalystcenter_sda_port_assignments_example" {
  value = data.catalystcenter_sda_port_assignments.example.items
}
