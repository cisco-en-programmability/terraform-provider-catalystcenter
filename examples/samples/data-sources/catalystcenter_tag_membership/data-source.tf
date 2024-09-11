
data "catalystcenter_tag_membership" "example" {
  provider                = catalystcenter
  id                      = "string"
  level                   = "string"
  limit                   = 1
  member_association_type = "string"
  member_type             = "string"
  offset                  = 1
}

output "catalystcenter_tag_membership_example" {
  value = data.catalystcenter_tag_membership.example.items
}
