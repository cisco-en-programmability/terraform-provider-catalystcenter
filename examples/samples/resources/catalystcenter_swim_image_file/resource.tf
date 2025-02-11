
terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_swim_image_file" "example" {
  provider = catalystcenter
  parameters {
    file_path = "$PATH/terraform-provider-catalystcenter/examples/samples/resources/catalystcenter_swim_image_file/testIMG1.zip"
    file_name = "testIMG1.zip"
  }
}

output "catalystcenter_swim_image_file_example" {
  value = catalystcenter_swim_image_file.example
}
