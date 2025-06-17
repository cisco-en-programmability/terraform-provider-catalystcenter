terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}
/*
data "catalystcenter_tag" "example" {
  provider                   = catalystcenter
  additional_info_attributes = "string"
  additional_info_name_space = "string"
  field                      = "1"
  level                      = "0"
  limit                      = "1"
  name                       = "WAN"
  offset                     = "10"
  order                      = "string"
  size                       = "1"
  sort_by                    = "string6"
  system_tag                 = "true"
}

output "catalystcenter_tag_example" {
  value = data.catalystcenter_tag.example.items
}
*/
data "catalystcenter_tag" "example" {
  provider = catalystcenter
  id       = "eacdbb1f-da5f-4acc-978f-bc723a5f29bc"
}

output "catalystcenter_tag_example" {
  value = data.catalystcenter_tag.example.item
}
