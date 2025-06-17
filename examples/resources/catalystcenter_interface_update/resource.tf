
resource "catalystcenter_interface_update" "example" {
  provider        = catalystcenter
  deployment_mode = "string"
  interface_uuid  = "string"
  parameters = [{

    admin_status  = "string"
    description   = "string"
    vlan_id       = 1
    voice_vlan_id = 1
  }]
}

output "catalystcenter_interface_update_example" {
  value = catalystcenter_interface_update.example
}