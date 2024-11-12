
resource "catalystcenter_analytics_endpoints_anc_policy_update" "example" {
  provider = meraki
  ep_id    = "string"
  parameters {

    anc_policy = "string"
    granular_anc_policy {

      name           = "string"
      nas_ip_address = "string"
    }
  }
}

output "catalystcenter_analytics_endpoints_anc_policy_update_example" {
  value = catalystcenter_analytics_endpoints_anc_policy_update.example
}