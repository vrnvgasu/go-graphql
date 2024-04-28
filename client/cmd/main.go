package main

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

var client *graphql.Client

type NewNote struct {
	Text   string `json:"text"`
	UserID int    `json:"userId"`
}

type NewUser struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type NoteUser struct {
	id int `json:"id"`
}

type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	User User   `json:"user"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client = graphql.NewClient("http://localhost:8080/query")

	// 1. Создание двух пользователей
	user1 := createUser("user1", 20)
	user2 := createUser("user2", 40)

	// 2. Создание постов для этих пользователей
	createPost("messageUser1", user1.ID)
	createPost("messageUser1", user2.ID)

	// 3. Получение всех постов
	allNotes := getAllNotes()

	// 4. Получение постов по каждому отдельному пользователю
	notesMap := make(map[int]Note)
	notesMap[user1.ID] = *getFirstNote(getNotesByUserId(user1.ID))
	notesMap[user2.ID] = *getFirstNote(getNotesByUserId(user1.ID))

	fmt.Println(allNotes)
	fmt.Println(notesMap)
	fmt.Println(len(allNotes) == len(notesMap))
}

func getFirstNote(notes []Note) *Note {
	if notes == nil || len(notes) == 0 {
		return nil
	}
	return &notes[0]
}

func getNotesByUserId(userId int) []Note {
	req := graphql.NewRequest(fmt.Sprintf(`
		query {
		   noteByUser(userId: %d) {
			id
			text
			user {
			  id
			}
		  }
		}
    `, userId))
	req.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	type data struct {
		Notes []Note `json:"noteByUser"`
	}
	var respData data
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	return respData.Notes
}

func getAllNotes() []Note {
	req := graphql.NewRequest(`
		query {
			notes {
				id
				text
				user {
					id
				}
		  }
		}
    `)
	req.Header.Set("Cache-Control", "no-cache")

	ctx := context.Background()

	type data struct {
		Notes []Note `json:"notes"`
	}
	var respData data
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	return respData.Notes
}

func createPost(text string, userID int) Note {
	req := graphql.NewRequest(`
		mutation ($input: NewNote!) {
		  createNote(input: $input) {
			id
			text
			user {
			  id
			}
		  }
		}
    `)
	req.Var("input", NewNote{
		Text:   text,
		UserID: userID,
	})
	req.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()
	// run it and capture the response

	type data struct {
		CreateNote Note `json:"createNote"`
	}
	var respData data
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	return respData.CreateNote
}

func createUser(name string, age int) User {
	req := graphql.NewRequest(`
		mutation ($input: NewUser!) {
		  createUser(input: $input) {
			id
			name
			age
		  }
		}
    `)
	req.Var("input", NewUser{
		Name: name,
		Age:  age,
	})
	req.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()
	// run it and capture the response

	type data struct {
		CreateUser User `json:"createUser"`
	}
	var respData data
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	return respData.CreateUser
}
