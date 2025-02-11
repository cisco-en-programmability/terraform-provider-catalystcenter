
data "catalystcenter_floors_floor_id_planned_access_point_positions_count" "example" {
    provider = catalystcenter
    floor_id = "string"
    mac_address = "string"
    name = "string"
    type = "string"
}

output "catalystcenter_floors_floor_id_planned_access_point_positions_count_example" {
    value = data.catalystcenter_floors_floor_id_planned_access_point_positions_count.example.item
}
