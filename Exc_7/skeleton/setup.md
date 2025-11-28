# todo note commands
Since we were allowed to do this exercise alone as well, I went for a single-machine Swarm setup.  
I ran everything locally on my Windows laptop using Docker Desktop.

To reset any old configuration, I first left the previous swarm and created a fresh one:
"docker swarm leave --force"
"docker swarm init"

After that I cleaned up leftover local networks that were blocking the deploy:
"docker network rm intercom"
"docker network rm web"


Then I simply deployed the stack with the provided compose file:
"docker stack deploy -c docker-compose.yml sbd3ue"


Whenever I changed something in the compose file, I redeployed by removing and deploying again:
"docker stack rm sbd3ue"
"docker stack deploy -c docker-compose.yml sbd3ue"