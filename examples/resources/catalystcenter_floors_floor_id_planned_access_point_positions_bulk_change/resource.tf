
resource "catalystcenter_floors_floor_id_planned_access_point_positions_bulk_change" "example" {
    provider = meraki
    floor_id = "string"
    parameters = [{
      
      id = "string"
      mac_address = "string"
      name = "string"
      position = [{
        
        x = 1.0
        y = 1.0
        z = 1.0
      }]
      radios = [{
        
        antenna = [{
          
          azimuth = 1
          elevation = 1
          name = "string"
        }]
        channel = 1
        id = "string"
        tx_power = 1
      }]
      type = "string"
    }]
}

output "catalystcenter_floors_floor_id_planned_access_point_positions_bulk_change_example" {
    value = catalystcenter_floors_floor_id_planned_access_point_positions_bulk_change.example
}