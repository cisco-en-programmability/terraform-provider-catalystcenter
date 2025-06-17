
resource "catalystcenter_icap_radios_id_stats" "example" {
  provider    = catalystcenter
  id          = "string"
  xca_lle_rid = "string"
  parameters = [{

    end_time = 1
    filters = [{

      key      = "string"
      operator = "string"
      value    = 1
    }]
    page = [{

      limit           = 1
      offset          = 1
      time_sort_order = "string"
    }]
    start_time = 1
  }]
}

output "catalystcenter_icap_radios_id_stats_example" {
  value = catalystcenter_icap_radios_id_stats.example
}