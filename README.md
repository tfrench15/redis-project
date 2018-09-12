# Redis-Proxy

Redis-Proxy provides an HTTP interface to a lightweight, read-through LRU cache for a single backing Redis instance.

### Overview

The proxy listens for incoming HTTP requests on a configurable port, and maps HTTP GET requests to Redis GET commands using the base of the URL path as the key.  Non-GET requests return a Bad Request (400) HTTP error.

For example, a GET request issued to "http://localhost:8000/hi" parses 'hi' as the key to look up in the cache, or to fetch from Redis if it is not yet cached.  Keys not in Redis return a Status Not Found (404) HTTP error.

### Features

#### Configuration
The proxy is configurable via command-line flags.  You may customize:
1. `redisAddr`: The address of the backing Redis
2. `proxyAddr`: The port the proxy will listen on
3. `capacity`: The size of the cache
4. `expiry`: The time to live of a key in cache

#### How the Code works

- `redis.go` provides a function to initialize a connection to the backing Redis.
- `cache.go` defines the LRU cache struct and provides methods to retrieve keys from the cache.
- `proxy.go` composes the backing Redis and cache into the `RedisProxy` type, which implements the core service.
- `http.go` provides methods for handling HTTP requests to the service and satisfies the http.Handler interface.
- `main.go` defines command-line flags with which a user can configure their own Redis, HTTP port, and capacity / ttl of the cache.

#### Tests

The proxy comes with unit tests, leveraging Go's testing framework in `cache_test.go`, `http_test.go`.  One thing I need to do is write black-box tests and verify the tests cover the entire package's functionality.

#### Algorithmic Complexity

The cache, imported from Hashicorp's `simplelru` library, is implemented with a map. Looking up a key has linear time complexity (O(log n)).  Adding a key to the cache has constant time complexity (O(1)).

#### Instructions (TODO)

To test the cache, enter the top-level directory and run `make test`.  Note that I'm currently re-writing the Makefile, so this will be added shortly :-).