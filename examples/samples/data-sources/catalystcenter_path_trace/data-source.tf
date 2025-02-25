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

data "catalystcenter_path_trace" "example" {
  provider         = catalystcenter
  dest_ip          = "192.168.1.145"
  dest_port        = "123"
  gt_create_time   = "1484241890"
  last_update_time = "1484241890"
  limit            = "1"
  lt_create_time   = "1484241890"
  offset           = "1"
  order            = "string7"
  periodic_refresh = "false"
  protocol         = "ospf"
  sort_by          = "Test"
  source_ip        = "192.168.1.145"
  source_port      = "123"
  status           = "string10"
  task_id          = "442368cb-bb21-4d01-b91a-543aba6287a4"
}

output "catalystcenter_path_trace_example" {
  value = data.catalystcenter_path_trace.example.items
}
/*
data "catalystcenter_path_trace" "example" {
  provider         = catalystcenter
  flow_analysis_id = "442368cb-bb21-4d01-b91a-543aba6287a4"
}

output "catalystcenter_path_trace_example" {
  value = data.catalystcenter_path_trace.example.item
}
*/
