
data "catalystcenter_diagnostic_tasks_id_detail" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_diagnostic_tasks_id_detail_example" {
  value = data.catalystcenter_diagnostic_tasks_id_detail.example.item
}
