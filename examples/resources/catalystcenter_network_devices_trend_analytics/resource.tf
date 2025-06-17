
resource "catalystcenter_network_devices_trend_analytics" "example" {
  provider = catalystcenter
  parameters = [{

    aggregate_attributes = ["string"]
    attributes           = ["string"]
    end_time             = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = "string"
    }]
    group_by = ["string"]
    page = [{

      limit           = 1
      offset          = 1
      timestamp_order = "string"
    }]
    start_time     = 1
    trend_interval = "string"
  }]
}

output "catalystcenter_network_devices_trend_analytics_example" {
  value = catalystcenter_network_devices_trend_analytics.example
}