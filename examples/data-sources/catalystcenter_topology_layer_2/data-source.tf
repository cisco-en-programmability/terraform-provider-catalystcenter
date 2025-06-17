
data "catalystcenter_topology_layer_2" "example" {
  provider = catalystcenter
  vlan_id  = "string"
}

output "catalystcenter_topology_layer_2_example" {
  value = data.catalystcenter_topology_layer_2.example.item
}
