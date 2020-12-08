#!/usr/bin/bash
cd static/

tree -t -r temp-images -H 'temp-images' --noreport --charset utf-8 > downloads.html

sed -i 's_<style type="text/css">_<link rel="stylesheet" type="text/css" href="downloads.css">_' downloads.html

sed -i 's_Directory Tree_Downloads_' downloads.html

sed -i '8i<meta name ="viewport" content="width=device-width, initial-scale=1">' downloads.html

sed -i 's_<p class="VERSION">_<script src="/u/jquery.min.js"></script><script src="/u/switcher.js"></script>_' downloads.html

sed -i '33i<ol class="rounded">' downloads.html

#sed -i "35i</ol>" downloads.html

sed -i 's_</a><br></br>_</a><ol/><br></br>_' downloads.html


sed -i 's_<h1>Downloads</h1><p>_<header><nav class="container"><a class="logo" href="../"><span>F</span><span>T</span><span>W</span><span>A</span></a><div class="nav-toggle"><span></span></div><ul id="menu"><div class="text">Dark Mode:   </div><div class="toggle-switch"><label class="switch"><input id="theme-switch" type="checkbox"><span class="slider round"></span></span></div><li><a href="">About</a></li><li><a href="u/index.html">Upload</a></li></ul></nav></header><a href="home.html" class="home">Go Back</a>_' downloads.html

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
