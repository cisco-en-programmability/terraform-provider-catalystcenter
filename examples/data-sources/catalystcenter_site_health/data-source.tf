
data "catalystcenter_site_health" "example" {
  provider  = catalystcenter
  limit     = 1
  offset    = 1
  site_type = "string"
  timestamp = 1.0
}

output "catalystcenter_site_health_example" {
  value = data.catalystcenter_site_health.example.items
}
