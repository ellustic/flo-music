package melon

type User struct {
  ID    int32
  Name  string
  Email string
}

func (p *User) Playlists() []Playlist {
  return make([]Playlist, 0)
}

func (p *User) TableName() string {
  return "users"
}

func (p *User) Paths() []string {
  return []string{
    "/me",
    "/me/player",
    "/me/favorites?type=playlist",
  }
}
