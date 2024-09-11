
data "catalystcenter_area" "example" {
  provider = catalystcenter
  limit    = 1
  name     = "string"
  offset   = 1
  site_id  = "string"
  type     = "string"
}

output "catalystcenter_area_example" {
  value = data.catalystcenter_area.example.items
}
