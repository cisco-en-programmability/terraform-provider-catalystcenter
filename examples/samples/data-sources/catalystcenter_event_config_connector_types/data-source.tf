terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

data "catalystcenter_event_config_connector_types" "example" {
  provider = catalystcenter
}

output "catalystcenter_event_config_connector_types_example" {
  value = data.catalystcenter_event_config_connector_types.example.items
}
