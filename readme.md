# Appmonitor checks

This repository contains checks for the [Appmonitor Go Client](https://github.com/sleepycrew/appmonitor-client) that don't fit into the repository.

Checks implemented here target very specific systems and software, for instance systemd.

I've experimented with golang's plugin system and dynamic linking a bit here, which is still in use for the integration tests using shunit2. But I would not recommend using it for anything else and it might be removed in the future.

## Tests
Integration tests are written using shunit2, and run in lxd containers providing full linux environments. Should you not be able to run lxd you can also use its client cli and connect to a remote instance before running the tests.