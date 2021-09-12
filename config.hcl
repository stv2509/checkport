ConnectTimeout = 2
RepeatTimeout = 60

host "google.com" {
  proto = "udp"
  ports = ["53","68", 79]
}

host "example.com" {
  proto = "udp"
  ports = ["53","68", 6579]
}

host "ya.ru" {
  proto = "tcp"
  ports = [443,80, 8080]
}
