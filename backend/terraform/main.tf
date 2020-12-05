provider "random" {

}

resource "random_id" "server" {
  byte_length = 8
}

output "result" {
  value = random_id.server.hex
}
