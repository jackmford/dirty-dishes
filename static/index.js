function changeDishesColor() {
  let btncolor = "";
  if (document.getElementById("dirty_dishes") != null ) {
    document.getElementById("dirty_dishes").innerHTML = "Clean";
    document.getElementById("dirty_dishes").id = "clean_dishes";
    btncolor = "clean_dishes";
  }
  else {
    document.getElementById("clean_dishes").innerHTML = "Dirty";
    document.getElementById("clean_dishes").id = "dirty_dishes";
    btncolor = "dirty_dishes";
  }
  var data = new FormData();
  data.append("color", btncolor);
  fetch("http://localhost:8787/write", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    //body: JSON.stringify(data)
    body: {"color": btncolor},
  })
  .then(response => response.json())
  .then(data => {
    console.log("success");
  })
  .catch((error) => {
    console.error("error");
  });
}

function changeRunColor() {
  if (document.getElementById("not_running") != null) {
    document.getElementById("not_running").innerHTML = "Running";
    document.getElementById("not_running").id = "running";
  }
  else {
    document.getElementById("running").innerHTML = "Not running";
    document.getElementById("running").id = "not_running";
  }
}

function addItem() {
  document.getElementById("grocery_list");
  console.log('test')
}
