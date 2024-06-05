#!/bin/bash
set -e

# required root

# apt update -y;
# apt upgrade -y;
# reboot

apt -y update
add-apt-repository universe
apt apt -y install git curl wget libnewt-dev libssl-dev libncurses5-dev subversion libsqlite3-dev build-essential libjansson-dev libxml2-dev uuid-dev

cd ~/tmp
wget http://downloads.asterisk.org/pub/telephony/asterisk/asterisk-20-current.tar.gz

tar xvf asterisk-20-current.tar.gz
cd asterisk-20.8.1/ # 細かいバージョンが変わるかも
contrib/scripts/get_mp3_source.sh
sudo contrib/scripts/install_prereq install

./configure
make menuselect

## 変なTUIが出るので、JA-WAVっぽいやつにチェックを付けておく ##

make
make install
make progdocs
make samples
make config
ldconfig
