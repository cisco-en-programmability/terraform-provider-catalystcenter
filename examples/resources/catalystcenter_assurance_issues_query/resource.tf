
resource "catalystcenter_assurance_issues_query" "example" {
  provider        = catalystcenter
  accept_language = "string"
  xca_lle_rid     = "string"
  parameters = [{

    end_time = 1
    filters = [{

      filters = [{

        key      = "string"
        operator = "string"
        value    = "string"
      }]
      key              = "string"
      logical_operator = "string"
      operator         = "string"
      value            = "string"
    }]
    start_time = 1
  }]
}

output "catalystcenter_assurance_issues_query_example" {
  value = catalystcenter_assurance_issues_query.example
}