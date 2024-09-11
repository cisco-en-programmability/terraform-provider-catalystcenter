
resource "catalystcenter_service_provider_v2" "example" {
  provider = catalystcenter

  parameters {

    settings {

      qos {

        model            = "string"
        old_profile_name = "string"
        profile_name     = "string"
        wan_provider     = "string"
      }
    }
  }
}

output "catalystcenter_service_provider_v2_example" {
  value = catalystcenter_service_provider_v2.example
}