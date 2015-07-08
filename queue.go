package gojenkins

type Queue struct {
    Items []Item `json:"items"`
}

type Item struct {
    Actions []Action `json:"actions"`
    Blocked bool `json:"blocked"`
    Buildable bool `json:"buildable"`
    Id int `json:"id"`
    InQueueSince int64 `json:"inQueueSince"`
    Params string `json:"params"`
    Stuck bool `json:"stuck"`
    Task Task `json:"task"`
    URL string `json:"url"`
    Why string `json:"why"`
    BuildableStartMilliseconds int64 `json:"buildableStartMilliseconds"`
    Pending bool `json:"pending"`
}

type Action struct {
    Causes []Cause
}

type Cause struct {
    ShortDescription string `json:"shortDescription"`
    UserId string `json:"userId"`
    UserName string `json:"userName"`
}

type Task struct {
    Name string `json:"name"`
    Url string `json:"url"`
    Color string `json:"color"`
}
