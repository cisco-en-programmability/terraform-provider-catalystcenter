
data "catalystcenter_sda_fabrics_vlan_to_ssids_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_sda_fabrics_vlan_to_ssids_count_example" {
  value = data.catalystcenter_sda_fabrics_vlan_to_ssids_count.example.item
}
