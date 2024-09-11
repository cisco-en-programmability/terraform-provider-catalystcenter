
data "catalystcenter_system_performance_historical" "example" {
  provider   = catalystcenter
  end_time   = 1609459200
  kpi        = "string"
  start_time = 1609459200
}

output "catalystcenter_system_performance_historical_example" {
  value = data.catalystcenter_system_performance_historical.example.item
}
