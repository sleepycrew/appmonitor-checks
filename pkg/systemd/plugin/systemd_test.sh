#! /bin/sh

testService() {
  docker build pkg/systemd/plugin -t amc-sysd
  container_name=$(docker run -d amc-sysd)
  docker cp ./ $container_name:/root/checks
  docker exec $container_name systemctl start nginx
  docker exec $container_name cd /root/checks && make 
  docker exec $container_name go run /root/checks/cmd/run-checks 
}

# Load shUnit2.
. third_party/shunit2/shunit2
