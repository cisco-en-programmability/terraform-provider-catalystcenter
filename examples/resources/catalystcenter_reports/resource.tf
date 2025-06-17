
resource "catalystcenter_reports" "example" {
  provider = catalystcenter
 
  parameters {

    data_category = "string"
    deliveries    = ["string"]
    name          = "string"
    report_id     = "string"
    schedule      = "string"
    tags          = ["string"]
    view {

      field_groups {

        field_group_display_name = "string"
        field_group_name         = "string"
        fields {

          display_name = "string"
          name         = "string"
        }
      }
      filters {

        display_name = "string"
        name         = "string"
        type         = "string"
        value        = "string"
      }
      format {

        format_type = "string"
        name        = "string"
      }
      name    = "string"
      view_id = "string"
    }
    view_group_id      = "string"
    view_group_version = "string"
  }
}

output "catalystcenter_reports_example" {
  value = catalystcenter_reports.example
}