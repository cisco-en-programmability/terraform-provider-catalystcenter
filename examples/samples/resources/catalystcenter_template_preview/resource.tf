terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_template_preview" "example" {
  provider = catalystcenter

  parameters {
    device_id       = "string"
    params          = "string"
    resource_params = ["string"]
    template_id     = "string"
  }
}
