package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestUsersHandler(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
    w := httptest.NewRecorder()
    usersHandler(w, req)

    res := w.Result()
    defer res.Body.Close()

    if res.StatusCode != http.StatusOK {
        t.Fatalf("expected status 200; got %v", res.StatusCode)
    }

    var list []UserResponse
    if err := json.NewDecoder(res.Body).Decode(&list); err != nil {
        t.Fatalf("failed to decode JSON: %v", err)
    }

    if len(list) != 5 {
        t.Fatalf("expected 5 users; got %d", len(list))
    }

    first := list[0]
    if first.Name != "Foo Bar1" {
        t.Errorf("expected first Name Foo Bar1; got %q", first.Name)
    }
    if first.CreateDate != "Oct 1 2020" {
        t.Errorf("expected CreateDate Oct 1 2020; got %q", first.CreateDate)
    }
    if first.DaysSincePasswordChange <= 0 {
        t.Errorf("expected positive DaysSincePasswordChange; got %d", first.DaysSincePasswordChange)
    }
    if !first.MFAEnabled {
        t.Errorf("expected MFAEnabled true; got false")
    }

    last := list[4]
    if last.Name != "Foo Bar4" {
        t.Errorf("expected last Name Foo Bar4; got %q", last.Name)
    }
    if last.LastAccessDate != "Oct 4 2022" {
        t.Errorf("expected LastAccessDate Oct 4 2022; got %q", last.LastAccessDate)
    }
}
