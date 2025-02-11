
data "catalystcenter_task_tree" "example" {
    provider = catalystcenter
    task_id = "string"
}

output "catalystcenter_task_tree_example" {
    value = data.catalystcenter_task_tree.example.items
}
