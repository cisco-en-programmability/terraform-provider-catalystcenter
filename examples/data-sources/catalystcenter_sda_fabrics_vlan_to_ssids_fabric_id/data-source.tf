
data "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id" "example" {
  provider  = catalystcenter
  fabric_id = "string"
  limit     = 1
  offset    = 1
}

output "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id_example" {
  value = data.catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id.example.items
}
