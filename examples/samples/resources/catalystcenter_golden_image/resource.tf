
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_golden_tag_image" "example" {
  provider = catalystcenter
  parameters {
    image_id                 = "e7f80aaa-62d3-4390-a8ee-49bbfba036a3"
    site_id                  = "2397da83-4e12-4d04-9bd3-a57b2ad91652"
    device_role              = "ALL"
    device_family_identifier = "284389362"
  }
}

output "catalystcenter_golden_image_example" {
  value = catalystcenter_golden_image.example
}
