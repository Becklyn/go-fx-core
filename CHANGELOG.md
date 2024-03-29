# v2.0.2

- (bug) Fix fiber middleware definition check

# v2.0.1

- (bug) Fix module version

# v2.0.0

- (bc) Add `context.Context` as first param in all cqrs handler intarfaces
- (bc) Replace `MiddlewareRegistry` by `MiddlewareHandlerMap`
- (bc) Move `cors.Module` to `middleware.NoCors`
- (internal) Update all dependencies

# v1.1.1

- (bug) Fix module declared as program

# v1.1.0

- (feature) Add injectable middleware for fiber
- (feature) Add nocors module for development without cors

# v1.0.1

- (bug) Set correct content type when returning error messages in fiber

# v1.0.0

- (feature) v1 release

# v1.0.0-alpha10

- (feature) Add error handling middleware to fiber module

# v1.0.0-alpha9

- (bug) Fix bool env variable always returning default value if set to false
- (improvement) Add keepalive rules to the grpc server

# v1.0.0-alpha8

- (feature) Add readyness module

# v1.0.0-alpha7

- (internal) Fix metrics middleware
- (bc) Remove fiber middleware registry

# v1.0.0-alpha6

- (improvement) Add metrics module to core module
- (internal) Move fiber module to web package

# v1.0.0-alpha5

- (internal) Fix env module initialization order
- (feature) Add health module
- (bug) Allow debug log level configuration
- (feature) Add prometheus metrics module
- (feature) Add fiber middleware registry

# v1.0.0-alpha4

- (improvement) Improve gRPC address logging

# v1.0.0-alpha3

- (improvement) Improve gRPC listen addr parsing

# v1.0.0-alpha2

- (feature) Add gRPC server module
- (improvement) Env variable functions now have methods with default values
- (improvement) Int env variable function now returns

# v1.0.0-alpha1

- (internal) Add code style tests
- (improvement) Support .env
- (feature) Add cqrs handler interfaces
- (feature) GraphQL module
- (feature) Core module

# v1.0.0-alpha

- (feature) Fiber module
- (feature) Logrus module
- (feature) Validator module
