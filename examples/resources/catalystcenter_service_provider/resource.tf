
resource "catalystcenter_service_provider" "example" {
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
    sp_profile_name = "string"
  }
}

output "catalystcenter_service_provider_example" {
  value = catalystcenter_service_provider.example
}