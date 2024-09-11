
resource "catalystcenter_network_devices_interfaces_query" "example" {
  provider  = catalystcenter
  device_id = "string"
  parameters {

    end_time = 1
    query {

      fields = ["string"]
      filters {

        key      = "string"
        operator = "string"
        value    = "string"
      }
      page {

        limit  = 1
        offset = 1
        order_by {

          name  = "string"
          order = "string"
        }
      }
    }
    start_time = 1
  }
}

output "catalystcenter_network_devices_interfaces_query_example" {
  value = catalystcenter_network_devices_interfaces_query.example
}