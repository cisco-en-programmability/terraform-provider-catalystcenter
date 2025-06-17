
resource "catalystcenter_network_devices_top_n_analytics" "example" {
  provider = catalystcenter
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
      value    = "string"
    }]
    group_by = ["string"]
    page = [{

      limit  = 1
      offset = 1
      sort_by = [{

        function = "string"
        name     = "string"
        order    = "string"
      }]
    }]
    start_time = 1
    top_n      = 1
  }]
}

output "catalystcenter_network_devices_top_n_analytics_example" {
  value = catalystcenter_network_devices_top_n_analytics.example
}