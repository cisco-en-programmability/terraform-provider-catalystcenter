
data "catalystcenter_analytics_tasks" "example" {
    provider = catalystcenter
    task_id = "string"
}

output "catalystcenter_analytics_tasks_example" {
    value = data.catalystcenter_analytics_tasks.example.item
}
