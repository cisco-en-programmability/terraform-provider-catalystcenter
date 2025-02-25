terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}


resource "catalystcenter_custom_issue_definitions" "example" {
  provider = catalystcenter

  parameters {

    description             = "string"
    id                      = "string"
    is_enabled              = "false"
    is_notification_enabled = "false"
    name                    = "string"
    priority                = "string"
    rules {

      duration_in_minutes = 1
      facility            = "string"
      mnemonic            = "string"
      occurrences         = 1
      pattern             = "string"
      severity            = 1
    }
  }
}

output "catalystcenter_custom_issue_definitions_example" {
  value = catalystcenter_custom_issue_definitions.example
}
