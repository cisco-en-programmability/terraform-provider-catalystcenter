
data "catalystcenter_reports_view_group" "example" {
  provider = catalystcenter
}

output "catalystcenter_reports_view_group_example" {
  value = data.catalystcenter_reports_view_group.example.items
}

data "catalystcenter_reports_view_group" "example" {
  provider      = catalystcenter
  view_group_id = "string"
}

output "catalystcenter_reports_view_group_example" {
  value = data.catalystcenter_reports_view_group.example.item
}
