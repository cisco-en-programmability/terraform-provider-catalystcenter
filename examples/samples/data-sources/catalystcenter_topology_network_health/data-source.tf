
data "catalystcenter_topology_network_health" "example" {
  provider  = catalystcenter
  timestamp = "string"
}

output "catalystcenter_topology_network_health_example" {
  value = data.catalystcenter_topology_network_health.example.items
}
