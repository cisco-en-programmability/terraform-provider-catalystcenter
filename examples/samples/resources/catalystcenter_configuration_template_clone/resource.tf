terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_configuration_template_clone" "example" {
  provider = catalystcenter

  parameters {
    name        = "hola"
    project_id  = "42f4a526-7498-4593-88f5-d45e22e924ea"
    template_id = "fcfd4d19-99e2-494e-9c6f-0d85cf3094e5"
  }
}
