
data "catalystcenter_sda_fabric_sites" "example" {
  provider = catalystcenter
  id       = "string"
  limit    = 1
  offset   = 1
  site_id  = "string"
}

output "catalystcenter_sda_fabric_sites_example" {
  value = data.catalystcenter_sda_fabric_sites.example.items
}
