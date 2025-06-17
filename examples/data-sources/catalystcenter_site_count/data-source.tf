
data "catalystcenter_site_count" "example" {
  provider = catalystcenter
  site_id  = "string"
}

output "catalystcenter_site_count_example" {
  value = data.catalystcenter_site_count.example.item
}
