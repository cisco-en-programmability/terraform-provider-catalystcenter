
resource "catalystcenter_aaa_services_query_count" "example" {
  provider    = catalystcenter
  xca_lle_rid = "string"
  parameters = [{

    end_time = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = ["string"]
    }]
    start_time = 1
  }]
}

output "catalystcenter_aaa_services_query_count_example" {
  value = catalystcenter_aaa_services_query_count.example
}