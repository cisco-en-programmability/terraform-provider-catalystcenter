
data "catalystcenter_tasks_detail" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_tasks_detail_example" {
  value = data.catalystcenter_tasks_detail.example.item
}
