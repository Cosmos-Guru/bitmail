# How to Bootstrap a Bitmail Node Tutorial

For this tutorial we will be creating a node and connecting it to the bitmailtestnet, version 001. Note it has a `chain-id` = `bitmailtestnet001` . For this tutorial we assume you have some basic knowledge of Linux and how to SSH into a VPS cloud server.

## Step 1: Log into your VPS. 

```bash
ssh username@your.node.ip.address
````

Note: We recommend not using `root` as login, we recommend creating a `newuser`, giving it `sudo` priveleges and removing root login capability.  For instructions on how to od so, please click HERE

## Step 2: Set your Environment Variables

You will need to update your bash profile with the ENV variables below. Note, for this tutorial, we are using Linux Ubuntu, if you use mac, this file may be called `.zhrc` of something similar, you can find these files in your `$HOME`Update bashrc and recome from shell script. 

```
cd $HOME
nano .bashrc
```

#### Copy & Paste into your `.bashrc` file
```bash
# Set GOPATH and optionally GOROOT if needed
export GOPATH=$HOME/go
# Optionally set GOROOT if you're using a custom Go installation
export GOROOT=/usr/local/go
# Consolidate PATH updates
export PATH=$PATH:$GOPATH/bin:/usr/local/bin:/usr/local/go/bin
```

Upadate your shell using the `source` command or start at new shell prompt. 

```
source ~/.bashrc
```

## Step 3: Create a bootstap.sh Shell Script 

The following should be done from your `$HOME` directory. 

```bash
cd $HOME
```
Create the bootstrap.sh file

```bash
nano bootstrap.sh
```




Copy and paste contents below into the `bootstrap.sh` file

```
#!/bin/bash
# This Script is intended to be installed on a VPS either is root or you have already created a user name and have proper SSH access privileges. 
sudo apt update -y
sudo apt install build-essential -y
sudo apt install gcc -y
sudo apt install git -y
sudo apt install wget -y
sudo apt install curl -y

#On Linux
sudo wget https://go.dev/dl/go1.23.1.linux-amd64.tar.gz
#unpack it
sudo tar -C /usr/local -xzf go1.23.1.linux-amd64.tar.gz
source ~/.bashrc
go version
#should output golang 1.23.1

curl https://get.ignite.com/cli@v0.27.1 | bash
sudo mv ignite /usr/local/bin
git clone https://github.com/Cosmos-Guru/bitmail.git
cd bitmail
ignite chain build

bitmaild init mytestnetnode --default-denom ubtml --chain-id bitmailtestnet001
 
curl http://172.233.144.166:26657/genesis | jq '.result.genesis' > ~/.bitmail/config/genesis.json
sed -i '/\[api\]/,+3 s/enable = false/enable = true/' ~/.bitmail/config/app.toml
sed -i '/\[api\]/,+3 s/swagger = false/swagger = true/' ~/.bitmail/config/app.toml
sed -i 's/minimum-gas-prices = "0stake"/minimum-gas-prices = "0.0025bitmail"/' ~/.bitmail/config/app.toml
sed -i 's/^persistent_peers = .*/persistent_peers = "288d66e8c02539a430598ac5ffa01795584531e3@172.233.144.166:26656,51e019d3decdd7077a07b7110cad0ea5c49bdf3f@194.195.118.93:26656"/' ~/.bitmail/config/config.toml
sed -i 's/^seeds = .*/seeds = "288d66e8c02539a430598ac5ffa01795584531e3@172.233.144.166:26656,51e019d3decdd7077a07b7110cad0ea5c49bdf3f@194.195.118.93:26656"/' ~/.bitmail/config/config.toml
sed -i 's/timeout_commit = "5s"/timeout_commit = "20s"/' ~/.bitmail/config/config.toml
sed -i 's#^laddr = "tcp://127.0.0.1:26657"#laddr = "tcp://0.0.0.0:26657"#' ~/.bitmail/config/config.toml
bitmaild start
```

Next we need to turn our shell script into an executable
```
chmod +x boostrap.sh
```

```
./bootstrap.sh
```

It will take some time for this to complete but when it does, you should see it start syncing blocks with the network. 

You can then view your API and RPC endpoints, and submit transactions to their endpoints. 

#### For API
your.node.ip.address:1317

#### For RPC
your.node.ip.address:26657


CONGRATULATIONS!!!! You have successfully connected to the network. In our next tutorial, I will walk you through on how to update your node to a Validator. 



## Other References

### How to create a new user and disable root login

```
#You can change username to whatever you like. 
sudo adduser username

sudo usermod -aG sudo username

sudo nano /etc/ssh/sshd_config

#Change the PermitRootLogin from yes to no. 
PermitRootLogin no
```

