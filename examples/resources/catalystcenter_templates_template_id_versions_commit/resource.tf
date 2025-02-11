
resource "catalystcenter_templates_template_id_versions_commit" "example" {
    provider = meraki
    template_id = "string"
    parameters = [{
      
      commit_note = "string"
    }]
}

output "catalystcenter_templates_template_id_versions_commit_example" {
    value = catalystcenter_templates_template_id_versions_commit.example
}