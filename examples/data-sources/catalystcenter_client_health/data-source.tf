
data "catalystcenter_client_health" "example" {
  provider  = catalystcenter
  timestamp = 1.0
}

output "catalystcenter_client_health_example" {
  value = data.catalystcenter_client_health.example.items
}
