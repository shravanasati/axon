#!/bin/bash

echo "Downloading axon..."
curl -L "https://github.com/Shravan-1908/axon/releases/latest/download/axon-linux-amd64" -o axon

echo "Adding axon into PATH..."

mkdir -p ~/.axon

chmod u+x ./axon

mv ./axon ~/.axon
echo "export PATH=$PATH:~/.axon" >> ~/.bashrc

echo "axon installation is completed!"
echo "You need to restart the shell to use axon."
