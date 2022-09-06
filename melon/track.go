package melon

type Track struct {
  ID          int32
  Name        string
  Genres      []string
  Artists     []*Artist
  AlbumID     int32
  Duration    int
  DiscNumber  int
  TrackNumber int
}

func (p *Track) TableName() string {
  return "tracks"
}
