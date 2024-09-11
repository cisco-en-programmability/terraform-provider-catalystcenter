
terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_itsm_integration_events_retry" "example" {
  provider = catalystcenter

  parameters {
    payload = ["string2"]
  }
}

output "catalystcenter_itsm_integration_events_retry_example" {
  value = catalystcenter_itsm_integration_events_retry.example
}


data "catalystcenter_dnacaap_management_execution_status" "example" {
  depends_on   = [catalystcenter_itsm_integration_events_retry.example]
  provider     = catalystcenter
  execution_id = catalystcenter_itsm_integration_events_retry.example.item.0.execution_id
}

output "catalystcenter_dnacaap_management_execution_status_example" {
  value = data.catalystcenter_dnacaap_management_execution_status.example.item
}
