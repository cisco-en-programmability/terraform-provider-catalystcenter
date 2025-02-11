
resource "catalystcenter_event_subscription_syslog" "example" {
    provider = catalystcenter

    parameters {

      description = "string"
      filter {

        categories = ["string"]
        domains_subdomains {

          domain = "string"
          sub_domains = ["string"]
        }
        event_ids = ["string"]
        severities = ["string"]
        site_ids = ["string"]
        sources = ["string"]
        types = ["string"]
      }
      name = "string"
      subscription_endpoints {

        instance_id = "string"
        subscription_details {

          connector_type = "string"
        }
      }
      subscription_id = "string"
      version = "string"
    }
}

output "catalystcenter_event_subscription_syslog_example" {
    value = catalystcenter_event_subscription_syslog.example
}
