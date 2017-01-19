package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// repos
// https://api.github.com/users/bolcom/repos

// Tags
// https://api.github.com/repos/bolcom/kibana/tags

// START OMIT
type Repository struct {
	Name    string `json:"name"`
	TagsUrl string `json:"tags_url"`
}

func getRepos(user string) (list []Repository, err error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", user))
	if err != nil {
		return
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	json.NewDecoder(resp.Body).Decode(&list)
	return
}

// END OMIT

// START OMIT2

type Tag struct {
	Name string `json:"name"`
}

func getTagsForRepo(url string) (list []Tag, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	json.NewDecoder(resp.Body).Decode(&list)
	return
}

// END OMIT2

// START OMIT3

func main() {
	repos, err := getRepos("coreos")
	if err != nil {
		log.Fatalln(err)
	}
	if len(repos) == 0 {
		log.Println("no repos for this user (or did you hit the rate limitter?)")
	}
	for _, each := range repos {
		if list, err := getTagsForRepo(each.TagsUrl); err != nil {
			log.Fatalln(each, err)
		} else {
			if len(list) > 0 {
				fmt.Println(each.Name)
				fmt.Println(list)
			}
		}
	}
}

// END OMIT3
