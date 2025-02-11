
data "catalystcenter_pnp_device" "example" {
    provider = catalystcenter
    hostname = "string"
    last_contact = "false"
    limit = 1
    mac_address = "string"
    name = ["string"]
    offset = 1
    onb_state = ["string"]
    pid = ["string"]
    serial_number = ["string"]
    site_name = "string"
    smart_account_id = ["string"]
    sort = ["string"]
    sort_order = "string"
    source = ["string"]
    state = ["string"]
    virtual_account_id = ["string"]
    workflow_id = ["string"]
    workflow_name = ["string"]
}

output "catalystcenter_pnp_device_example" {
    value = data.catalystcenter_pnp_device.example.items
}

data "catalystcenter_pnp_device" "example" {
    provider = catalystcenter
    id = "string"
}

output "catalystcenter_pnp_device_example" {
    value = data.catalystcenter_pnp_device.example.item
}
