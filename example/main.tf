terraform {
  required_providers {
    abiraj = {
      source = "abirajvp/abiraj"
      version = "0.1.1"
    }
  }
}

provider abiraj{
    username = "abirajvp"
    password = "http://localhost:8000"
}

