
data "catalystcenter_site" "example" {
  provider = catalystcenter
  limit    = 1
  name     = "string"
  offset   = 1
  site_id  = "string"
  type     = "string"
}

output "catalystcenter_site_example" {
  value = data.catalystcenter_site.example.items
}
