
data "catalystcenter_site_v2" "example" {
  provider             = catalystcenter
  group_name_hierarchy = "string"
  id                   = "string"
  limit                = "string"
  offset               = "string"
  type                 = "string"
}

output "catalystcenter_site_v2_example" {
  value = data.catalystcenter_site_v2.example.items
}
