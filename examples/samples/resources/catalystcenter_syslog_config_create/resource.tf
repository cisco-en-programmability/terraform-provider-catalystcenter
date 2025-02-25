terraform {
  required_providers {
    catalystcenter = {
      version = "1.1.1-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}
resource "catalystcenter_syslog_config_create" "example" {
  provider = catalystcenter
  parameters {

    //config_id = "string"
    description = "string"
    host        = "string"
    name        = "string"
    port        = "string"
    protocol    = "string"
  }
}

output "catalystcenter_syslog_config_create_example" {
  value = catalystcenter_syslog_config_create.example
}
