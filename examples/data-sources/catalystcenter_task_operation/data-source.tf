
data "catalystcenter_task_operation" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
    operation_id = "string"
}

output "catalystcenter_task_operation_example" {
    value = data.catalystcenter_task_operation.example.items
}
