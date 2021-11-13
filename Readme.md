# AnchAmber Lab

Software to manage the fish in a laboratory. It is side project with no real use. 
Therefore, technology and design decisions will be made just to try  them out and get some experience.

The Service is developed with DDD in mind and event sourcing. The services will expose an gRPC API to the outside. 
A frontend application will be developed with Flutter.

## Domain Model

There are different roles which may have different views for the entities.
- scientist
- BTA
- fish guys

### Line
A line is a group of fish which have the same base DNA reference.
It can be active or inactive. If it is active at least one fish exists.
Inactive lines can be reactivated.

#### Properties
- **Name** - Name of the line.
- **DNA String** - reference to the DNA string of the line
- **Active** - Flag to show if there are currently living fish of this line

### Fish
A representation of a fish. A fish has a reference to which line it belongs.
It also has a reference in which tank the fish lives.

#### Properties
- **Name** - Name of the fish.
- **Kind** - The kind of the fish.
- **Tank** - Reference to the tank in which the fish is hold.
- **Birthday** - Birthday of the fish.
- **Gender** - Gender of the fish.
- **Father** - Reference to the father of the fish.
- **Mother** - Reference to the mother of the fish.
- **Generation** - Count of the generation
- **Alive** - Flag to show if the fish is alive

### System
A system is module which holds multiple tanks.

#### Properties
- **Manufacturer** - Company which produces the system.
- **System Name** - Product name of the system.
- **Name** - A name to reference the system in the lab. This name is independent of the system name.
- **Responsible Person** - A reference to the person who is responsible to maintain the system.
- **Cleaning Interval** - The interval in days in which the system needs to be cleaned.
- **Location** - A reference to the room in which the system stands.
- **Last Cleaned** - Date when the system was cleaned the last time.

### Room

#### Properties
- **Building** - Building in which the room is located.
- **Level** - Level on which the room is located.
- **Name** - Name of the room. 
- **Open** - Is the room currently open and can be used for systems.
- **Features** - Features the room has, e.g. cold room, water supply.

### Person

#### Properties
- **First Name**
- **Last Name**
- **Email**

### Tank
A tank is a container in which the fish are kept. It can vary in size.

#### Properties
- **Size** - The size of the tank in liter.
- **In Use** - Flag to show if the tank is currently used.


## Services

### Systems Service



