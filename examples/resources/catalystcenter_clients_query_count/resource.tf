
resource "catalystcenter_clients_query_count" "example" {
  provider    = catalystcenter
  xca_lle_rid = "string"
  parameters = [{

    end_time = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = 1
    }]
    start_time = 1
  }]
}

output "catalystcenter_clients_query_count_example" {
  value = catalystcenter_clients_query_count.example
}