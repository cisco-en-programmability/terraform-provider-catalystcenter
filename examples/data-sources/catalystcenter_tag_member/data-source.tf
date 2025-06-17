
data "catalystcenter_tag_member" "example" {
  provider                = catalystcenter
  id                      = "string"
  level                   = "string"
  limit                   = 1
  member_association_type = "string"
  member_type             = "string"
  offset                  = 1
}

output "catalystcenter_tag_member_example" {
  value = data.catalystcenter_tag_member.example.items
}
