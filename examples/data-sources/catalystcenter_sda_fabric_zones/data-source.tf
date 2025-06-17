
data "catalystcenter_sda_fabric_zones" "example" {
  provider = catalystcenter
  id       = "string"
  limit    = 1
  offset   = 1
  site_id  = "string"
}

output "catalystcenter_sda_fabric_zones_example" {
  value = data.catalystcenter_sda_fabric_zones.example.items
}
