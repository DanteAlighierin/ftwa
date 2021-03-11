if (window.FileReader) {					     
	var reader = new FileReader();
						        
	reader.onload = function (oFREvent) { 
		preview = document.getElementById("uploadPreview")
		preview.src = oFREvent.target.result;  
		preview.style.display = "block";

		

};




function doTest() {
		if (document.getElementById("myFile").files.length === 0) { return; }  
		var file = document.getElementById("myFile").files[0];
		reader.readAsDataURL(file); 
	}					          
}
else {					
	alert("FileReader object not found :( \nTry using Chrome, Firefox or WebKit");



							      }
