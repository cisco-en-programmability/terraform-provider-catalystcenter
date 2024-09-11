
resource "catalystcenter_sda_multicast_update" "example" {
  provider            = catalystcenter
  site_name_hierarchy = "string"
}

output "catalystcenter_sda_multicast_update_example" {
  value = catalystcenter_sda_multicast_update.example
}