
data "catalystcenter_floors_floor_id_access_point_positions_count" "example" {
  provider    = catalystcenter
  floor_id    = "string"
  mac_address = "string"
  model       = "string"
  name        = "string"
  type        = "string"
}

output "catalystcenter_floors_floor_id_access_point_positions_count_example" {
  value = data.catalystcenter_floors_floor_id_access_point_positions_count.example.item
}
