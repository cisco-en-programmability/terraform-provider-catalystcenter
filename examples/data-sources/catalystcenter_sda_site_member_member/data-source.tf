
data "catalystcenter_sda_site_member_member" "example" {
  provider    = catalystcenter
  id          = "string"
  level       = "string"
  limit       = "string"
  member_type = "string"
  offset      = "string"
}

output "catalystcenter_sda_site_member_member_example" {
  value = data.catalystcenter_sda_site_member_member.example.items
}
