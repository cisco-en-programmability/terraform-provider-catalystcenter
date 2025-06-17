
resource "catalystcenter_clients_query" "example" {
  provider    = catalystcenter
  xca_lle_rid = "string"
  parameters = [{

    aggregate_attributes = [{

      function = "string"
      name     = "string"
    }]
    attributes = ["string"]
    end_time   = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = 1
    }]
    page = [{

      limit  = 1
      offset = 1
      sort_by = [{

        name  = "string"
        order = "string"
      }]
    }]
    start_time = 1
    views      = ["string"]
  }]
}

output "catalystcenter_clients_query_example" {
  value = catalystcenter_clients_query.example
}