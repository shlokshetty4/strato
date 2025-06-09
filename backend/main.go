package main

import (
    "encoding/json"
    "log"
    "net/http"
    "time"
)

type User struct {
    Name                  string
    CreateDate            time.Time
    PasswordChangedDate   time.Time
    LastAccessDate        time.Time
    MFAEnabled            bool
}

type UserResponse struct {
    Name                      string `json:"name"`
    CreateDate                string `json:"createDate"`
    PasswordChangedDate       string `json:"passwordChangedDate"`
    DaysSincePasswordChange   int    `json:"daysSincePasswordChange"`
    LastAccessDate            string `json:"lastAccessDate"`
    DaysSinceLastAccess       int    `json:"daysSinceLastAccess"`
    MFAEnabled                bool   `json:"mfaEnabled"`
}

func computeDaysAgo(t time.Time) int {
    return int(time.Since(t).Hours() / 24)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
    // sample data matching your test set
    users := []User{
        {
            Name:                "Foo Bar1",
            CreateDate:          time.Date(2020, time.October, 1, 0, 0, 0, 0, time.UTC),
            PasswordChangedDate: time.Date(2021, time.October, 1, 0, 0, 0, 0, time.UTC),
            LastAccessDate:      time.Date(2025, time.January, 4, 0, 0, 0, 0, time.UTC),
            MFAEnabled:          true,
        },
        {
            Name:                "Foo1 Bar1",
            CreateDate:          time.Date(2019, time.September, 20, 0, 0, 0, 0, time.UTC),
            PasswordChangedDate: time.Date(2019, time.September, 22, 0, 0, 0, 0, time.UTC),
            LastAccessDate:      time.Date(2025, time.February, 8, 0, 0, 0, 0, time.UTC),
            MFAEnabled:          false,
        },
        {
            Name:                "Foo2 Bar2",
            CreateDate:          time.Date(2022, time.February, 3, 0, 0, 0, 0, time.UTC),
            PasswordChangedDate: time.Date(2022, time.February, 3, 0, 0, 0, 0, time.UTC),
            LastAccessDate:      time.Date(2025, time.February, 12, 0, 0, 0, 0, time.UTC),
            MFAEnabled:          false,
        },
        {
            Name:                "Foo3 Bar3",
            CreateDate:          time.Date(2023, time.March, 7, 0, 0, 0, 0, time.UTC),
            PasswordChangedDate: time.Date(2023, time.March, 10, 0, 0, 0, 0, time.UTC),
            LastAccessDate:      time.Date(2022, time.January, 3, 0, 0, 0, 0, time.UTC),
            MFAEnabled:          true,
        },
        {
            Name:                "Foo Bar4",
            CreateDate:          time.Date(2018, time.April, 8, 0, 0, 0, 0, time.UTC),
            PasswordChangedDate: time.Date(2020, time.April, 12, 0, 0, 0, 0, time.UTC),
            LastAccessDate:      time.Date(2022, time.October, 4, 0, 0, 0, 0, time.UTC),
            MFAEnabled:          false,
        },
    }

    resp := make([]UserResponse, len(users))
    for i, u := range users {
        resp[i] = UserResponse{
            Name:                    u.Name,
            CreateDate:              u.CreateDate.Format("Jan 2 2006"),
            PasswordChangedDate:     u.PasswordChangedDate.Format("Jan 2 2006"),
            DaysSincePasswordChange: computeDaysAgo(u.PasswordChangedDate),
            LastAccessDate:          u.LastAccessDate.Format("Jan 2 2006"),
            DaysSinceLastAccess:     computeDaysAgo(u.LastAccessDate),
            MFAEnabled:              u.MFAEnabled,
        }
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/api/users", usersHandler)

    fs := http.FileServer(http.Dir("../frontend/build"))
    http.Handle("/", fs)

    log.Println("Server listening on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
