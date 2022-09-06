package melon

type Playlist struct {
  ID          int32
  Name        string
  Tags        []string
  Images      []string
  Tracks      []*Track
  Public      bool
  OwnerID     int32
  Duration    int
  Description string
}

func (p *Playlist) TableName() string {
  return "playlists"
}

func (p *Playlist) Paths() []string {
  return []string{
    "/playlists/{id}",
    "/playlists/{id}/tracks",
  }
}
