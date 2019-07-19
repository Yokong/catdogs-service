package main

import (
	pb "catdogs-proto"
	"context"
	"fmt"

	"github.com/micro/go-micro"
)

const (
	content = `
### 1. 两数之和
golang
func twoSum(nums []int, target int) []int {
	m := make(map[int] int, len(nums))
	
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return []int{}
}
	`
)

func main() {
	setpost()
}

func setpost() {
	service := micro.NewService(micro.Name("post.client"))
	service.Init()

	post := pb.NewPostService("post", service.Client())
	rsp, err := post.Poster(context.TODO(), &pb.SetPostReq{
		Title:   "Python Test",
		Author:  "Yoko",
		Content: content,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rsp)
}
