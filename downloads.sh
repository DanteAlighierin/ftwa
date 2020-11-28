#!/usr/bin/bash
cd static/

tree -t -r temp-images -H 'temp-images' --noreport --charset utf-8 > downloads.html

sed -i 's_<style type="text/css">_<link rel="stylesheet" type="text/css" href="downloads.css">_' downloads.html

sed -i 's_Directory Tree_Downloads_' downloads.html

sed -i '8i<meta name ="viewport" content="width=device-width, initial-scale=1">' downloads.html

sed -i 's_<h1>Downloads</h1><p>_<header><nav class="container"><a class="logo" href="../"><span>F</span><span>T</span><span>W</span><span>A</span></a><div class="nav-toggle"><span></span></div><ul id="menu"><li><a href="">About</a></li><li><a href="downloads.html">Downloads</a></li></ul></nav></header><a href="home.html">Go Back</a>_' downloads.html

sed -i '/<a href="temp-images">te/d' downloads.html
sed -i '/Steve/d' downloads.html
sed -i '/Moore/d' downloads.html
sed -i '/Rocher/d' downloads.html
sed -i '/Sesser/d' downloads.html
sed -i '/<hr>/d' downloads.html
sed -i 's_└──_ _' downloads.html
sed -i 's_├──_ _' downloads.html
sed -i '/Tokoro/d' downloads.html
sed -i '/indexses/d' downloads.html
sed -i '/style.css">style.css</d' downloads.html
sed -i '/Go%20Back/d' downloads.html
