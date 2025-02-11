
resource "catalystcenter_dhcp_services_query" "example" {
    provider = meraki
    xca_lle_rid = "string"
    parameters = [{
      
      end_time = 1
      filters = [{
        
        key = "string"
        operator = "string"
        value = ["string"]
      }]
      page = [{
        
        limit = 1
        offset = 1
        sort_by = [{
          
          name = "string"
          order = "string"
        }]
      }]
      start_time = 1
    }]
}

output "catalystcenter_dhcp_services_query_example" {
    value = catalystcenter_dhcp_services_query.example
}