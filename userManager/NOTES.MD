The middleware pattern makes it easy for us to create our functions and add Logging, security and other standardization by just feeding in our business logic to the middlware functions, resulting in an entire stack of appropriate functions.

Authentication is handled by wrapping sensitive functions in an Auth function.

In a production scenario you would need the ability pass different permissions, so a configurable auth function would take in those permissions and check a pernission store.