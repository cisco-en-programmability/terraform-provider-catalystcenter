
data "catalystcenter_site_count_v2" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_site_count_v2_example" {
  value = data.catalystcenter_site_count_v2.example.item
}
