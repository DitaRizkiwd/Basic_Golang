package basic

type Post struct {
	Title   string
	Content string
	Author  User
}


//construktor = function yang membentuk struct

func NewPost(title, content string, author User) *Post {
	return &Post{title, content, author}
}

func (p *Post) EditContent(newContent string) {
	p.Content = newContent
}
