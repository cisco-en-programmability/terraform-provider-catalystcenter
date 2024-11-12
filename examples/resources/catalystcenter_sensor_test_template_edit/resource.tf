
resource "catalystcenter_sensor_test_template_edit" "example" {
  provider = meraki
  parameters {

    id                 = "string"
    action_in_progress = "string"
    ap_coverage {

      bands                 = "string"
      number_of_aps_to_test = 1
      rssi_threshold        = 1
    }
    connection      = "string"
    encryption_mode = "string"
    frequency {

      unit  = "string"
      value = 1
    }
    last_modified_time = 1
    location           = "string"
    location_info_list {

      all_sensors            = "false"
      custom_management_vlan = "false"
      location_id            = "string"
      location_type          = "string"
      mac_address_list       = ["string"]
      management_vlan        = "string"
      site_hierarchy         = "string"
    }
    model_version            = 1
    name                     = "string"
    num_associated_sensor    = 1
    num_neighbor_apthreshold = 1
    profiles {

      auth_protocol           = "string"
      auth_type               = "string"
      certdownloadurl         = "string"
      certfilename            = "string"
      certpassphrase          = "string"
      certstatus              = "string"
      certxferprotocol        = "string"
      device_type             = "string"
      eap_method              = "string"
      ext_web_auth            = "false"
      ext_web_auth_access_url = "string"
      ext_web_auth_html_tag {

        label = "string"
        tag   = "string"
        value = "string"
      }
      ext_web_auth_portal     = "string"
      ext_web_auth_virtual_ip = "string"
      location_vlan_list {

        location_id = "string"
        vlans       = ["string"]
      }
      password      = "******"
      password_type = "******"
      profile_name  = "string"
      psk           = "string"
      qos_policy    = "string"
      scep          = "false"
      tests {

        config {

          direction        = "string"
          domains          = ["string"]
          downlink_test    = "false"
          end_port         = 1
          exit_command     = "string"
          final_prompt     = "string"
          ndt_server       = "string"
          ndt_server_path  = "string"
          ndt_server_port  = "string"
          num_packets      = 1
          password         = "******"
          password_prompt  = "******"
          path_to_download = "string"
          port             = 1
          probe_type       = "string"
          protocol         = "string"
          proxy_password   = "string"
          proxy_port       = "string"
          proxy_server     = "string"
          proxy_user_name  = "string"
          server           = "string"
          servers          = ["string"]
          shared_secret    = "string"
          start_port       = 1
          transfer_type    = "string"
          udp_bandwidth    = 1
          uplink_test      = "false"
          url              = "string"
          user_name        = "string"
          user_name_prompt = "string"
        }
        name = "string"
      }
      username   = "string"
      vlan       = "string"
      white_list = "false"
    }
    radio_as_sensor_removed = "false"
    rssi_threshold          = 1
    run_now                 = "string"
    schedule_in_days        = 1
    sensors {

      all_sensor_addition       = "false"
      assigned                  = "false"
      config_updated            = "string"
      host_name                 = "string"
      i_perf_info               = ["string"]
      id                        = "string"
      ip_address                = "string"
      location_id               = "string"
      mac_address               = "string"
      marked_for_uninstall      = "false"
      name                      = "string"
      run_now                   = "string"
      sensor_type               = "string"
      service_policy            = "string"
      status                    = "string"
      switch_mac                = "string"
      switch_serial_number      = "string"
      switch_uuid               = "string"
      target_a_ps               = ["string"]
      test_mac_addresses        = "------"
      wired_application_message = "string"
      wired_application_status  = "string"
      xor_sensor                = "false"
    }
    show_wlc_upgrade_banner = "false"
    site_hierarchy          = "string"
    ssids {

      auth_protocol           = "string"
      auth_type               = "string"
      bands                   = "string"
      certdownloadurl         = "string"
      certfilename            = "string"
      certpassphrase          = "string"
      certstatus              = "string"
      certxferprotocol        = "string"
      eap_method              = "string"
      ext_web_auth            = "false"
      ext_web_auth_access_url = "string"
      ext_web_auth_html_tag {

        label = "string"
        tag   = "string"
        value = "string"
      }
      ext_web_auth_portal          = "string"
      ext_web_auth_virtual_ip      = "string"
      id                           = 1
      layer3web_auth_email_address = "string"
      layer3web_authpassword       = "******"
      layer3web_authsecurity       = "string"
      layer3web_authuser_name      = "string"
      num_aps                      = 1
      num_sensors                  = 1
      password                     = "******"
      password_type                = "******"
      profile_name                 = "string"
      proxy_password               = "string"
      proxy_port                   = "string"
      proxy_server                 = "string"
      proxy_user_name              = "string"
      psk                          = "string"
      qos_policy                   = "string"
      scep                         = "false"
      ssid                         = "string"
      status                       = "string"
      tests {

        config {

          direction        = "string"
          domains          = ["string"]
          downlink_test    = "false"
          end_port         = 1
          exit_command     = "string"
          final_prompt     = "string"
          ndt_server       = "string"
          ndt_server_path  = "string"
          ndt_server_port  = "string"
          num_packets      = 1
          password         = "******"
          password_prompt  = "******"
          path_to_download = "string"
          port             = 1
          probe_type       = "string"
          protocol         = "string"
          proxy_password   = "string"
          proxy_port       = "string"
          proxy_server     = "string"
          proxy_user_name  = "string"
          server           = "string"
          servers          = ["string"]
          shared_secret    = "string"
          start_port       = 1
          transfer_type    = "string"
          udp_bandwidth    = 1
          uplink_test      = "false"
          url              = "string"
          user_name        = "string"
          user_name_prompt = "string"
        }
        name = "string"
      }
      third_party {

        selected = "false"
      }
      username   = "string"
      valid_from = 1
      valid_to   = 1
      white_list = "false"
      wlan_id    = 1
      wlc        = "string"
    }
    start_time         = 1
    status             = "string"
    template_name      = "string"
    test_schedule_mode = "string"
    version            = 1
    wlans              = ["string"]
  }
}

output "catalystcenter_sensor_test_template_edit_example" {
  value = catalystcenter_sensor_test_template_edit.example
}