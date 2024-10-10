package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/careofyou/music-api/helpers"
	"github.com/careofyou/music-api/services"
	"github.com/go-chi/chi/v5"
)


var song services.Song

// GET/songs
func GetAllSongs(w http.ResponseWriter, r *http.Request) {
    all, err := song.GetAllSongs()
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        return
    }
    helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"songs": all})
}

// GET/songs/song/{id}
func GetSongById(w http.ResponseWriter, r* http.Request) {
    id := chi.URLParam(r, "id")
    song, err := song.GetSongById(id)
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        return
    }
    helpers.WriteJSON(w, http.StatusOK, song)
}

// POST/songs/song
func CreateSong(w http.ResponseWriter, r *http.Request) {
    var songData services.Song
    err := json.NewDecoder(r.Body).Decode(&songData)

    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        return 
    }

    songCreated, err := song.CreateSong(songData)
    // CHECK
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
        return
    }
    helpers.WriteJSON(w, http.StatusOK, songCreated)
}

// PUT/songs/song/{id}
func UpdateSong(w http.ResponseWriter, r *http.Request) {
    var songData services.Song
    id := chi.URLParam(r, "id")
    err := json.NewDecoder(r.Body).Decode(&songData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    songUpdated, err := song.UpdateSong(id, songData)
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
    }
    helpers.WriteJSON(w, http.StatusOK, songUpdated)
}

// DELETE/songs/song/{id}
func DeleteSong(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    err := song.DeleteSong(id)
    if err != nil {
        helpers.MessageLogs.ErrorLog.Println(err)
    }
    helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "successfully deleted"})

}
