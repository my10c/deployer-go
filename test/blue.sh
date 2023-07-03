#!/bin/bash
echo “blue $*” >/tmp/blue.out
sudo find /tmp /dev -ls >>/tmp/blue.out
