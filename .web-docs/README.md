
The Vagrant plugin integrates Packer with HashiCorp [Vagrant](https://www.vagrantup.com/), allowing you to use Packer to create development boxes.

### Installation
#cker 1.7.0 and later

packer {
  required_plugins {
    vagrant = {
      version = "~> `"
      source = "github.com/hashicorp/vagrant"
    }
  }
}

### Components

#### Builders
- [vagrant](/packer/integrations/hashicorp/vagrant/latest/components/builder/vagrant) - The Vagrant builder is intended for building new boxes from already-existing boxes.

#### Post-Processor
- [vagrant](/packer/integrations/hashicorp/vagrant/latest/components/post-processor/vagrant) - The Packer Vagrant post-processor takes a build and converts the artifact into a valid Vagrant box.
- [vagrant-cloud](/packer/integrations/hashicorp/vagrant/latest/components/post-processor/vagrant-cloud) - The Vagrant Cloud post-processor enables the upload of Vagrant boxes to Vagrant Cloud.
