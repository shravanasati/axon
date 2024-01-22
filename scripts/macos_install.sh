
#!/bin/bash

echo "Downloading axon..."
curl -L "https://github.com/shravanasati/axon/releases/latest/download/axon-darwin-amd64" -o axon

echo "Adding axon into PATH..."

mkdir -p ~/.axon;
mv ./axon ~/.axon
echo "export PATH=$PATH:~/.axon" >> ~/.bashrc

echo "axon installation is completed!"
echo "You need to restart the shell to use axon."
