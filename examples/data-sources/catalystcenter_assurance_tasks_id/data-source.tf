
data "catalystcenter_assurance_tasks_id" "example" {
    provider = catalystcenter
    id = "string"
    xca_lle_rid = "string"
}

output "catalystcenter_assurance_tasks_id_example" {
    value = data.catalystcenter_assurance_tasks_id.example.item
}
