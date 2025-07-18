
resource "catalystcenter_network_applications_trend_analytics" "example" {
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
    group_by = ["string"]
    page = [{

      cursor          = "string"
      limit           = 1
      time_sort_order = "string"
    }]
    site_ids       = ["string"]
    start_time     = 1
    trend_interval = "string"
  }]
}

output "catalystcenter_network_applications_trend_analytics_example" {
  value = catalystcenter_network_applications_trend_analytics.example
}