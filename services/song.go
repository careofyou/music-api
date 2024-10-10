package services

import (
	"context"
	"time"
)

type Song struct {
    ID string `json:"id"`
    Name string `json:"name"`
    Group string `json:"group"`
    ReleaseDate string `json:"release_date"`
    Text string `json:"text"`
    Link string `json:"link"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`

}

func (s *Song) GetAllSongs() ([]*Song, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()

    query := `SELECT id, name, "group", "releaseDate", text, link, created_at, updated_at FROM songs`
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return nil, err 
    }

    var songs []*Song
    for rows.Next() {
        var song Song
        err := rows.Scan(
            &song.ID,
            &song.Name,
            &song.Group,
            &song.ReleaseDate,
            &song.Text,
            &song.Link,
            &song.CreatedAt,
            &song.UpdatedAt,
            )
            
            if err != nil {
                return nil, err
            }
            
            songs = append(songs, &song)
    }

    return songs, nil
}

func (s *Song) GetSongById(id string) (*Song, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()
    query := `
        SELECT id, name, "group", "releaseDate", text, link, created_at, updated_at 
        FROM songs
        WHERE id = $1
    `
    var song Song
    
    row := db.QueryRowContext(ctx, query, id)
    err := row.Scan(
        &song.ID,
        &song.Name,
        &song.Group,
        &song.ReleaseDate,
        &song.Text,
        &song.Link,
        &song.CreatedAt,
        &song.UpdatedAt,
    )

    if err != nil {
        return nil, err
    }
    return &song, nil
}

func (s *Song) CreateSong(song Song) (*Song, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()

    song.CreatedAt = time.Now()
    song.UpdatedAt = time.Now()

    query := `
    INSERT INTO songs (name, "group", "releaseDate", text, link, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7) returning *
    `

    _, err := db.ExecContext(
        ctx,
        query,
        song.Name,
        song.Group,
        song.ReleaseDate,
        song.Text,
        song.Link,
        song.CreatedAt,
        song.UpdatedAt,
        )
    
        if err != nil {
            return nil, err
        }    
    return &song, nil
}

func (s *Song) UpdateSong(id string, body Song) (*Song, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()

    query := `
        UPDATE songs
        SET 
            name = $1,
            "group" = $2,
            "releaseDate" = $3,
            text = $4,
            link = $5,
            updated_at = $6
        WHERE id = $7
        returning id, name, "group", "releaseDate", text, link, created_at, updated_at
    `

    _, err := db.ExecContext(
        ctx,
        query,
        body.Name,
        body.Group,
        body.ReleaseDate,
        body.Text,
        body.Link,
        time.Now(),
        id,
        )
    if err != nil {
        return nil, err
    }
    return &body, nil
}

func (s *Song) DeleteSong(id string) error {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()

    query := `
        DELETE FROM songs 
        WHERE id = $1`
    _, err := db.ExecContext(ctx, query, id)
    if err != nil {
        return err
    }
    return nil
}
