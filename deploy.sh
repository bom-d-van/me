#!/usr/bin/env bash

scp /Users/bom_d_van/Code/go/workspace/src/github.com/bom-d-van/me/server.sh root@bomdvan.com:~/server.sh
ssh root@bomdvan.com "chmod +x ~/server.sh && sudo ~/server.sh"
