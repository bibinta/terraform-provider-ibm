// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package mqcloud_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmMqcloudCapacityDetailsDataSourceBasic(t *testing.T) {
	service_instance_guid := acc.MqcloudDeploymentID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckMqcloud(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmMqcloudCapacityDetailsDataSourceConfigBasic(service_instance_guid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_capacity_details.mqcloud_capacity_details_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_capacity_details.mqcloud_capacity_details_instance", "service_instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_capacity_details.mqcloud_capacity_details_instance", "capacity_units"),
					resource.TestCheckResourceAttrSet("data.ibm_mqcloud_capacity_details.mqcloud_capacity_details_instance", "capacity_status"),
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
