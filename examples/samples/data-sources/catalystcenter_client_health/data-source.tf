
data "catalystcenter_client_health" "example" {
  provider  = catalystcenter
  timestamp = "string"
}

output "catalystcenter_client_health_example" {
  value = data.catalystcenter_client_health.example.items
}
