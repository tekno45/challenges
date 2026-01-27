This proxy takes the config file and sets up a connection randomly from a set of backend targets.

Each pool of listeners has its own set of targets.

We guarantee a connection will be made by retrying with a recursive call to the dialer function without a timeout.

In production there would be a timeout and most likely a limit on retries on a speciic target.

Using a random number based on the length of the address we may retry an endpoint more than once, but is a simpe round robin since we retry until a connection is made and never timeout.