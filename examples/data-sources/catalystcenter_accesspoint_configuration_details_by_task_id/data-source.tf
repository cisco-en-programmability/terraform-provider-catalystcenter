
data "catalystcenter_accesspoint_configuration_details_by_task_id" "example" {
  provider = catalystcenter
  task_id  = "string"
}

output "catalystcenter_accesspoint_configuration_details_by_task_id_example" {
  value = data.catalystcenter_accesspoint_configuration_details_by_task_id.example.items
}
