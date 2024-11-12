
resource "catalystcenter_reports" "example" {
  provider = catalystcenter
  item {




    executions {








    }





    view {


      field_groups {



        fields {



        }
      }
      filters {





      }
      format {




      }



    }


  }
  parameters {

    data_category = "string"
    deliveries    = ["string"]
    name          = "string"
    report_id     = "string"
    schedule      = "------"
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
        value        = "------"
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