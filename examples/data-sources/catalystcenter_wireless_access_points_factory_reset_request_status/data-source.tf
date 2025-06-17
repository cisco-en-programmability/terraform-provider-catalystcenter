
data "catalystcenter_wireless_access_points_factory_reset_request_status" "example" {
  provider = catalystcenter
  task_id  = "string"
}

output "catalystcenter_wireless_access_points_factory_reset_request_status_example" {
  value = data.catalystcenter_wireless_access_points_factory_reset_request_status.example.items
}
