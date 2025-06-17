
resource "catalystcenter_floors_floor_id_planned_access_point_positions_bulk" "example" {
  provider = catalystcenter
  floor_id = "string"
  parameters = [{

    mac_address = "string"
    name        = "string"
    position = [{

      x = 1.0
      y = 1.0
      z = 1.0
    }]
    radios = [{

      antenna = [{

        azimuth   = 1
        elevation = 1
        name      = "string"
      }]
      bands    = []
      channel  = 1
      tx_power = 1
    }]
    type = "string"
  }]
}

output "catalystcenter_floors_floor_id_planned_access_point_positions_bulk_example" {
  value = catalystcenter_floors_floor_id_planned_access_point_positions_bulk.example
}