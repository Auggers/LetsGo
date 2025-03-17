#!/bin/bash

OS=""
PREPARE_PATH=""
VAGRANT_FOLDER=""
PSMDB_FOLDER=""
REPO_TYPE=""
PSMDB_VERSION=""
PLAYBOOK_PATH=""

usage() {
    echo "Usage: $0 [OPTIONS]"
    echo "Options:"
    echo "  -os --operating-system, Specify the Operating System you want to install"
    echo "  -pp, --prepare_path FILE,  Specify a file path to prepare.yml file"
    echo "  -vp, --vagrantpath, Location of your vagrantfile"
    echo "  -pf, --psmdb_folder, Location of your psmdb folder e.g. psmdb-testing/psmdb"
    echo "  -rt, --repo_type, Repo type where PSMDB is being pulled from e.g. release, testing, experimental"
    echo "  -pv, --psmdb_version, Version of PSMDB e.g. 6.0.20, 8.0.0"
    echo "  -pbp, --playbook_path, Specify a file path to playbook.yml file"
    exit 0
}

# Parse command-line arguments in any order
while [[ "$#" -gt 0 ]]; do
    case "$1" in
        -h|--help)
            usage
            ;;
        -os|--operating-system)
            if [[ -n "$2" && "$2" != -* ]]; then
                OS="$2"
                shift 2
            else
                echo "Error: Missing argument for $1" >&2
                exit 1
            fi
            ;;
        -pp|--prepare-path)
            if [[ -n "$2" && "$2" != -* ]]; then
                PREPARE_PATH="$2"
                shift 2
            else
                echo "Error: Missing argument for $1" >&2
                exit 1
            fi
            ;;
          -vf|--vagrantfile)
            if [[ -n "$2" && "$2" != -* ]]; then
                VAGRANT_FOLDER="$2"
                shift 2
            else
                echo "Error: Missing argument for $1" >&2
                exit 1
            fi
            ;;
          -pf|--psmdb_folder)
            if [[ -n "$2" && "$2" != -* ]]; then
                PSMDB_FOLDER="$2"
                shift 2
            else
                echo "Error: Missing argument for $1" >&2
                exit 1
            fi
            ;;
          -rt|--repo_type)
            if [[ -n "$2" && "$2" != -* ]]; then
                REPO_TYPE="$2"
                shift 2
            else
                echo "Error: Missing argument for $1" >&2
                exit 1
            fi
            ;;
          -pv|--psmdb_version)
            if [[ -n "$2" && "$2" != -* ]]; then
                PSMDB_VERSION="$2"
                shift 2
            else
                echo "Error: Missing argument for $1" >&2
                exit 1
            fi
            ;;
          -pbp|--playbook-path)
            if [[ -n "$2" && "$2" != -* ]]; then
                PLAYBOOK_PATH="$2"
                shift 2
            else
                echo "Error: Missing argument for $1" >&2
                exit 1
            fi
            ;;
        *)
            echo "Error: Unknown option '$1'" >&2
            usage
            ;;
    esac
done

#Change directory to vagrant folder containing all necessary files
cd "$VAGRANT_FOLDER" || exit

# Create VM and bring the host up
vagrant up $OS || exit

# Ping to make sure that VM is online
PING=$(ansible all -m ping | grep -A 6 "$OS" | grep '"ping": "pong"')
if [[ $PING =~ "pong" ]]; then
  echo "$OS is up and running"
else
  echo "$OS is not online"
  exit 1
fi

#Run prepare playbook for the selected VM
ansible-playbook "$PREPARE_PATH" --limit "$OS" || exit
echo "Prepare was successful"

# Run "main" playbook for selected VM
ANSIBLE_ROLES_PATH=$PSMDB_FOLDER REPO=$REPO_TYPE PSMDB_VERSION=$PSMDB_VERSION ansible-playbook $PLAYBOOK_PATH --limit $OS || exit
echo "Playbook was successful"