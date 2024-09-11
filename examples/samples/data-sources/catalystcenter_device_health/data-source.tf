
data "catalystcenter_device_health" "example" {
  provider    = catalystcenter
  device_role = "string"
  end_time    = 1609459200
  health      = "string"
  limit       = 1
  offset      = 1
  site_id     = "string"
  start_time  = 1609459200
}

output "catalystcenter_device_health_example" {
  value = data.catalystcenter_device_health.example.items
}
