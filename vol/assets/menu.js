window.onload = function() {
  // document.getElementById('PYstring').focus();
  elem = document.getElementById('container');
  if(elem != null){
    elem.onkeydown = function(e){
      // alert(String.fromCharCode(e.keyCode)+" --> "+e.keyCode);
      if(e.keyCode==80) pinyintozi();
      if(e.keyCode==90) zitopinyin();
    };
    elem.focus();
  };  
};

function setmenu(option) {
    // get the origin of the url 
    var host = window.location.origin;
    
    window.location.replace(host + option );
}
function pinyintozi() {
    var retVal = prompt("Enter Pinyin (+ tone if you know it) : ");
    if(retVal !="" && retVal.length <8) setmenu("/listpy/"+retVal )
    else  setmenu("/listpy/invalid%20data");
}
function zitopinyin() {
    var retVal = prompt("Enter chinese character : ");
    if(retVal !="" && retVal.length <8) setmenu("/listzi/"+retVal )
    else  setmenu("/listzi/invalid%20data");
}

function add(zi) {
    var s= document.getElementById("zistring");
    s.value = s.value + zi;
    entree = document.getElementById("zilist"); // reset list of displayed zi buttons
    if(entree != null)entree.innerHTML = "";
    document.getElementById('PYstring').focus();
}

function copyTextToClipboard() {
    // source : https://stackoverflow.com/questions/400212/how-do-i-copy-to-the-clipboard-in-javascript
    var text = document.getElementById("zistring").value ;
    navigator.clipboard.writeText(text).then(function() {
        console.log('Async: Copying to clipboard was successful!');
    }, function(err) {
        console.error('Async: Could not copy text: ', err);
    });
    document.getElementById('PYstring').focus();
}

function lookup() {
  var text = document.getElementById("zistring").value ;
  var url = "https://dictionary.writtenchinese.com/#sk="+text+"&svt=pinyin";
  window.open(url);
  // document.getElementById('PYstring').focus();
}

  function reset(){
    var entree = document.getElementById("zistring");
    if (entree != null) entree.value = "";
    entree = document.getElementById("zilist");
    if(entree != null)entree.innerHTML = "";
    document.getElementById('PYstring').value = "" ;
    document.getElementById('PYstring').focus();
  }

  function visible(){
    document.getElementById("solution").style.visibility = "visible" ;
    document.getElementById("retry").style.visibility = "visible" ;
    document.getElementById("retry").focus({ focusVisible: true }) ;
  }

  function convertToZi(){ 
    var pyentree = document.getElementById('PYstring').value ;
    var lon = pyentree.length;
    if (lon ==0)return;
    var dernier = pyentree.charAt(lon-1) ;
    if ('01234/ '.includes(dernier)){
      if(dernier=='/' || dernier ==' ' ) pyentree = pyentree.substring(0,lon-1);
      var entree = document.getElementById("zistring");
      var currentZiString;
      if(entree !=null && entree.value != "")currentZiString = entree.value;
      else currentZiString = "v";
      if(pyentree!="") setmenu("/zistring/"+pyentree+"/"+currentZiString); // state conservation via the url
    }
  }

  
