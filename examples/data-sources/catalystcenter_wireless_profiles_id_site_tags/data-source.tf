
data "catalystcenter_wireless_profiles_id_site_tags" "example" {
  provider      = catalystcenter
  id            = "string"
  limit         = 1
  offset        = 1
  site_tag_name = "string"
}

output "catalystcenter_wireless_profiles_id_site_tags_example" {
  value = data.catalystcenter_wireless_profiles_id_site_tags.example.items
}
