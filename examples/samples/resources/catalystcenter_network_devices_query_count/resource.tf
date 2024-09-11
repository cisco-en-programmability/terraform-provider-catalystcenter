
terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_network_devices_query_count" "example" {
  provider = catalystcenter
  parameters {

    # aggregate_attributes {

    #   function = "string"
    #   name     = "string"
    # }
    # attributes = ["string"]
    # end_time   = 1
    # filters {

    #   key      = "string"
    #   operator = "string"
    #   value    = "string"
    # }
    # page {

    #   limit  = 1
    #   offset = 1
    #   sort_by {

    #     name  = "string"
    #     order = "string"
    #   }
    # }
    start_time = 1
    # views      = ["string"]
  }
}

output "catalystcenter_network_devices_query_count_example" {
  value = catalystcenter_network_devices_query_count.example
}