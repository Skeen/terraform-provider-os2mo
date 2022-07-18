terraform {
  required_providers {
    os2mo = {
      source = "github.com/Skeen/os2mo"
    }
  }
}

provider "os2mo" {
  url = "http://localhost:5000/graphql"
}

data "os2mo_itsystems" "all" {}

output "itoutput" {
  value = data.os2mo_itsystems.all
}

resource "os2mo_organisation" "root" {
  name = "Wowzers2"
}


