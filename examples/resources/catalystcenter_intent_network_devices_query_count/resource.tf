
resource "catalystcenter_intent_network_devices_query_count" "example" {
  provider = catalystcenter
  parameters = [{

    filter = [{

      filters = [{

        key      = "string"
        operator = "string"
        value    = "string"
      }]
      logical_operator = "string"
    }]
  }]
}

output "catalystcenter_intent_network_devices_query_count_example" {
  value = catalystcenter_intent_network_devices_query_count.example
}