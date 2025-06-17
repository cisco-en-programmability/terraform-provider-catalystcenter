
resource "catalystcenter_sda_port_assignments" "example" {
  provider = catalystcenter
 
  parameters {

    allowed_vlan_ranges        = "string"
    authenticate_template_name = "string"
    connected_device_type      = "string"
    data_vlan_name             = "string"
    fabric_id                  = "string"
    id                         = "string"
    interface_description      = "string"
    interface_name             = "string"
    native_vlan_id             = 1
    network_device_id          = "string"
    scalable_group_name        = "string"
    security_group_name        = "string"
    voice_vlan_name            = "string"
  }
}

output "catalystcenter_sda_port_assignments_example" {
  value = catalystcenter_sda_port_assignments.example
}