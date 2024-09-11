
resource "catalystcenter_event_webhook_update" "example" {
  provider = catalystcenter
  parameters {

    description = "string"
    headers {

      default_value = "string"
      encrypt       = "false"
      name          = "string"
      value         = "string"
    }
    is_proxy_route = "false"
    method         = "string"
    name           = "string"
    trust_cert     = "false"
    url            = "string"
    webhook_id     = "string"
  }
}

output "catalystcenter_event_webhook_update_example" {
  value = catalystcenter_event_webhook_update.example
}