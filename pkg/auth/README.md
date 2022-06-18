Domain
Ports (Interfaces) - just an interface declarations
Repo - It's An output port (driven port) is another type of interface that is used by the application core to reach things outside of itself (like getting some data from a database).
Service - An input port (driving port) lets the application core to expose the functionality to the outside of the world.


Service can be injected to
 > HTTP Handlers
 > GRPC
 > Command Line