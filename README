## Wireless and Mobile Networks Project 2


Isaac Coggin, Tyler Pittman, and Mitchell Price

### Project Setup

To set up the project for use, you must execute `make setup`. This command
will download and install a local copy of the go binary into the repository,
along with all project dependencies.

### Project Compilation

To compile the project, simply run `make`. This should create 4 binaries in the
`bin` directory.

### Running the Project

The following are the four binaries produced by the project:

1. __Data Source__

The data source binary is used to simulate a data sensor, generating data from
one of four types of sensors and sending it to the sensor hub. It is run by
calling `bin/data_source <sensor name> <generate port> <hub port>` where:

  * <sensor name> Denotes the name of the sensor to generate data from. This
  must be one of "heartrate", "location", "oxygen", or "gas"

  * <generate port> Indicates the port number on which the data should be
  generated.

  * <hub port> Indicates the port to which the data should be sent.

2. __Display Hub__

The display hub binary is used to display data for the fire chief once it is
received from the first responder. It is run by calling
`bin/display_hub <config file> <error file> <hostname> <display port>` where:

  * <config file> Denotes the path to the configuration file, which must follow
  the format described below, relative to the project root.

  * <error file> Denotes the path to the error file, which must follow the
  format described below, relative to the project root.

  * <hostname> Denotes the hostname of this node. This value should be "Display"

  * <display port> Indicates the port on which a webpage containing the display
  should be served. This value should be 10109.

3. __Manet Node__

The manet node binary is used to run a single node in the MANET. It is run by
calling `bin/manet_node` <config file> <hostname>` where:

  * <config file> Denotes the path to the configuration file, which must follow
  the format described below, relative to the project root.

  * <hostname> Is the hostname for this manet node, which must appear in the
  configuration file.

4. __Sensor Hub__

The sensor hub binary serves as the sensor hub on the first responder. It is run
by calling `bin/sensor_hub <config file> <error file> <hostname> <listenport>` where:

  * <config file> Denotes the path to the configuration file, which must follow
  the format described below, relative to the project root.

  * <error file> Denotes the path to the error file, which must follow the
  format described below, relative to the project root.

  * <hostname> Denotes the hostname of this node. This value should be "Sensor"

  * <listenport> Indicates on what port this binary will listen to sensor data.

Note that the Display Hub and Manet Nodes should be run before the sensor hub,
and that the sensor hub should be run before the data sources.

Once the network is running properly, you can visit `localhost:10109/home` to
see the project in action.

### Config File Format

The config file should consist of a complete description of every device in
the network, with each device having one line formatted the following way:

<manet address> <manet hostname> <host:port> <x> <y> <neighbor>...

Where:
  * <manet address> is the 16-bit address that uniquely identifies this node.

  * <manet hostname> is the hostname of this node to the network. Note that
  there should always be a hostname of "Sensor" and "Display", the sensor and
  display hubs respectively.

  * <host:port> indicates the hostname and port that the node will actually be
  run at. For example, this could be `tux065.eng.auburn.edu:10102`.

  * <x> is the X coordinate of the node.

  * <y> is the Y coordinate of the node.

  * <neighbor>... is a space separated list of the manet addresses of all
  nodes that this node is neighbors with.

### Error File Format

The error file should consist of a single line containing a single number, the
error rate of a 1-way transmission in the network. For example, the number
`0.40` corresponds to a 40% chance a given packet will be dropped by the
gremlin function.

### Convenient Startup and Teardown

A test script that sets up the entire network with 1 manet node on the given
device can be launched by running `make start`. To kill the network,
run `make kill`.
