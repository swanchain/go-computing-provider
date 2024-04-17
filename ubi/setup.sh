#!/bin/bash
#Lagrange computer-provider initializes the ubi-task environment.
set -euxo pipefail
export DEBIAN_FRONTEND=noninteractive
sudo dpkg --set-selections <<< "cloud-init install" || true
# Detect OS type
OS="$(uname)"
    case $OS in
        "Linux")
            if [ -f /etc/os-release ]; then
                . /etc/os-release
                DISTRO=$ID
                VERSION=$VERSION_ID
            else
                echo "Your Linux distribution is not supported."
                exit 1
            fi
            ;;
    esac

NVIDIA_PRESENT=$(lspci | grep -i nvidia || true)

if [[ -z "$NVIDIA_PRESENT" ]]; then
    echo "No NVIDIA device detected on this system."
else
# Check if nvidia-smi is available and working
    if command -v nvidia-smi &>/dev/null; then
        echo "CUDA drivers already installed as nvidia-smi works."
    else

                # Depending on Distro
                case $DISTRO in
                    "ubuntu")
                        case $VERSION in
                            "20.04")
                                sudo -- sh -c 'apt-get update; apt-get upgrade -y; apt-get autoremove -y; apt-get autoclean -y'
                                sudo -- sh -c 'apt-get update; apt-get upgrade -y; apt-get autoremove -y; apt-get autoclean -y'
                                sudo apt install linux-headers-$(uname -r) -y
								sudo apt del 7fa2af80 || true
                                sudo apt remove 7fa2af80 || true
                                sudo apt install build-essential apt-transport-https ca-certificates cmake gpg unzip pkg-config software-properties-common ubuntu-drivers-common alsa-utils -y
                                sudo apt update -y
                                sudo dirmngr </dev/null
                                if sudo apt-add-repository -y ppa:graphics-drivers/ppa && sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys FCAE110B1118213C; then
                                    echo "Alternative method succeeded."
                                else
                                    echo "Alternative method failed. Trying the original method..."
                                    sudo dirmngr </dev/null
                                    sudo apt-add-repository -y ppa:graphics-drivers/ppa
                                    sudo gpg --no-default-keyring --keyring gnupg-ring:/etc/apt/trusted.gpg.d/graphics-drivers.gpg --keyserver keyserver.ubuntu.com --recv-keys FCAE110B1118213C
                                    sudo chmod 644 /etc/apt/trusted.gpg.d/graphics-drivers.gpg
                                fi
                                sudo ubuntu-drivers autoinstall
								sudo apt-mark hold nvidia* libnvidia*
                                ;;

                            "22.04")
                                sudo -- sh -c 'apt-get update; apt-get remove needrestart -y; apt-get upgrade -y; apt-get autoremove -y; apt-get autoclean -y'
                                sudo -- sh -c 'apt-get update; apt-get remove needrestart -y; apt-get upgrade -y; apt-get autoremove -y; apt-get autoclean -y'
                                sudo apt install linux-headers-$(uname -r) -y
                                sudo apt del 7fa2af80 || true
                                sudo apt remove 7fa2af80 || true
								sudo apt update -y
                                sudo apt install build-essential apt-transport-https ca-certificates cmake gpg unzip pkg-config software-properties-common ubuntu-drivers-common alsa-utils -y
                                sudo dirmngr </dev/null
                                if sudo apt-add-repository -y ppa:graphics-drivers/ppa && sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys FCAE110B1118213C; then
                                    echo "Alternative method succeeded."
                                else
                                    echo "Alternative method failed. Trying the original method..."
                                    sudo dirmngr </dev/null
                                    sudo apt-add-repository -y ppa:graphics-drivers/ppa
                                    sudo gpg --no-default-keyring --keyring gnupg-ring:/etc/apt/trusted.gpg.d/graphics-drivers.gpg --keyserver keyserver.ubuntu.com --recv-keys FCAE110B1118213C
                                    sudo chmod 644 /etc/apt/trusted.gpg.d/graphics-drivers.gpg
                                fi
                                sudo ubuntu-drivers autoinstall
								sudo apt-mark hold nvidia* libnvidia*
                                ;;

                            "18.04")
                                # Commands specific to Ubuntu 18.04
                                sudo -- sh -c 'apt-get update; apt-get upgrade -y; apt-get autoremove -y; apt-get autoclean -y'
                                sudo apt-get install linux-headers-$(uname -r) -y
                                sudo apt del 7fa2af80 || true
                                sudo apt remove 7fa2af80 || true
                                sudo apt install build-essential apt-transport-https cmake gpg unzip pkg-config software-properties-common ubuntu-drivers-common alsa-utils -y
                                sudo apt update -y
                                sudo ubuntu-drivers install
								sudo apt-mark hold nvidia* libnvidia*
                                ;;

                            *)
                                echo "This version of Ubuntu is not supported in this script."
                                exit 1
                                ;;
                        esac
                        ;;

                    "debian")
                        case $VERSION in
                            "10"|"11")
                                sudo -- sh -c 'apt update; apt upgrade -y; apt autoremove -y; apt autoclean -y'
                                sudo apt install linux-headers-$(uname -r) -y
                                sudo apt update -y
                                sudo apt install nvidia-driver firmware-misc-nonfree
                                ;;

                            *)
                                echo "This version of Debian is not supported in this script."
                                exit 1
                                ;;
                        esac
                        ;;

                    *)
                        echo "Your Linux distribution is not supported."
                        exit 1
                        ;;

            *)
                echo "Your OS is not supported."
                exit 1
                ;;
        esac
    fi
fi
# For testing purposes, this should output NVIDIA's driver version
if [[ ! -z "$NVIDIA_PRESENT" ]]; then
    nvidia-smi
fi

# Check if Docker is installed
if command -v docker &>/dev/null; then
    echo "Docker is already installed."
else
    echo "Docker is not installed. Proceeding with installations..."
    # Install Docker-ce keyring
    sudo apt update -y
    sudo apt install -y ca-certificates curl gnupg
    sudo install -m 0755 -d /etc/apt/keyrings
    FILE=/etc/apt/keyrings/docker.gpg
    if [ -f "$FILE" ]; then
        sudo rm "$FILE"
    fi
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o "$FILE"
    sudo chmod a+r /etc/apt/keyrings/docker.gpg

    # Add Docker-ce repository to Apt sources and install
    echo \
      "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
      $(. /etc/os-release; echo "$VERSION_CODENAME") stable" | \
      sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    sudo apt update -y
    sudo apt -y install docker-ce
	# Add user to group docker
	sudo usermod -aG docker $USER || true
#
fi

# Check if docker-compose is installed
if command -v docker-compose &>/dev/null; then
    echo "Docker-compose is already installed."
else
    echo "Docker-compose is not installed. Proceeding with installations..."
    # Install docker-compose subcommand
    sudo apt -y install docker-compose-plugin
    sudo ln -sv /usr/libexec/docker/cli-plugins/docker-compose /usr/bin/docker-compose
    docker-compose --version
fi


# Test nvidia-docker
if [[ ! -z "$NVIDIA_PRESENT" ]]; then
    if sudo docker run --rm --gpus all nvidia/cuda:11.0.3-base-ubuntu18.04 nvidia-smi &>/dev/null; then
        echo "nvidia-docker is enabled and working. Exiting script."
    else
        echo "nvidia-docker does not seem to be enabled. Proceeding with installations..."
        distribution=$(. /etc/os-release;echo $ID$VERSION_ID)
        curl -s -L https://nvidia.github.io/nvidia-docker/gpgkey | sudo apt-key add
        curl -s -L https://nvidia.github.io/nvidia-docker/$distribution/nvidia-docker.list | sudo tee /etc/apt/sources.list.d/nvidia-docker.list
        sudo apt-get update && sudo apt-get install -y nvidia-container-toolkit
		sudo bash -c 'cat <<EOF > /etc/docker/daemon.json
{
	"default-runtime": "nvidia",
	"runtimes": {
	"nvidia": {
	"path": "/usr/bin/nvidia-container-runtime",
	"runtimeArgs": []
		}
	}
}
EOF'
        sudo systemctl restart docker
        sudo docker run --rm --gpus all nvidia/cuda:11.0.3-base-ubuntu18.04 nvidia-smi
#
    fi
fi
