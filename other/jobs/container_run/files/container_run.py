import argparse
import sys
import syslog
import docker

import collect_agent


def build_parser():
    parser = argparse.ArgumentParser(description='Docker container runner')
    parser.add_argument('docker_image', type=str, help='Name of docker image to run')
    parser.add_argument('-c', '--command', type=str, help='Command to run on the container', default="")
    parser.add_argument('-n', '--name', type=str, help='Container\'s name', default="")
    parser.add_argument('-p', '--ports', type=str, help='List of comma separated ports to bind', default="")

    return parser


def main(image, command, name, ports):
    """Program to run docker container"""
    try:
        client = docker.from_env()
        list_port = {}

        if not command:
            command = None

        if ports:
            ports = ports.split(',')

            for port in ports:
                parts = port.split(':')

                if len(parts) == 2:
                    key = f'{parts[1]}/tcp'
                    value = int(parts[0])
                    list_port[key] = value

        if not name:
            client.containers.run(
                image=image,
                command=command,
                detach=True,
                ports=list_port
            )
        else:
            client.containers.run(
                image=image,
                command=command,
                detach=True,
                name=name,
                ports=list_port
            )

    except Exception as ex:
        message = 'Error while executing container {} : {}'.format(image, ex)
        collect_agent.send_log(syslog.LOG_ERR, message)
        sys.exit(message)


if __name__ == '__main__':
    with collect_agent.use_configuration("/opt/openbach/agent/jobs/container_run/container_run_rstat_filter.conf"):
        args = build_parser().parse_args()
        main(args.docker_image, args.command, args.name, args.ports)
