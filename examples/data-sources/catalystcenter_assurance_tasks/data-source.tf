
data "catalystcenter_assurance_tasks" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
    order = "string"
    sort_by = "string"
    status = "string"
    xca_lle_rid = "string"
}

output "catalystcenter_assurance_tasks_example" {
    value = data.catalystcenter_assurance_tasks.example.items
}
