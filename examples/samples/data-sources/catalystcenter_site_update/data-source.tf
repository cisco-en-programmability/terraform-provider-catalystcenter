terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
}

data "catalystcenter_site_update" "example" {
  provider = catalystcenter
  site_id  = "93d566d4-cdd9-447d-9dde-14c812dba13f"
  site {
    building {
      address     = "1dc6731f-6e17-4bc8-a9f1-e0f095cec64d"
      latitude    = 10
      longitude   = 10
      name        = "CDMX"
      parent_name = "Global/Mexico"
    }
  }
  type = "building"
}
