function changeColor() {
  if (document.getElementById("redbutton") != null){
    document.getElementById("redbutton").innerHTML = "Clean";
    document.getElementById("redbutton").id = "bluebutton";
    let btn-color = "bluebutton";
  }
  else{
    document.getElementById("bluebutton").innerHTML = "Dirty";
    document.getElementById("bluebutton").id = "redbutton";
    let btn-color = "redbutton";
  }

  var data = new FormData();
  data.append("color", btn-color);
  fetch("localhost:8787/write", { method: "POST", body: data})
}

