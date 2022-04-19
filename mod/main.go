package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var (
	__BASEURL__ = "https://jsonplaceholder.typicode.com"
)

func main() {
	// posts, err := getAllPosts("/posts")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Printf("%+v\n", posts)

	// post, err := getPostById("1")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println("\nGet post by id :")
	// fmt.Printf("%+v\n", post)

	var post = Post{
		UserId: 1,
		Id:     1,
		Title:  "foo",
		Body:   "bar",
	}

	err := createNewPost(&post)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func getAllPosts(enpoint string) (*[]Post, error) {
	res, err := http.Get(__BASEURL__ + enpoint)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res.Body)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	sb := string(body)

	var posts []Post
	err = json.Unmarshal([]byte(sb), &posts)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	defer res.Body.Close()

	return &posts, nil
}

func getPostById(id string) (*Post, error) {
	res, err := http.Get(__BASEURL__ + "/posts/" + id)
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	sb := string(body)

	var post Post
	err = json.Unmarshal([]byte(sb), &post)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	defer res.Body.Close()

	return &post, nil
}

func createNewPost(post *Post) error {
	data, err := json.Marshal(post)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", __BASEURL__+"/posts", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	sb := string(body)

	fmt.Println(sb)

	return nil
}
