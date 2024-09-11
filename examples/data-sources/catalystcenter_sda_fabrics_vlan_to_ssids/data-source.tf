
data "catalystcenter_sda_fabrics_vlan_to_ssids" "example" {
  provider = catalystcenter
  limit    = 1
  offset   = 1
}

output "catalystcenter_sda_fabrics_vlan_to_ssids_example" {
  value = data.catalystcenter_sda_fabrics_vlan_to_ssids.example.items
}
