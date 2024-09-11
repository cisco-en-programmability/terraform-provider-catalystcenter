
data "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id_count" "example" {
  provider  = catalystcenter
  fabric_id = "string"
}

output "catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id_count_example" {
  value = data.catalystcenter_sda_fabrics_vlan_to_ssids_fabric_id_count.example.item
}
