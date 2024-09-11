
resource "catalystcenter_path_trace" "example" {
  provider = catalystcenter

  parameters {

    control_path     = "false"
    dest_ip          = "string"
    dest_port        = "string"
    flow_analysis_id = "string"
    inclusions       = ["string"]
    periodic_refresh = "false"
    protocol         = "string"
    source_ip        = "string"
    source_port      = "string"
  }
}

output "catalystcenter_path_trace_example" {
  value = catalystcenter_path_trace.example
}