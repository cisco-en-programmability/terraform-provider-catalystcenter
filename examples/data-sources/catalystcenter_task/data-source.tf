
data "catalystcenter_task" "example" {
    provider = catalystcenter
    data = "string"
    end_time = "string"
    error_code = "string"
    failure_reason = "string"
    is_error = "string"
    limit = 1
    offset = 1
    order = "string"
    parent_id = "string"
    progress = "string"
    service_type = "string"
    sort_by = "string"
    start_time = "string"
    username = "string"
}

output "catalystcenter_task_example" {
    value = data.catalystcenter_task.example.items
}

data "catalystcenter_task" "example" {
    provider = catalystcenter
    task_id = "string"
}

output "catalystcenter_task_example" {
    value = data.catalystcenter_task.example.item
}
