
data "catalystcenter_topology_vlan_details" "example" {
    provider = catalystcenter
}

output "catalystcenter_topology_vlan_details_example" {
    value = data.catalystcenter_topology_vlan_details.example.items
}
