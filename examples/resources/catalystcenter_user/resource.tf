
resource "catalystcenter_user" "example" {
  provider = catalystcenter
 
  parameters {

    email      = "string"
    first_name = "string"
    last_name  = "string"
    password   = "******"
    role_list  = ["string"]
    user_id    = "string"
    username   = "string"
  }
}

output "catalystcenter_user_example" {
  value = catalystcenter_user.example
}