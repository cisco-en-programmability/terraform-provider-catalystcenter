
data "catalystcenter_wireless_profiles_id_site_tags_count" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_wireless_profiles_id_site_tags_count_example" {
  value = data.catalystcenter_wireless_profiles_id_site_tags_count.example.item
}
