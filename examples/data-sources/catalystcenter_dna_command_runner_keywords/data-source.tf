
data "catalystcenter_dna_command_runner_keywords" "example" {
  provider = catalystcenter
}

output "catalystcenter_dna_command_runner_keywords_example" {
  value = data.catalystcenter_dna_command_runner_keywords.example.items
}
