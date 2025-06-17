
resource "catalystcenter_tag" "example" {
  provider = catalystcenter
 
  parameters {

    description = "string"
    dynamic_rules {

      member_type = "string"
      rules {

        items     = ["string"]
        name      = "string"
        operation = "string"
        value     = "string"
        values    = ["string"]
      }
    }
    id                 = "string"
    instance_tenant_id = "string"
    name               = "string"
    system_tag         = "false"
  }
}

output "catalystcenter_tag_example" {
  value = catalystcenter_tag.example
}