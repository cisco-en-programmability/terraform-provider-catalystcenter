
resource "catalystcenter_event_webhook_create" "example" {
    provider = meraki
    parameters = [{
      
      description = "string"
      headers = [{
        
        default_value = "string"
        encrypt = "false"
        name = "string"
        value = "string"
      }]
      is_proxy_route = "false"
      method = "string"
      name = "string"
      trust_cert = "false"
      url = "string"
      webhook_id = "string"
    }]
}

output "catalystcenter_event_webhook_create_example" {
    value = catalystcenter_event_webhook_create.example
}