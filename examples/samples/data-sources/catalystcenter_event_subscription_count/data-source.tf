terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}

data "catalystcenter_event_subscription_count" "example" {
  provider  = catalystcenter
  event_ids = "string"
}

output "catalystcenter_event_subscription_count_example" {
  value = data.catalystcenter_event_subscription_count.example.item
}
