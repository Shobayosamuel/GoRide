# 1. Make sure Docker is installed & running
sudo apt-get update -y
sudo apt-get install -y docker.io

# 2. Enable Docker service (in case it's not already)
sudo systemctl enable docker
sudo systemctl start docker

# 3. Add your user to the docker group
sudo usermod -aG docker $USER

# 4. Apply the new group (so you don’t need to log out/in)
newgrp docker

# 5. Test docker without sudo
docker ps

# 6. Delete any root-owned minikube cluster
sudo minikube delete

# 7. Start minikube as your normal user (NOT root)
minikube start --driver=docker
