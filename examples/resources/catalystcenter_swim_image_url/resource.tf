
resource "catalystcenter_swim_image_url" "example" {
  provider        = catalystcenter
  schedule_at     = "string"
  schedule_desc   = "string"
  schedule_origin = "string"
  parameters {

    application_type = "string"
    image_family     = "string"
    source_url       = "string"
    third_party      = "false"
    vendor           = "string"
  }
}

output "catalystcenter_swim_image_url_example" {
  value = catalystcenter_swim_image_url.example
}