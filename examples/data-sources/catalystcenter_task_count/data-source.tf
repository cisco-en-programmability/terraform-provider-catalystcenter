
data "catalystcenter_task_count" "example" {
  provider       = catalystcenter
  data           = "string"
  end_time       = "string"
  error_code     = "string"
  failure_reason = "string"
  is_error       = "string"
  parent_id      = "string"
  progress       = "string"
  service_type   = "string"
  start_time     = "string"
  username       = "string"
}

output "catalystcenter_task_count_example" {
  value = data.catalystcenter_task_count.example.item
}
