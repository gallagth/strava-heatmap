package hello

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "google.golang.org/appengine"
    "google.golang.org/appengine/urlfetch"
    "github.com/strava/go.strava"
)

func init() {
    http.HandleFunc("/activities", handler)
    http.HandleFunc("/segment", segHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello Thomas!")
}

func segHandler(w http.ResponseWriter, r *http.Request) {
    var segmentId int64 = 229781
    var accessToken string = readAccessToken()

    ctx := appengine.NewContext(r)
    urlfetch_client := urlfetch.Client(ctx)
    client := strava.NewClient(accessToken, urlfetch_client)

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