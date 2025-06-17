terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}



data "catalystcenter_application_policy_application_set" "example" {
  provider = catalystcenter
  #  attributes = "APPLICATION_GROUP"
  # limit      = 1
  # name       = "string"
  # offset     = 1
}

output "catalystcenter_application_policy_application_set_example" {
  value = data.catalystcenter_application_policy_application_set.example.items
}
