
data "catalystcenter_assurance_tasks_count" "example" {
  provider    = catalystcenter
  status      = "string"
  xca_lle_rid = "string"
}

output "catalystcenter_assurance_tasks_count_example" {
  value = data.catalystcenter_assurance_tasks_count.example.item
}
