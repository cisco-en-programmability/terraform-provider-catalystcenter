
data "catalystcenter_pnp_device_count" "example" {
    provider = catalystcenter
    last_contact = "false"
    name = ["string"]
    onb_state = ["string"]
    pid = ["string"]
    serial_number = ["string"]
    smart_account_id = ["string"]
    source = ["string"]
    state = ["string"]
    virtual_account_id = ["string"]
    workflow_id = ["string"]
    workflow_name = ["string"]
}

output "catalystcenter_pnp_device_count_example" {
    value = data.catalystcenter_pnp_device_count.example.item
}
