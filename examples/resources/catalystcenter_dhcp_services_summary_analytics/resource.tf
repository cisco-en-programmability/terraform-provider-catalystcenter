
resource "catalystcenter_dhcp_services_summary_analytics" "example" {
    provider = meraki
    xca_lle_rid = "string"
    parameters = [{
      
      aggregate_attributes = [{
        
        function = "string"
        name = "string"
      }]
      attributes = ["string"]
      end_time = 1
      filters = [{
        
        filters = [{
          
          filters = ["string"]
          key = "string"
          logical_operator = "string"
          operator = "string"
          value = ------
        }]
        key = "string"
        logical_operator = "string"
        operator = "string"
        value = ------
      }]
      group_by = ["string"]
      page = [{
        
        limit = 1
        offset = 1
        sort_by = [{
          
          function = "string"
          name = "string"
          order = "string"
        }]
      }]
      start_time = 1
    }]
}

output "catalystcenter_dhcp_services_summary_analytics_example" {
    value = catalystcenter_dhcp_services_summary_analytics.example
}