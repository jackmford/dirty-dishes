function changeColor() {
  if (document.getElementById("redbutton") != null){
    document.getElementById("redbutton").innerHTML = "Clean";
    document.getElementById("redbutton").id = "bluebutton";
  }
  else{
    document.getElementById("bluebutton").innerHTML = "Dirty";
    document.getElementById("bluebutton").id = "redbutton";
  }
}

