#!/usr/bin/bash

rsync --archive --compress --delete ./bin/main root@127.0.0.1:/var/www/api.example.com/bin/

# Use scp to copy files to server
# scp -r bin/main root@127.0.0.1:/var/www/api.example.com/bin/