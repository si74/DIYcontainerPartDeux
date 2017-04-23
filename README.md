# diyContainer

Containers: What, Why, How?

This repo is the companion to a blog post on using linux namespaces for isolation -
a.k.a. how to make your own container. 

## Preparation

Many of the examples here will work only a Linux based distribution.
As a result, I recommend using the vagrant manager along with the VirtualBox hypervisor.

Download vagrant and virtual box on the MacOSX if you have not already done so:

`brew cask install virtualbox`

`brew cask install vagrant`

The vagrant file is in the repository so all you need to do is:

`vagrant up`

Note: as per the docs, the current working directory is shared in the /vagrant folder of the guest.

Install go and set the $GOPATH:

`sudo apt-get install golang
source env.sh`
