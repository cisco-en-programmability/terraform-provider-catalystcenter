
resource "catalystcenter_template_preview" "example" {
  provider = catalystcenter
  parameters {

    device_id       = "string"
    params          = "string"
    resource_params = "string"
    template_id     = "string"
  }
}

output "catalystcenter_template_preview_example" {
  value = catalystcenter_template_preview.example
}