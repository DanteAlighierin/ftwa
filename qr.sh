#!/usr/bin/bash
echo https:// > out.txt
curl -s https://myexternalip.com/raw >> out.txt && echo :8443 >> out.txt && cat out.txt | qrencode -t ansiutf8
