function changeColor(value) {
  let btncolor = "";
  if (value == "dirty_dishes"){
    document.getElementById("dirty_dishes").innerHTML = "Clean";
    document.getElementById("dirty_dishes").id = "clean_dishes";
    btncolor = "clean_dishes";
  }
  else if (value == "clean_dishes") {
    document.getElementById("clean_dishes").innerHTML = "Dirty";
    document.getElementById("clean_dishes").id = "dirty_dishes";
    btncolor = "dirty_dishes";
  }
  else if (value == "running") {
    document.getElementById("not_running").innerHTML = "Running";
    document.getElementById("not_running").id = "running";
  }
  else {
    document.getElementById("running").innerHTML = "Not running";
    document.getElementById("running").id = "not_running";
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

function addItem() {
  document.getElementById("grocery_list");
  console.log('test')
}
