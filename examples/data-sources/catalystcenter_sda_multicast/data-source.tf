
data "catalystcenter_sda_multicast" "example" {
  provider            = catalystcenter
  site_name_hierarchy = "string"
}

output "catalystcenter_sda_multicast_example" {
  value = data.catalystcenter_sda_multicast.example.item
}
