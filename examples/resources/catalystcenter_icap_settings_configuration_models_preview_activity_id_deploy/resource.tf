
resource "catalystcenter_icap_settings_configuration_models_preview_activity_id_deploy" "example" {
    provider = meraki
    preview_activity_id = "string"
    parameters = [{
      
      object = "string"
    }]
}

output "catalystcenter_icap_settings_configuration_models_preview_activity_id_deploy_example" {
    value = catalystcenter_icap_settings_configuration_models_preview_activity_id_deploy.example
}