terraform {
  required_providers {
    os2mo = {
      version = "0.2"
      source  = "hashicorp.com/edu/os2mo"
    }
  }
}

provider "os2mo" {
  username = "education"
  password = "test123"
}

resource "os2mo_order" "sample" {}

output "sample_order" {
  value = os2mo_order.sample
}
