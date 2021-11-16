function collection_modal() {
    document.getElementById("modal").style.display = "block";
    document.getElementById("collection_modal").style.display = "block";
    document.getElementById("download_modal").style.display = "none";
    document.getElementById("heat_map_modal").style.display = "none";
}

function download_modal() {
    document.getElementById("modal").style.display = "block";
    document.getElementById("collection_modal").style.display = "none";
    document.getElementById("download_modal").style.display = "block";
    document.getElementById("heat_map_modal").style.display = "none";
    download_data();
}

function heat_map_modal() {
    var button = document.getElementById("generate_button");
    var cameraView = document.getElementById("camera-view");
    if (button.innerHTML != 'Generate Heat Map') {
        button.innerHTML = 'Generate Heat Map';
        cameraView.style.backgroundImage = "url('res/image/stock.jpg')";
        return;
    }
    document.getElementById("modal").style.display = "block";
    document.getElementById("collection_modal").style.display = "none";
    document.getElementById("download_modal").style.display = "none";
    document.getElementById("heat_map_modal").style.display = "block";
}

function close_modal() {
    document.getElementById("modal").style.display = "none";
}

function edit_schedule() {
    collection_modal();
}

function data_collection() {
    dates = document.getElementsByClassName("bear-dates");
    months = document.getElementsByClassName("bear-months");
    years = document.getElementsByClassName("bear-years");
    start_date = years[0].value + '-' + months[0].value + '-' + dates[0].value;
    end_date = years[1].value + '-' + months[1].value + '-' + dates[1].value;
    $.ajax({
        type: "POST",
        url: "/form",
        data: JSON.stringify({
            "Request_Type": "data_schedule", 
            "Start_Date": start_date, 
            "End_Date": end_date
        }),
        success: function(response) {
            console.log(response);
            json = JSON.parse(response);
            document.getElementById("status_bar").style.backgroundColor = "#22bd0d";
            document.getElementById("status_bar").innerHTML = "<p>Curent Status: Collection Scheduled<br>Schedule: " 
                + json.Start_Date + " to " + json.End_Date + "</p>";
            document.getElementById("edit_schedule").style.display = "block";
            document.getElementById("cancel_schedule").style.display = "block";
            close_modal();
        }
    });
}

function generate_map() {
    checkbox = document.getElementById("data_checkbox");
    if (checkbox.checked) {
        data_content = "local_data";
    } else {
        data_content = "specific_data";
    }
    $.ajax({
        type: "POST",
        url: "/form",
        data: JSON.stringify({
            "Request_Type": "map_generation",
            "Data_Content": data_content
        }),
        success: function(response) {
            console.log(response);
            var button = document.getElementById("generate_button");
            var cameraView = document.getElementById("camera-view");
            if (button.innerHTML == 'Generate Heat Map') {
                button.innerHTML = 'Return to Map View';
                cameraView.style.backgroundImage = "url('res/image/filled.png')";
            } else {
                button.innerHTML = 'Generate Heat Map';
                cameraView.style.backgroundImage = "url('res/image/stock.jpg')";
            }
            close_modal();
        }
    })
}

function cancel_schedule() {
    $.ajax({
        type: "POST",
        url: "/form",
        data: JSON.stringify({
            "Request_Type": "cancel_schedule"
        }),
        success: function(response) {
            console.log(response);
            var status_bar = document.getElementById("status_bar");
            status_bar.innerHTML = "<p>Curent Status: Collection Scheduled <br> Schedule:</p>";
            status_bar.style.backgroundColor = "orange";
            document.getElementById("edit_schedule").style.display = "none";
            document.getElementById("cancel_schedule").style.display = "none";
        }
    })
}