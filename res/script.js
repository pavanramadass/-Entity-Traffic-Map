window.onload = function() {
    get_schedule();
    setInterval(get_schedule, 1000*60*60);
};

function collection_modal() {
    document.getElementById("modal").style.display = "block";
    document.getElementById("collection_modal").style.display = "block";
    document.getElementById("edit_modal").style.display = "none";
    document.getElementById("download_modal").style.display = "none";
    document.getElementById("heat_map_modal").style.display = "none";
}

function edit_modal() {
    document.getElementById("modal").style.display = "block";
    document.getElementById("collection_modal").style.display = "none";
    document.getElementById("edit_modal").style.display = "block";
    document.getElementById("download_modal").style.display = "none";
    document.getElementById("heat_map_modal").style.display = "none";
}

function download_modal() {
    document.getElementById("modal").style.display = "block";
    document.getElementById("collection_modal").style.display = "none";
    document.getElementById("edit_modal").style.display = "none";
    document.getElementById("download_modal").style.display = "block";
    document.getElementById("heat_map_modal").style.display = "none";
    download_data();
}

function heat_map_modal() {
    var button = document.getElementById("generate_button");
    if (button.innerHTML != 'Generate Heat Map') {
        document.getElementById("cam-iframe").width = 640;
        document.getElementById("cam-iframe").height = 442;
        document.getElementById("heatmap-img").width = 0;
        document.getElementById("heatmap-img").height = 0;  
        button.innerHTML = 'Generate Heat Map';
        return;
    }
    document.getElementById("modal").style.display = "block";
    document.getElementById("collection_modal").style.display = "none";
    document.getElementById("edit_modal").style.display = "none";
    document.getElementById("download_modal").style.display = "none";
    document.getElementById("heat_map_modal").style.display = "block";
}

function close_modal() {
    document.getElementById("modal").style.display = "none";
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
            var today = new Date();
            var start_date = new Date(json.Start_Date);
            var end_date = new Date(json.End_Date);
            if (start_date < today && today < end_date) {
                document.getElementById("status_bar").style.backgroundColor = "purple";
                document.getElementById("status_bar").innerHTML = "<p>Curent Status: Collection in Progress<br>Schedule: " 
                    + json.Start_Date + " to " + json.End_Date + "</p>";
            } else if (today > end_date) {
                document.getElementById("status_bar").style.backgroundColor = "#73b504";
                document.getElementById("status_bar").innerHTML = "<p>Curent Status: Collection Completed<br>Schedule: " 
                    + json.Start_Date + " to " + json.End_Date + "</p>";
            } else {
                document.getElementById("status_bar").style.backgroundColor = "orange";
                document.getElementById("status_bar").innerHTML = "<p>Curent Status: Collection Scheduled<br>Schedule: " 
                    + json.Start_Date + " to " + json.End_Date + "</p>";
            }
            document.getElementById("edit_schedule").style.display = "block";
            document.getElementById("cancel_schedule").style.display = "block";
            for (i = 0; i < 2; i++) {
                dates[i+2].value = dates[i].value
                months[i+2].value = months[i].value;
                years[i+2].value = years[i].value;
                dates[i].value = 1;
                months[i].value = 'January'
                years[i].value = 2021;
            }
            close_modal();
        }
    });
}

function edit_schedule() {
    dates = document.getElementsByClassName("bear-dates");
    months = document.getElementsByClassName("bear-months");
    years = document.getElementsByClassName("bear-years");
    start_date = years[2].value + '-' + months[2].value + '-' + dates[2].value;
    end_date = years[3].value + '-' + months[3].value + '-' + dates[3].value;
    $.ajax({
        type: "POST",
        url: "/form",
        data: JSON.stringify({
            "Request_Type": "edit_schedule", 
            "Start_Date": start_date, 
            "End_Date": end_date
        }),
        success: function(response) {
            console.log(response);
            json = JSON.parse(response);
            var today = new Date();
            var start_date = new Date(json.Start_Date);
            var end_date = new Date(json.End_Date);
            if (start_date < today && today < end_date) {
                document.getElementById("status_bar").style.backgroundColor = "purple";
                document.getElementById("status_bar").innerHTML = "<p>Curent Status: Collection in Progress<br>Schedule: " 
                    + json.Start_Date + " to " + json.End_Date + "</p>";
            } else if (today > end_date) {
                document.getElementById("status_bar").style.backgroundColor = "#73b504";
                document.getElementById("status_bar").innerHTML = "<p>Curent Status: Collection Completed<br>Schedule: " 
                    + json.Start_Date + " to " + json.End_Date + "</p>";
            } else {
                document.getElementById("status_bar").style.backgroundColor = "orange";
                document.getElementById("status_bar").innerHTML = "<p>Curent Status: Collection Scheduled<br>Schedule: " 
                    + json.Start_Date + " to " + json.End_Date + "</p>";
            }
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
            json = JSON.parse(response);
            var button = document.getElementById("generate_button");
            if (button.innerHTML == 'Generate Heat Map') {
                console.log('MAP');
                console.log(document.getElementById("cam-iframe"));
                console.log(document.getElementById("heatmap-img"));
                document.getElementById("cam-iframe").width = 0;
                document.getElementById("cam-iframe").height = 0;
                document.getElementById("heatmap-img").width = 640;
                document.getElementById("heatmap-img").height = 442;
                button.innerHTML = 'Return to Map View';
            }
            close_modal();
        }
    });
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
            status_bar.innerHTML = "<p>Curent Status: Collection Canceled</p>";
            status_bar.style.backgroundColor = "red";
            document.getElementById("edit_schedule").style.display = "none";
            document.getElementById("cancel_schedule").style.display = "none";
        }
    });
}

function get_schedule() {
    $.ajax({
        type: "GET",
        url: "/form",
        success: function(response) {
            console.log(response)
            json = JSON.parse(response);
            var today = new Date();
            var start_date = new Date(json.Start_Date);
            var end_date = new Date(json.End_Date);
            if (start_date < today && today < end_date) {
                document.getElementById("status_bar").style.backgroundColor = "purple";
                document.getElementById("status_bar").innerHTML = "<p>Curent Status: Collection in Progress<br>Schedule: " 
                    + json.Start_Date + " to " + json.End_Date + "</p>";
            } else if (today > end_date) {
                document.getElementById("status_bar").style.backgroundColor = "#73b504";
                document.getElementById("status_bar").innerHTML = "<p>Curent Status: Collection Completed<br>Schedule: " 
                    + json.Start_Date + " to " + json.End_Date + "</p>";
            } else {
                document.getElementById("status_bar").style.backgroundColor = "orange";
                document.getElementById("status_bar").innerHTML = "<p>Curent Status: Collection Scheduled<br>Schedule: " 
                    + json.Start_Date + " to " + json.End_Date + "</p>";
            }
        }
    });
}