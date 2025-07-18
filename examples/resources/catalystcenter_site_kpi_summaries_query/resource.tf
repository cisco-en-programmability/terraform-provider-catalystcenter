
resource "catalystcenter_site_kpi_summaries_query" "example" {
  provider    = catalystcenter
  xca_lle_rid = "string"
  parameters = [{

    attributes = ["string"]
    end_time   = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = "string"
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

output "catalystcenter_site_kpi_summaries_query_example" {
  value = catalystcenter_site_kpi_summaries_query.example
}