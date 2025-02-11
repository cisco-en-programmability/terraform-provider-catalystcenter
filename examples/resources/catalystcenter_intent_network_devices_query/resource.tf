
resource "catalystcenter_intent_network_devices_query" "example" {
    provider = meraki
    parameters = [{
      
      filter = [{
        
        filters = [{
          
          key = "string"
          operator = "string"
          value = ------
        }]
        logical_operator = "string"
      }]
      page = [{
        
        limit = 1
        offset = 1
        sort_by = [{
          
          name = "string"
          order = "string"
        }]
      }]
      views = ["string"]
    }]
}

output "catalystcenter_intent_network_devices_query_example" {
    value = catalystcenter_intent_network_devices_query.example
}