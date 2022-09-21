package main

import (
	"fmt"
	"strings"
)

type Post struct {
	Title     string
	Text      string
	published bool
}

func (p Post) Headline() string {
	return fmt.Sprintf("%s - %s", p.Title, p.Text[:50])
}

// Published - Creation d'un pointer receiver (sans grand interet puisqu'on ne fait pas d'écriture ici)
func (p *Post) Published() bool {
	return p.published
}

// Publish - Creation d'un pointer receiver publiant le Post
func (p *Post) Publish() {
	p.published = true
}

// Unpublish - Creation d'un pointer receiver dépubliant le Post
func (p *Post) Unpublish() {
	p.published = false
}

func UpperTitle(p *Post) {
	p.Title = strings.ToUpper(p.Title)
}

func main() {
	p := Post{
		Title: "Go release",
		Text: `Morbi a odio lacinia, mattis massa nec, fermentum eros. Donec eu varius nisl. 
Nam luctus massa tincidunt sodales dictum. Donec egestas erat felis, non mollis massa rhoncus sit amet. 
Nunc sagittis urna nibh, quis cursus elit molestie eu. Quisque vel gravida nulla. Phasellus efficitur, 
diam iaculis dignissim pellentesque, felis mi posuere purus, sit amet vestibulum massa nunc vel erat..`,
	}
	fmt.Println(p.Headline())
	fmt.Printf("Post published: %v\n", p.Published())
	p.Publish()
	fmt.Printf("Post published: %v\n", p.Published()) // true
	p.Unpublish()
	fmt.Printf("Post published: %v\n", p.Published()) // false
	UpperTitle(&p)
	fmt.Println(p.Headline())

	// On peut directement créé un nouveau post et référencé son adresse dans un pointeur avec la syntaxe `ptr = &Post{...}`
	pythonPost := &Post{
		Title: "Python intro",
		Text: `Donec egestas erat felis, non mollis massa rhoncus sit amet. 
Nunc sagittis urna nibh, quis cursus elit molestie eu. Quisque vel gravida nulla. Phasellus efficitur.`,
	}
	fmt.Println(pythonPost.Headline())
	UpperTitle(pythonPost)
	fmt.Println(pythonPost.Headline())
}
