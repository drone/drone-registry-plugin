This plugin provides global registry credentials for pulling pipeline images. This plugin is a direct port of the global registry credentials file in Drone 0.8. See https://0-8-0.docs.drone.io/setup-global-registry-credentials/

# Credentials File

Registry credentials are loaded from a yaml configuration file. Example registry credentials file:

```
- address: docker.io
  username: octocat
  password: correct-horse-batter-staple
- address: gcr.io
  username: _json_key
  password: |
    {
      "private_key_id": "...",
      "private_key": "...",
      "client_email": "...",
      "client_id": "...",
      "type": "..."
    }
```

Example registry credentials for an ECR repository:

```
- address: 012345678910.dkr.ecr.us-east-1.amazonaws.com
  aws_access_key_id: a50d28f4dd477bc184fbd10b376de753
  aws_secret_access_key: bc5785d3ece6a9cdefa42eb99b58986f9095ff1c
```

# Installation

Create a shared secret:

```
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

Configure the plugin:

```
DRONE_DEBUG=true
DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6
DRONE_CONFIG_FILE=/path/to/config.yml
```

Start the plugin:

```
docker run \
 --detach \
 --restart=always \
 -p 3000:3000 \
 -e DRONE_DEBUG=true \
 -e DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
 -e DRONE_CONFIG_FILE=/path/to/config.yml \
 -v /path/to/config.yml:/path/to/config.yml \
 drone/registry-plugin
```

Configure your runner to use the plugin:

```
DRONE_REGISTRY_ENDPOINT=http://plugin.address:3000
DRONE_REGISTRY_SECRET=bea26a2221fd8090ea38720fc445eca6
```

# Testing

Use the command line utility to test the plugin:

```
$ DRONE_REGISTRY_ENDPOINT=http://plugin.address:3000
$ DRONE_REGISTRY_SECRET=bea26a2221fd8090ea38720fc445eca6
$ 
$ drone plugins registry list
index.docker.io 
Username:  octocat
Password: correct-horse-battery-staple
```
