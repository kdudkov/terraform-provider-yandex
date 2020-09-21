package yandex

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-multierror"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/vpc/v1"
)

func init() {
	resource.AddTestSweepers("yandex_vpc_security_group", &resource.Sweeper{
		Name: "yandex_vpc_security_group",
		F:    testSweepVPCSecurityGroups,
		Dependencies: []string{
			"yandex_compute_instance",
			"yandex_compute_instance_group",
			"yandex_dataproc_cluster",
			"yandex_kubernetes_node_group",
			"yandex_kubernetes_cluster",
			"yandex_mdb_clickhouse_cluster",
			"yandex_mdb_mongodb_cluster",
			"yandex_mdb_mysql_cluster",
			"yandex_mdb_postgresql_cluster",
			"yandex_mdb_redis_cluster",
		},
	})
}

func testSweepVPCSecurityGroups(_ string) error {
	conf, err := configForSweepers()
	if err != nil {
		return fmt.Errorf("error getting client: %s", err)
	}

	it := conf.sdk.VPC().SecurityGroup().SecurityGroupIterator(conf.Context(), conf.FolderID)
	result := &multierror.Error{}
	for it.Next() {
		id := it.Value().GetId()
		if !sweepVPCSecurityGroup(conf, id) {
			result = multierror.Append(result, fmt.Errorf("failed to sweep VPC security group %q", it.Value().GetId()))
		}
	}

	return result.ErrorOrNil()
}

func sweepVPCSecurityGroup(conf *Config, id string) bool {
	return sweepWithRetry(sweepVPCSecurityGroupOnce, conf, "VPC Security Group", id)
}

func sweepVPCSecurityGroupOnce(conf *Config, id string) error {
	ctx, cancel := conf.ContextWithTimeout(yandexVPCNetworkDefaultTimeout)
	defer cancel()

	op, err := conf.sdk.VPC().SecurityGroup().Delete(ctx, &vpc.DeleteSecurityGroupRequest{
		SecurityGroupId: id,
	})
	return handleSweepOperation(ctx, conf, op, err)
}

func TestAccVPCSecurityGroup_basic(t *testing.T) {
	t.Parallel()

	var securityGroup vpc.SecurityGroup

	networkName := fmt.Sprintf("tf-test-%s", acctest.RandString(10))
	sg1Name := fmt.Sprintf("tf-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVPCSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVPCSecurityGroupBasic(networkName, sg1Name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVPCSecurityGroupExists("yandex_vpc_security_group.sgr1", &securityGroup),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.#", "1"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "egress.#", "1"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.3168776168.protocol", "TCP"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.3168776168.port", "8080"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "egress.740716492.protocol", "ANY"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "egress.740716492.port", "-1"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "egress.740716492.from_port", "8090"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "egress.740716492.to_port", "8099"),
					testAccCheckCreatedAtAttr("yandex_vpc_security_group.sgr1"),
				),
			},
			{
				ResourceName:      "yandex_vpc_security_group.sgr1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCSecurityGroup_update(t *testing.T) {
	t.Parallel()

	var securityGroup vpc.SecurityGroup
	var securityGroup2 vpc.SecurityGroup

	networkName := fmt.Sprintf("tf-test-%s", acctest.RandString(10))
	sg1Name := fmt.Sprintf("tf-test-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVPCSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVPCSecurityGroupBasic(networkName, sg1Name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVPCSecurityGroupExists("yandex_vpc_security_group.sgr1", &securityGroup),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.#", "1"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "egress.#", "1"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.3168776168.protocol", "TCP"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.3168776168.port", "8080"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "egress.740716492.protocol", "ANY"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "egress.740716492.port", "-1"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "egress.740716492.from_port", "8090"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "egress.740716492.to_port", "8099"),
					testAccCheckCreatedAtAttr("yandex_vpc_security_group.sgr1"),

					testAccCheckVPCSecurityGroupExists("yandex_vpc_security_group.sgr2", &securityGroup2),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr2", "ingress.#", "1"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr2", "egress.#", "1"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr2", "egress.669457800.protocol", "ANY"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr2", "egress.669457800.port", "9000"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr2", "egress.669457800.predefined_target", "self_security_group"),
					// It's hard for test rule with security_group_id because of not stable hash of rule with ID.
					// predefined_target has the same logic. Assume that test covers this situation.
					testAccCheckCreatedAtAttr("yandex_vpc_security_group.sgr2"),
				),
			},
			{
				Config: testAccVPCSecurityGroupBasic2(networkName, sg1Name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVPCSecurityGroupExists("yandex_vpc_security_group.sgr1", &securityGroup),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.#", "2"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.3168776168.protocol", "TCP"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.3168776168.port", "8080"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.583108381.protocol", "ANY"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.583108381.port", "-1"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.583108381.from_port", "8091"),
					resource.TestCheckResourceAttr("yandex_vpc_security_group.sgr1", "ingress.583108381.to_port", "8099"),
					testAccCheckCreatedAtAttr("yandex_vpc_security_group.sgr1"),
				),
			},
			{
				ResourceName:      "yandex_vpc_security_group.sgr1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckVPCSecurityGroupExists(name string, securityGroup *vpc.SecurityGroup) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("not found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		sdk := testAccProvider.Meta().(*Config).sdk
		found, err := sdk.VPC().SecurityGroup().Get(context.Background(), &vpc.GetSecurityGroupRequest{
			SecurityGroupId: rs.Primary.ID,
		})
		if err != nil {
			return err
		}

		if found.Id != rs.Primary.ID {
			return fmt.Errorf("security group not found")
		}

		securityGroup = found

		return nil
	}
}

func testAccVPCSecurityGroupBasic(networkName, sgr1Name string) string {
	return fmt.Sprintf(`
resource "yandex_vpc_network" "foo" {
  name = "%s"
}

resource "yandex_vpc_security_group" "sgr1" {
  name        = "%s"
  description = "description for security group"
  network_id  = "${yandex_vpc_network.foo.id}"
  folder_id   = "%s"

  labels = {
    tf-label    = "tf-label-value-a"
    empty-label = ""
  }

  ingress {
    description    = "rule1 description"
    protocol       = "TCP"
    v4_cidr_blocks = ["10.0.1.0/24", "10.0.2.0/24"]
    port           = 8080
  }

  egress {
    description    = "rule2 description"
    protocol       = "ANY"
    v4_cidr_blocks = ["10.0.1.0/24", "10.0.2.0/24"]
    from_port      = 8090
    to_port        = 8099
  }
}

resource "yandex_vpc_security_group" "sgr2" {
  network_id  = "${yandex_vpc_network.foo.id}"
  folder_id   = "%s"

  egress {
    description       = "rule3 description"
    protocol          = "ANY"
    predefined_target = "self_security_group"
    port              = 9000
  }

  ingress {
    description       = "rule4 description"
    protocol          = "TCP"
    security_group_id = "${yandex_vpc_security_group.sgr1.id}"
    port              = 9010
  }
}

`, networkName, sgr1Name, getExampleFolderID(), getExampleFolderID())
}

func testAccVPCSecurityGroupBasic2(networkName, sgr1Name string) string {
	return fmt.Sprintf(`
resource "yandex_vpc_network" "foo" {
  name = "%s"
}

resource "yandex_vpc_security_group" "sgr1" {
  name        = "%s"
  description = "description for security group"
  network_id  = "${yandex_vpc_network.foo.id}"
  folder_id   = "%s"

  labels = {
    tf-label    = "tf-label-value-a"
    empty-label = ""
  }

  ingress {
    description    = "rule1 description"
    protocol       = "TCP"
    v4_cidr_blocks = ["10.0.1.0/24", "10.0.2.0/24"]
    port           = 8080
  }

  ingress {
    description    = "rule2 description2"
    protocol       = "ANY"
    v4_cidr_blocks = ["10.0.1.0/24", "10.0.2.0/24"]
    from_port      = 8091
    to_port        = 8099
  }
}

`, networkName, sgr1Name, getExampleFolderID())
}

func testAccCheckVPCSecurityGroupDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "yandex_vpc_security_group" {
			continue
		}

		_, err := config.sdk.VPC().SecurityGroup().Get(context.Background(), &vpc.GetSecurityGroupRequest{
			SecurityGroupId: rs.Primary.ID,
		})
		if err == nil {
			return fmt.Errorf("Security group still exists")
		}
	}

	return nil
}
