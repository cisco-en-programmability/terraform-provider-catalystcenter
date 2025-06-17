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

data "catalystcenter_tag_member" "example" {
  provider                = catalystcenter
  id                      = "8f987b52-8698-45eb-9a80-bc5a02d4d972"
  level                   = "string"
  limit                   = 1
  member_association_type = "string"
  member_type             = "template"
  offset                  = 1
}

output "catalystcenter_tag_member_example" {
  value = data.catalystcenter_tag_member.example.items
}
