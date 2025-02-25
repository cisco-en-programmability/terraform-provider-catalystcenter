
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

resource "catalystcenter_service_provider" "example" {
  provider = catalystcenter
  parameters {

    settings {
      qos {
        profile_name = "Test_tf_new"
        model        = "8-class-model"
        wan_provider = "test1-provider"
      }
    }
  }
}

output "catalystcenter_service_provider_example" {
  value = catalystcenter_service_provider.example
}
