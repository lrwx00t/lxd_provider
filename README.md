# LXD Provider for Terraform

LXD provider for terraform to support managing LXC containers in Linux environment.

The provider currently supports managing the status (`Running` and `Stopped`) of the LXC containers using a Unix Domain Sockets.

## Installation
The provider is only available to be published from a local installation as it has not been published in any registry yet. However, the `makefile` provided in the repo should make the installtion step fairly easy and quick.

To install the the LXD provider, you need to run the `install-provider` in the `makefile`:

```bash
make install-provider
 ```

The task will build the binary and install the plugin in the `.terraformd/plugins` directory. 

## Testing the provider

In order to test the provider, you need to make sure that you have a running instance of LXD in your local machine.

The repo provides a terraform project to create an LXD resource to `start/stop` the container. It can be found in the `tests/lxd_provider_tf` directory.

The `makefile` task can be used to run the testing. Please ensure that it matches the container name in your local machine.

```bash
make test-provider-apply
```

You can also run both tasks at the same time to ensure that any develpoment on the provider can be tested immediately. Please note that this step will remove any terraform state file and will also remove the provider and re-install it.

```bash
make install-provider && make test-provider-apply
```

## Running Unit Tests

There is a couple of unit test cases provided in this project which can be used by triggering the `makfile` task:

```bash
make test
```