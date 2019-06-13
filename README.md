# Sample Micro Tests

This repo contains some super simple examples how to test services hosted using micro framework.

In general you have to take care of three points:

 * Start and stop the (referenced) referenced service in the test
 * Disable the service.Init() method while service startup
 * Choose a random port for the service so duplicate usage of ports is prevented.
