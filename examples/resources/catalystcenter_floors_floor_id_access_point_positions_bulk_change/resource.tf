
resource "catalystcenter_floors_floor_id_access_point_positions_bulk_change" "example" {
  provider = catalystcenter
  floor_id = "string"
  parameters = [{

    id = "string"
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
      id = "string"
    }]
  }]
}

output "catalystcenter_floors_floor_id_access_point_positions_bulk_change_example" {
  value = catalystcenter_floors_floor_id_access_point_positions_bulk_change.example
}