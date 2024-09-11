
resource "catalystcenter_sda_extranet_policies" "example" {
  provider = catalystcenter

  parameters {

    extranet_policy_name             = "string"
    fabric_ids                       = ["string"]
    id                               = "string"
    provider_virtual_network_name    = "string"
    subscriber_virtual_network_names = ["string"]
  }
}

output "catalystcenter_sda_extranet_policies_example" {
  value = catalystcenter_sda_extranet_policies.example
}