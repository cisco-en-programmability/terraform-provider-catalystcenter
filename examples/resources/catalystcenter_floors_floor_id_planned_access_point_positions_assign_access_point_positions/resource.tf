
resource "catalystcenter_floors_floor_id_planned_access_point_positions_assign_access_point_positions" "example" {
    provider = meraki
    floor_id = "string"
    parameters = [{
      
      access_point_id = "string"
      planned_access_point_id = "string"
    }]
}

output "catalystcenter_floors_floor_id_planned_access_point_positions_assign_access_point_positions_example" {
    value = catalystcenter_floors_floor_id_planned_access_point_positions_assign_access_point_positions.example
}