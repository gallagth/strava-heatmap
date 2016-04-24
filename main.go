package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
    "github.com/strava/go.strava"
)

var accessToken string

func init() {
    accessToken = strings.TrimSpace(readAccessToken())
    http.HandleFunc("/segment", segHandler)
    http.HandleFunc("/polylines", linesHandler)
}

func linesHandler(w http.ResponseWriter, r *http.Request) {
    client := getClient(r)
    service := strava.NewCurrentAthleteService(client)
    activities, err := service.ListActivities().
        Page(1).
        PerPage(50).
        Do()
    if err != nil {
        fmt.Fprint(w, err)
        return
    }
    for _,act := range activities {
        fmt.Fprintf(w, "%s\n", act.Map.SummaryPolyline)
    }
}

func segHandler(w http.ResponseWriter, r *http.Request) {
    var segmentId int64 = 229781

    client := getClient(r)

    segment, err := strava.NewSegmentsService(client).Get(segmentId).Do()

    if err != nil {
        fmt.Fprint(w, err)
        return
    }

    verb := "ridden"
    if segment.ActivityType == strava.ActivityTypes.Run {
        verb = "run"
    }
    fmt.Fprintf(w, "%s, %s %d times by %d athletes\n\n", segment.Name, verb, segment.EffortCount, segment.AthleteCount)
}

func readAccessToken() string {
    dat, err := ioutil.ReadFile("private/access.token")
    check(err)
    return string(dat)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getClient(r *http.Request) (*strava.Client) {
    urlFetchClient := urlfetch.Client(appengine.NewContext(r))
    return strava.NewClient(accessToken, urlFetchClient)
}