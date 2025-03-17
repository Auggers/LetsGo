#!/usr/bin/env python3
import subprocess
import json
import re

def get_vagrant_machines():
    result = subprocess.run(["vagrant", "status"], capture_output=True, text=True)
    machines = []
    for line in result.stdout.splitlines():
        match = re.match(r"^([a-zA-Z0-9_-]+)\s+running", line)
        if match:
            machines.append(match.group(1))
    return machines

def get_vagrant_ssh_configs():
    inventory = {"all": {"hosts": []}, "_meta": {"hostvars": {}}}
    machines = get_vagrant_machines()

    for machine in machines:
        result = subprocess.run(["vagrant", "ssh-config", machine], capture_output=True, text=True)
        config_lines = result.stdout.splitlines()
        ssh_config = {}

        for line in config_lines:
            if "HostName" in line:
                ssh_config["ansible_host"] = line.split()[1]
            elif "User" in line:
                ssh_config["ansible_user"] = "vagrant"
            elif "Port" in line:
                ssh_config["ansible_port"] = line.split()[1]
            elif "IdentityFile" in line:
                ssh_config["ansible_ssh_private_key_file"] = line.split()[1]

        if ssh_config:
            inventory["all"]["hosts"].append(machine)
            inventory["_meta"]["hostvars"][machine] = ssh_config

    return inventory

if __name__ == "__main__":
    print(json.dumps(get_vagrant_ssh_configs(), indent=2))

