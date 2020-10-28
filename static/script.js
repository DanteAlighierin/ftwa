if (window.FileReader) {
						      
						        var reader = new FileReader(), rFilter = /^(image\/bmp|image\/cis-cod|image\/gif|image\/ief|image\/jpeg|image\/jpeg|image\/jpeg|image\/pipeg|image\/png|image\/svg\+xml|image\/tiff|image\/x-cmu-raster|image\/x-cmx|image\/x-icon|image\/x-portable-anymap|image\/x-portable-bitmap|image\/x-portable-graymap|image\/x-portable-pixmap|image\/x-rgb|image\/x-xbitmap|image\/x-xpixmap|image\/x-xwindowdump)$/i; 
						        
						        reader.onload = function (oFREvent) { 
								        preview = document.getElementById("uploadPreview")
								        preview.src = oFREvent.target.result;  
								        preview.style.display = "block";
								      };  
						    
						        function doTest() {
								        
								        if (document.getElementById("myFile").files.length === 0) { return; }  
								        var file = document.getElementById("myFile").files[0];  
								        if (!rFilter.test(file.type)) { alert("You must select a valid image file!"); return; }  
								        reader.readAsDataURL(file); 
								      }
						          
						    } else {
							        alert("FileReader object not found :( \nTry using Chrome, Firefox or WebKit");
							      }
