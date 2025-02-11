
data "catalystcenter_topology_layer_3" "example" {
    provider = catalystcenter
    topology_type = "string"
}

output "catalystcenter_topology_layer_3_example" {
    value = data.catalystcenter_topology_layer_3.example.item
}
