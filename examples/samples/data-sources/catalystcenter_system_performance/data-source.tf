
data "catalystcenter_system_performance" "example" {
  provider   = catalystcenter
  end_time   = 1609459200
  function   = "string"
  kpi        = "string"
  start_time = 1609459200
}

output "catalystcenter_system_performance_example" {
  value = data.catalystcenter_system_performance.example.item
}
