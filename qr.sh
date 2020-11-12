#!/usr/bin/bash
curl -s https://myexternalip.com/raw > out.txt && echo :8080 >> out.txt && cat out.txt | qrencode -t ansiutf8
