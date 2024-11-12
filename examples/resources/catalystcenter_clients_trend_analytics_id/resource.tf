
resource "catalystcenter_clients_trend_analytics_id" "example" {
  provider    = meraki
  id          = "string"
  xca_lle_rid = "string"
  parameters {

    aggregate_attributes {

      function = "string"
      name     = "string"
    }
    attributes = ["string"]
    end_time   = 1
    filters {

      key      = "string"
      operator = "string"
      value    = 1
    }
    group_by = ["string"]
    page {

      cursor          = "string"
      limit           = 1
      time_sort_order = "string"
    }
    start_time     = 1
    trend_interval = "string"
  }
}

output "catalystcenter_clients_trend_analytics_id_example" {
  value = catalystcenter_clients_trend_analytics_id.example
}