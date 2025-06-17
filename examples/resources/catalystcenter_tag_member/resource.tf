
resource "catalystcenter_tag_member" "example" {
  provider = catalystcenter
 
  parameters {

    id          = "string"
    member_id   = "string"
    member_type = ["string"]
  }
}

output "catalystcenter_tag_member_example" {
  value = catalystcenter_tag_member.example
}