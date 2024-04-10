# README.md for Builder Design Pattern in Go

## Overview

This Go module implements the Builder design pattern to construct complex objects step by step. The pattern is particularly useful when an object must be created in various configurations, or when the construction process involves several steps that can be executed in different orders or configurations.

The provided code demonstrates the construction of network port configurations using a `Director` to control the building process and a `Builder` interface implemented by concrete builders for different types of port configurations.

## Structure

- **IPAddress**: A struct that represents an IP address, containing the address itself and a boolean indicating if it's a primary address.
- **Director**: Orchestrates the construction process, controlling which steps are executed to assemble the final product.
- **Builder Interface**: Defines the steps needed to build a network component. Concrete builders implement this interface.
- **Speed**: An enumeration representing the speed of the network port.
- **Component**: Represents a network component, holding properties such as port, IP, description, and speed.
- **PointToPointBuilder**: A concrete builder implementing the `Builder` interface to create a point-to-point network component.

## Usage

### Creating a Network Component

To create a network component, instantiate a `Director` and a concrete builder (e.g., `PointToPointBuilder`). Use the director to set the builder and define the construction steps:

```go
director := Director{}
portBuilder := PointToPointBuilder{}

director.SetBuilder(&portBuilder)
director.L3PortConstruct("Gi0/0/1", "192.168.1.1", true, "L3 PORT", Ethernet)
```

After executing the necessary steps, retrieve the final product:

```go
component := portBuilder.Build()
```

### Testing

The included test (`TestBuilder`) demonstrates how to construct two different port configurations and verify their properties.

## Explanation of Functions

- **SetBuilder**: Associates a specific builder with the director.
- **L3PortConstruct**: Constructs a Layer 3 network port with specified properties.
- **L2PortConstruct**: Constructs a Layer 2 network port without an IP address.
- **Build**: Finalizes the construction and returns the resulting network component.

## Test

Run the test using the standard Go testing tool:

```sh
go test
```

The test ensures that the components are constructed correctly with the specified attributes.

## Conclusion

This module provides a flexible and structured approach to constructing complex objects in Go, demonstrating the Builder design pattern's effectiveness in managing intricate object creation scenarios.
