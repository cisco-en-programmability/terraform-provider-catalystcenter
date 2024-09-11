
resource "catalystcenter_assurance_events_query_count" "example" {
  provider    = catalystcenter
  xca_lle_rid = "string"
  parameters {

    device_family = ["string"]
    end_time      = 1
    filters {

      key      = "string"
      operator = "string"
      value    = "string"
    }
    start_time = 1
  }
}

output "catalystcenter_assurance_events_query_count_example" {
  value = catalystcenter_assurance_events_query_count.example
}