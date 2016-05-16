var map;
var heatmapData;

function initMap() {
  // Create a map object and specify the DOM element for display.
    map = new google.maps.Map(document.getElementById('map'), {
        center: {lat: 48.853, lng: 2.3499},
        scrollwheel: false,
        zoom: 13
    });
    heatmapData = new google.maps.MVCArray();
    heatmap = new google.maps.visualization.HeatmapLayer({
        data: heatmapData,
        map: map,
        radius: 50,
        opacity: 0.5
    });
    drawAllLines();
}

function drawEncodedPolyline(encodedPolyline) {
    var decodedPath = google.maps.geometry.encoding.decodePath(encodedPolyline);
    for (var i = 0; i < decodedPath.length; i++) {
        heatmapData.push(decodedPath[i]);
    }
    var line = new google.maps.Polyline({
        path: decodedPath,
        geodesic: true,
        strokeColor: '#FF0000',
        strokeOpacity: 1.0,
        strokeWeight: 2
    });
    line.setMap(map);
}

function drawAllLines() {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function() {
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
            var lines = xmlHttp.responseText.split("\n");
            var numLines = lines.length;
            for (var i = 0; i < numLines; i++) {
                drawEncodedPolyline(lines[i]);
            }
        }
    }
    xmlHttp.open("GET", "/polylines", true); // true for asynchronous
    xmlHttp.send(null);
}