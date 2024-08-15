# This Kurtosis package spins up a defund rollup that connects to a DA node

# Import the local da kurtosis package
da_node = import_module("github.com/rollkit/local-da/main.star@v0.3.0")


def run(plan):
    # Start the DA node
    da_address = da_node.run(
        plan,
    )
    plan.print("connecting to da layer via {0}".format(da_address))

    # Define the defund start command
    defund_start_cmd = [
        "rollkit",
        "start",
        "--rollkit.aggregator",
        "--rollkit.da_address {0}".format(da_address),
        "--rollkit.block_time 0.1ms",
        "--minimum-gas-prices 0.01stake",
    ]
    # Define the jsonrpc ports
    defund_ports = {
        "jsonrpc": PortSpec(
            number=26657, transport_protocol="TCP", application_protocol="http"
        ),
    }
    # Start the defund chain
    defund = plan.add_service(
        name="defund",
        config=ServiceConfig(
            # Use the defund image we just built
            image="mysticlabss/defund:v0.1.0",
            # Set the command to start the defund chain in the docker container
            cmd=["/bin/sh", "-c", " ".join(defund_start_cmd)],
            ports=defund_ports,
            public_ports=defund_ports,
        ),
    )