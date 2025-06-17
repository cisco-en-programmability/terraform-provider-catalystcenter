
resource "catalystcenter_icap_settings_deploy" "example" {
  provider            = catalystcenter
  preview_description = "string"
  parameters = [{

    apid              = "string"
    capture_type      = "string"
    client_mac        = "string"
    duration_in_mins  = 1
    ota_band          = "string"
    ota_channel       = 1
    ota_channel_width = 1
    slot              = []
    wlc_id            = "string"
  }]
}

output "catalystcenter_icap_settings_deploy_example" {
  value = catalystcenter_icap_settings_deploy.example
}