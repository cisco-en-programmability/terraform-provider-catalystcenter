
data "catalystcenter_topology_physical" "example" {
    provider = catalystcenter
    node_type = "string"
}

output "catalystcenter_topology_physical_example" {
    value = data.catalystcenter_topology_physical.example.item
}
