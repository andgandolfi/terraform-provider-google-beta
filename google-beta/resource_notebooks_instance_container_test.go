package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccNotebooksInstance_create_container(t *testing.T) {
	t.Parallel()

	prefix := fmt.Sprintf("%d", randInt(t))
	name := fmt.Sprintf("tf-%s", prefix)

	vcrTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNotebooksInstance_create_container(name),
			},
			{
				ResourceName:            "google_notebooks_instance.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ExpectNonEmptyPlan:      true,
				ImportStateVerifyIgnore: []string{"container_image", "metadata", "vm_image"},
			},
		},
	})
}

func testAccNotebooksInstance_create_container(name string) string {
	return fmt.Sprintf(`

resource "google_notebooks_instance" "test" {
  name = "%s"
  location = "us-west1-a"
  machine_type = "n1-standard-1"
  metadata = {
    proxy-mode = "service_account"
    terraform  = "true"
  }
  container_image {
    repository = "gcr.io/deeplearning-platform-release/base-cpu"
    tag = "latest"
  }
}
`, name)
}
