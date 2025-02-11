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

data "dnaccatalystcenter_swim_trigger_distribution" "example" {
  provider = catalystcenter
  payload {

    device_uuid = "3eb928b8-2414-4121-ac35-1247e5d666a4"
    image_uuid  = "6af2b040-a312-4f57-8c8e-21f5e3e07597"
  }
}
