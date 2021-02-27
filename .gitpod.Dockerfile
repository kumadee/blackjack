FROM gitpod/workspace-full-vnc

# Install custom tools, runtimes, etc.
# For example "bastet", a command-line tetris clone:
# RUN brew install bastet
#
# More information: https://www.gitpod.io/docs/config-docker/
RUN pip install pre-commit && \
        sudo apt-get -q update && \
        sudo apt-get install -yq graphviz && \
        sudo rm -rf /var/lib/apt/lists/*