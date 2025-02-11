
resource "catalystcenter_tags_interfaces_members_associations_bulk" "example" {
    provider = meraki
    parameters = [{
      
      id = "string"
      tags = [{
        
        id = "string"
      }]
    }]
}

output "catalystcenter_tags_interfaces_members_associations_bulk_example" {
    value = catalystcenter_tags_interfaces_members_associations_bulk.example
}