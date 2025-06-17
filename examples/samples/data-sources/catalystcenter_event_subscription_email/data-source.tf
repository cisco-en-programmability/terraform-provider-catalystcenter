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

data "catalystcenter_event_subscription_email" "example" {
  provider  = catalystcenter
  event_ids = "string"
  limit     = 1
  offset    = 1
  order     = "string"
  sort_by   = "string"
}

output "catalystcenter_event_subscription_email_example" {
  value = data.catalystcenter_event_subscription_email.example.items
}
