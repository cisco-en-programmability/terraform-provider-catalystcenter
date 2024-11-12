
resource "catalystcenter_application_set_v2" "example" {
  provider = catalystcenter

  parameters {

    default_business_relevance     = "string"
    id                             = "string"
    name                           = "string"
    namespace                      = "string"
    qualifier                      = "string"
    scalable_group_external_handle = "string"
    scalable_group_type            = "string"
    type                           = "string"
  }
}

output "catalystcenter_application_set_v2_example" {
  value = catalystcenter_application_set_v2.example
}