
data "catalystcenter_planned_access_points" "example" {
  provider = catalystcenter
  floor_id = "string"
  limit    = 1
  offset   = 1
  radios   = "false"
}

output "catalystcenter_planned_access_points_example" {
  value = data.catalystcenter_planned_access_points.example.items
}
