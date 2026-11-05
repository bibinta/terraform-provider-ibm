// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmMqcloudCapacityDetailsDataSourceBasic(t *testing.T) {
	service_instance_guid := acc.MqcloudCapacityID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckMqcloud(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmMqcloudCapacityDetailsDataSourceConfigBasic(service_instance_guid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_capacity_details.mqcloud_capacity_details_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_capacity_details.mqcloud_capacity_details_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_capacity_details.mqcloud_capacity_details_instance", "capacity_units"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_capacity_details.mqcloud_capacity_details_instance", "vpc_entitlement"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_capacity_details.mqcloud_capacity_details_instance", "vpc_usage"),
				),
			},
		},
	})
}

func testAccCheckIbmMqcloudCapacityDetailsDataSourceConfigBasic(service_instance_guid string) string {
	return fmt.Sprintf(`
		data "ibm_mqcloud_capacity_details" "mqcloud_capacity_details_instance" {
			service_instance_guid = "%s"
		}
	`, service_instance_guid)
}
