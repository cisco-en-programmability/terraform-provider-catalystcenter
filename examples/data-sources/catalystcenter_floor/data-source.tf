
data "catalystcenter_floor" "example" {
  provider = catalystcenter
  limit    = 1
  name     = "string"
  offset   = 1
  site_id  = "string"
  type     = "string"
}

output "catalystcenter_floor_example" {
  value = data.catalystcenter_floor.example.items
}
