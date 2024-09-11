
resource "catalystcenter_pnp_device_claim" "example" {
  provider = catalystcenter
  parameters {

    authorization_needed = "false"
    config_file_url      = "string"
    config_id            = "string"
    device_claim_list {

      config_list {

        config_id = "string"
        config_parameters {

          key   = "string"
          value = "string"
        }
      }
      device_id                  = "string"
      license_level              = "string"
      license_type               = "string"
      top_of_stack_serial_number = "string"
    }
    file_service_id    = "string"
    image_id           = "string"
    image_url          = "string"
    populate_inventory = "false"
    project_id         = "string"
    workflow_id        = "string"
  }
}

output "catalystcenter_pnp_device_claim_example" {
  value = catalystcenter_pnp_device_claim.example
}