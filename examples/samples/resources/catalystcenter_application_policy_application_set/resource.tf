

terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_application_policy_application_set" "example" {
  provider = catalystcenter
  parameters {
    payload {
      default_business_relevance = "default"
      # id                             = "string"
      name = "Terraform Test"
      # namespace                      = "string"
      # qualifier                      = "string"
      # scalable_group_external_handle = "string"
      # scalable_group_type            = "string"
      # type                           ="string"
    }
  }
}

output "catalystcenter_application_policy_application_set_example" {
  value = catalystcenter_application_policy_application_set.example
}