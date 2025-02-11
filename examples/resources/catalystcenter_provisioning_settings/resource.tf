
resource "catalystcenter_provisioning_settings" "example" {
    provider = catalystcenter

    parameters {

      require_itsm_approval = "false"
      require_preview = "false"
    }
}

output "catalystcenter_provisioning_settings_example" {
    value = catalystcenter_provisioning_settings.example
}
