function changeColor() {
  let btncolor = "";
  if (document.getElementById("redbutton") != null){
    document.getElementById("redbutton").innerHTML = "Clean";
    document.getElementById("redbutton").id = "bluebutton";
    btncolor = "bluebutton";
  }
  else{
    document.getElementById("bluebutton").innerHTML = "Dirty";
    document.getElementById("bluebutton").id = "redbutton";
    btncolor = "redbutton";
  }
  var data = new FormData();
  data.append("color", btncolor);
  fetch("http://localhost:8787/write", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    //body: JSON.stringify(data)
    body: {"color": btncolor}
  })
  .then(response => response.json())
  .then(data => {
    console.log("success");
  })
  .catch((error) => {
    console.error("error");
  });
}

