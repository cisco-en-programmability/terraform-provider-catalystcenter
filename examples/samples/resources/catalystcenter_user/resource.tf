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

resource "catalystcenter_user" "example" {
  provider = catalystcenter
  parameters {

    email      = "john.doe@example.com"
    first_name = "John"
    last_name  = "Doe"
    password   = "Secure_password"
    role_list  = ["63060e76f5bc0031ecc68c14"]
    user_id    = "johndoe123"
    username   = "johndoe"
  }
}

output "catalystcenter_user_example" {
  sensitive = true
  value     = catalystcenter_user.example
}
