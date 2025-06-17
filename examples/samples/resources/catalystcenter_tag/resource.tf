terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

provider "catalystcenter" {
  debug = "true"
}

resource "catalystcenter_tag" "example" {
  provider = catalystcenter
  parameters {

    description = "test33"
    dynamic_rules {

      member_type = "networkdevice"
      rules {

        # items     = "string"
        operation = "OR"
        items {
          name      = "hostname"
          operation = "ILIKE"
          value     = "%tf-rules%"
        }
        items {
          name      = "hostname"
          operation = "ILIKE"
          value     = "%-border-%"
        }
      }
      # rules {

      #   # items     = ["string"]
      #   name      = "string"
      #   operation = "string"
      #   value     = "string"
      #   values    = ["string"]
      # }
    }

    # instance_tenant_id = "62fe613eaa85d640eb70561e"
    //id                 = "0c3c64ea-69de-4686-8114-3f3dd7d97882" //Just necesesary for update...
    name = "api-test"
    //system_tag         = "false"
  }
}

output "catalystcenter_tag_example" {
  value = catalystcenter_tag.example
}
