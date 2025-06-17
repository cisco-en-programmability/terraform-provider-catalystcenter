
data "catalystcenter_tasks" "example" {
  provider   = catalystcenter
  end_time   = 1
  limit      = 1
  offset     = 1
  order      = "string"
  parent_id  = "string"
  root_id    = "string"
  sort_by    = "string"
  start_time = 1
  status     = "string"
}

output "catalystcenter_tasks_example" {
  value = data.catalystcenter_tasks.example.items
}

data "catalystcenter_tasks" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_tasks_example" {
  value = data.catalystcenter_tasks.example.item
}
