
resource "catalystcenter_dna_network_devices_query_count" "example" {
    provider = meraki
    parameters = [{
      
      end_time = 1
      filters = [{
        
        key = "string"
        operator = "string"
        value = "string"
      }]
      start_time = 1
    }]
}

output "catalystcenter_dna_network_devices_query_count_example" {
    value = catalystcenter_dna_network_devices_query_count.example
}