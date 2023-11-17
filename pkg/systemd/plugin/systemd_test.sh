#! /bin/sh

UUID=$(uuidgen)
CONTAINER_NAME="amc-sysd-$UUID"

oneTimeSetUp() {
  # setup container
  lxc launch images:debian/bookworm $CONTAINER_NAME
  lxc exec $CONTAINER_NAME -- apt-get update > /dev/null
  lxc exec $CONTAINER_NAME -- apt-get install -y golang make ca-certificates mpd > /dev/null
  lxc exec $CONTAINER_NAME -- mkdir -p /root/amc > /dev/null
  lxc file push -r ./ $CONTAINER_NAME/root/amc > /dev/null
  lxc exec $CONTAINER_NAME --cwd /root/amc -- make clean > /dev/null
  lxc exec $CONTAINER_NAME --cwd /root/amc -- make > /dev/null
}

oneTimeTearDown() {
  lxc stop $CONTAINER_NAME
  lxc delete $CONTAINER_NAME
  echo "done" > /dev/null
}

testServiceRunning() {
  lxc exec $CONTAINER_NAME --cwd /root/amc --cwd /root/amc -- go run ./cmd/run-check /root/amc/build/systemd.so  '{"Name":"systemd-udevd.service", "Status": "running"}'
}


testServiceNotRunning() {
  lxc exec $CONTAINER_NAME --cwd /root/amc --cwd /root/amc -- go run ./cmd/run-check /root/amc/build/systemd.so  '{"Name":"mpd.service", "Status": "running"}'
  assertSame "1" "$?"
}

# Load shUnit2.
. third_party/shunit2/shunit2
