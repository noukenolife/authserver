#!/bin/bash

cd .ssl
mkcert -key-file server.key -cert-file server.crt 0.0.0.0 localhost 127.0.0.1 ::1
cd ../
