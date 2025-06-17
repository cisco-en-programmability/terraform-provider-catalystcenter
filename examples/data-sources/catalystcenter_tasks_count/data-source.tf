
data "catalystcenter_tasks_count" "example" {
  provider   = catalystcenter
  end_time   = 1
  parent_id  = "string"
  root_id    = "string"
  start_time = 1
  status     = "string"
}

output "catalystcenter_tasks_count_example" {
  value = data.catalystcenter_tasks_count.example.item
}
