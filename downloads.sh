#!/usr/bin/bash
cd static/temp-images

tree -H '' --noreport --charset utf-8 > indexses.html
sed -i 's_<style type="text/css">_<style type="text/css" href="style.css">_' indexses.html


