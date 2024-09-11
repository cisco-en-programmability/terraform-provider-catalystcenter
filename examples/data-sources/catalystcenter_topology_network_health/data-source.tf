
data "catalystcenter_topology_network_health" "example" {
  provider  = catalystcenter
  timestamp = 1.0
}

output "catalystcenter_topology_network_health_example" {
  value = data.catalystcenter_topology_network_health.example.items
}
