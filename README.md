# Raspberry Pi Xmas Tree :christmas_tree: 

Used for Christmas December 2018 on a **Pi Zero W** with [Alpine Linux](https://alpinelinux.org/).

Prerequisites:

1. a [PiHut 3D Xmas Tree](https://thepihut.com/products/3d-xmas-tree-for-raspberry-pi)
2. [Go](https://golang.org/)


## Setup

1. Compile the binary and upload it to the raspberry:

```
make pizerow
```

Then:

```
sudo mv /tmp/xmastree /usr/local/bin/xmastree
sudo mv /tmp/xmastree.openrc /etc/init.d/xmastree
sudo chown root:root /usr/local/bin/xmastree /etc/init.d/xmastree
sudo chmod a+x /etc/init.d/xmastree
```

2. Starts the service with:

```
sudo rc-service xmastree start
```

3. Enable start at boot with:

```
sudo rc-update add xmastree default
```

4. And finally for an Alpine Linux run-from-RAM configuration:

```
$ sudo lbu add /usr/loca/bin/xmastree
$ sudo lbu add /etc/init.d/xmastree
$ sudo lbu commit -d
```

:santa: Ho ho ho
