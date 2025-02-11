
data "catalystcenter_floors_floor_id_planned_access_point_positions" "example" {
    provider = catalystcenter
    floor_id = "string"
    limit = 1
    mac_address = "string"
    name = "string"
    offset = 1
    type = "string"
}

output "catalystcenter_floors_floor_id_planned_access_point_positions_example" {
    value = data.catalystcenter_floors_floor_id_planned_access_point_positions.example.items
}
