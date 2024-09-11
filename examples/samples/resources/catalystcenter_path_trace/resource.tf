
terraform {
  required_providers {
    catalystcenter = {
      version = "1.0.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_path_trace" "example" {
  provider = catalystcenter
  parameters {

    control_path = "true"
    dest_ip      = "10.1.1.2"
    dest_port    = "8080"
    #flow_analysis_id = "false"
    #inclusions       = ["string"]
    periodic_refresh = "false"
    protocol         = "udp"
    source_ip        = "10.1.1.3"
    source_port      = "8082"
  }
}

output "catalystcenter_path_trace_example" {
  value = catalystcenter_path_trace.example
}
