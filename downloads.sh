#!/usr/bin/bash
cd static/

tree  temp-images -H 'temp-images'  --noreport --charset utf-8 > indexses.html

sed -i 's_<style type="text/css">_<link rel="stylesheet" type="text/css" href="downloads.css">_' indexses.html

sed -i 's_Directory Tree_Downloads_' indexses.html

sed -i '8i<meta name ="viewport" content="width=device-width, initial-scale=1">' indexses.html

sed -i 's_<h1>Downloads</h1><p>_<header><nav class="container"><a class="logo" href="../"><span>F</span><span>T</span><span>W</span><span>A</span></a><div class="nav-toggle"><span></span></div><ul id="menu"><li><a href="">About</a></li><li><a href="/temp-images">Downloads</a></li></ul></nav></header>_' indexses.html

sed -i '/<a href="temp-images">te/d' indexses.html
sed -i '/Steve/d' indexses.html
sed -i '/Moore/d' indexses.html
sed -i '/Rocher/d' indexses.html
sed -i '/Sesser/d' indexses.html
sed -i '/<hr>/d' indexses.html
sed -i 's_└──_ _' indexses.html
sed -i 's_├──_ _' indexses.html
sed -i '/Tokoro/d' indexses.html
sed -i '/indexses/d' indexses.html
sed -i '/style.css">style.css</d' indexses.html
