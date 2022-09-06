package melon

type Artist struct {
  ID        int32
  Name      string
  Images    []string
  Genres    []string
  Albums    []*Album
  Related   []*Artist
  DebutDate string
}

func (p *Artist) TopTracks() []Track {
  return make([]Track, 0)
}

func (p *Artist) TableName() string {
  return "artists"
}

func (p *Artist) Paths() []string {
  return []string{
    "/artists/{id}/albums",
    "/artists/{id}/top-tracks",
    "/artists/{id}/related-artists",
  }
}
