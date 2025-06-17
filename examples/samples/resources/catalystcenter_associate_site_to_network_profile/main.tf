
terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}
resource "catalystcenter_associate_site_to_network_profile" "example" {
  provider = catalystcenter

  parameters {
    network_profile_id = "string"
    site_id            = "string"
  }
}
