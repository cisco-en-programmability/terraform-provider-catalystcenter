
resource "catalystcenter_connection_modesetting" "example" {
    provider = catalystcenter

    parameters {

      connection_mode = "string"
      parameters {

        client_id = "string"
        client_secret = "string"
        on_premise_host = "string"
        smart_account_name = "string"
      }
    }
}

output "catalystcenter_connection_modesetting_example" {
    value = catalystcenter_connection_modesetting.example
}
