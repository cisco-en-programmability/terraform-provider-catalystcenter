terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}

data "catalystcenter_client_enrichment_details" "example" {
  provider = catalystcenter
}

output "catalystcenter_client_enrichment_details_example" {
  value = data.catalystcenter_client_enrichment_details.example.items
}
