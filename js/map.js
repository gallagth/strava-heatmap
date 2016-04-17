function initMap() {
  // Create a map object and specify the DOM element for display.
  var map = new google.maps.Map(document.getElementById('map'), {
    center: {lat: 48.853, lng: 2.3499},
    scrollwheel: false,
    zoom: 13
  });
}