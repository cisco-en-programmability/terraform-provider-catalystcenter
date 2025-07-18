terraform {
  required_providers {
    catalystcenter = {
      version = "1.2.0-beta"
      source  = "hashicorp.com/edu/catalystcenter"
      # "hashicorp.com/edu/catalystcenter" is the local built source, change to "cisco-en-programmability/catalystcenter" to use downloaded version from registry
    }
  }
}

resource "catalystcenter_event_subscription_syslog" "example" {
  provider = catalystcenter
  parameters {
    payload {
      description = "Test Terraform 4"
      filter {

        # categories = ["string"]
        /*domains_subdomains {

        domain      = "string"
        sub_domains = ["string"]
      }*/
        event_ids = ["LICMGMT-DEV-REG-SUCCESS", "LICMGMT-DEV-REG-FAILURE", "LICMGMT-DEV-DEREG-SUCCESS"]
        #severities = ["string"]
        #site_ids   = ["string"]
        #sources    = ["string"]
        #types      = ["string"]
      }
      name = "Test Terraform 2"
      subscription_endpoints {

        instance_id = "cc5f19b2-baac-4835-8285-8b9e91c15cc4"
        subscription_details {
          connector_type = "SYSLOG"
        }
      }
      #subscription_id = "string"
      #version         = "string"
    }
  }
}

