import argparse
import sys
import syslog
import docker

import collect_agent


def build_parser():
    parser = argparse.ArgumentParser(description='Docker container runner')
    parser.add_argument('docker_image', type=str, help='Name of docker image to run')
    parser.add_argument('command', type=str, help='Command to run on the container')
    parser.add_argument('-n', '--name', type=str, help='Container\'s name', default="")

    return parser


def main(image, command, name):
    """Program to run docker container"""
    try:
        client = docker.from_env()

        if not name:
            client.containers.run(
                image=image,
                command=command,
                detach=True
            )
        else:
            client.containers.run(
                image=image,
                command=command,
                detach=True,
                name=name
            )

    except Exception as ex:
        message = 'Error while executing container {} : {}'.format(image, ex)
        collect_agent.send_log(syslog.LOG_ERR, message)
        sys.exit(message)


if __name__ == '__main__':
    with collect_agent.use_configuration("/opt/openbach/agent/jobs/container_run/container_run_rstat_filter.conf"):
        args = build_parser().parse_args()
        main(args.docker_image, args.command, args.name)
