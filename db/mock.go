package db

type Node struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var (
	Nodes = []*Node{
		{
			Slug:  "node-1",
			Title: "This title of node 1",
			Body:  "1.It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout.",
		},
		{
			Slug:  "node-2",
			Title: "This title of node 2",
			Body:  "2.It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout.",
		}, {
			Slug:  "node-3",
			Title: "This title of node 3",
			Body:  "3.It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout.",
		}, {
			Slug:  "node-4",
			Title: "This title of node 3",
			Body:  "4.It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout.",
		},
	}
)

func FindBySlug(slug string) *Node {
	for _, node := range Nodes {
		if node.Slug == slug {
			return node
		}
	}
	return nil
}
