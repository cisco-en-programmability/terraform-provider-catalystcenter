
data "catalystcenter_diagnostic_tasks_id" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_diagnostic_tasks_id_example" {
    value = data.catalystcenter_diagnostic_tasks_id.example.item
}
