
data "catalystcenter_reports_view_group_view" "example" {
  provider      = catalystcenter
  view_group_id = "string"
  view_id       = "string"
}

output "catalystcenter_reports_view_group_view_example" {
  value = data.catalystcenter_reports_view_group_view.example.item
}
