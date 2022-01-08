terraform {
  required_providers {
    os2mo = {
      source = "github.com/Skeen/os2mo"
    }
  }
}

provider "os2mo" {
  url = "http://localhost:5000"
}

resource "os2mo_organisation" "root" {
    name = "Wowzers"
}
