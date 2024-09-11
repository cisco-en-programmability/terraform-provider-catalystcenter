
data "catalystcenter_site_health" "example" {
  provider  = catalystcenter
  limit     = 1
  offset    = 1
  site_type = "string"
  timestamp = "string"
}

output "catalystcenter_site_health_example" {
  value = data.catalystcenter_site_health.example.items
}
