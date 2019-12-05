# Raspberry Pi Xmas Tree :christmas_tree: 

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Prerequisites:

1. a [PiHut 3D Xmas Tree](https://thepihut.com/products/3d-xmas-tree-for-raspberry-pi)
2. [Go](https://golang.org/)

Installation instructions are for a Raspberry Pi powered by [Alpine Linux](https://alpinelinux.org/).

:warning: **Do NOT place the Xmas Tree on your Raspberry Pi backwards or bad things will happen!**

![Assembled with PiZero](https://github.com/vinymeuh/raspi-xmastree/blob/master/assets/assembled_pizero.jpg)

## Manual setup

1. Compile the binary for your favorite Raspberry:

```
make build-arm6
```

2. Upload **xmastree** and **xmastree.openrc** under ```/tmp``` on the Raspberry


3. Connect to the Raspberry and move files to their final destination

```
sudo mv /tmp/xmastree /usr/local/bin/xmastree
sudo mv /tmp/xmastree.openrc /etc/init.d/xmastree
sudo chown root:root /usr/local/bin/xmastree /etc/init.d/xmastree
sudo chmod a+x /etc/init.d/xmastree
```

4. Create the user **santa** :santa:

```
sudo adduser -D -H santa
sudo addgroup santa gpio
```

5. Starts the service with ```sudo rc-service xmastree start```

6. Enable start at boot with ```sudo rc-update add xmastree default```

7. And finally

```
$ sudo lbu add /usr/local/bin/xmastree
$ sudo lbu add /etc/init.d/xmastree
$ sudo lbu commit -d
```

:santa: Ho ho ho
