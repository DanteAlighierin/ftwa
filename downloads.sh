#!/usr/bin/bash
cd static/temp-images

tree -H '' --noreport --charset utf-8 > indexses.html
sed -i 's_<style type="text/css">_<link rel="stylesheet" type="text/css" href="style.css">_' indexses.html

sed -i 's_Directory Tree_Downloads_' indexses.html

sed -i '8i<meta name ="viewport" content="width=device-width, initial-scale=1">' indexses.html

sed -i 's_<h1>Downloads</h1><p>_<header><nav class="container"><a class="logo" href="../"><span>F</span><span>T</span><span>W</span><span>A</span></a><div class="nav-toggle"><span></span></div><ul id="menu"><li><a href="">About</a></li><li><a href="/temp-images">Downloads</a></li></ul></nav></header>_' indexses.html


sed -i 's_├──_ _' indexses.html

sed -i '52,63d' indexses.html

sed -i '/Tokoro/d' indexses.html

sed -i '9,26d' indexses.html
