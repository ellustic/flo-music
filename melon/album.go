package melon

type Album struct {
  ID        int32
  Name      string
  Kind      string
  Images    []string
  Genres    []string
  Tracks    []*Track
  Artists   []*Artist
  Copyright string
  IssueDate string
}

func (p *Album) TableName() string {
  return "albums"
}

func (p *Album) Paths() []string {
  return []string{
    "/albums/{id}",
    "/albums/{id}/tracks",
  }
}
