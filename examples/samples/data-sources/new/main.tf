

terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

data "catalystcenter_device_details" "example" {
  provider   = catalystcenter
  identifier = "uuid"
  search_by  = "57d9e4e4-e655-4512-a137-8f8c90e59ab1"
  #   timestamp  = "string"
}

output "catalystcenter_device_details_example" {
  value = data.catalystcenter_device_details.example.item
}
