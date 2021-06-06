package hcpairing

import (
	"errors"
	"fmt"
	"strings"
)

const (

	// Tags
	BlurredVision = "Blurred Vision"
	Cough         = "Cough"
	ItchySkin     = "Itchy Skin"
	Pregnancy     = "Pregnancy"
	SoreMuscles   = "Sore Muscles"
	Stomachache   = "Stomachache"
	Toothache     = "Toothache"
	Vomit         = "Vomit"
)

var (
	allTags = []string{
		BlurredVision,
		Cough,
		ItchySkin,
		Pregnancy,
		SoreMuscles,
		Stomachache,
		Toothache,
		Vomit,
	}
	tagsPrefixTreeRoot TagsPrefixTreeNode = buildTagsPrefixTree()
)

type TagsPrefixTreeNode interface {
	getLeaf(key string) (TagsPrefixTreeNode, error)
	addLeafIfNotExists(key string)
	setTag(tag string)
	getTag() string
	scanTags() []string
}

type tagsPrefixTreeNode struct {
	tag    string
	leaves map[string]TagsPrefixTreeNode
}

func buildTagsPrefixTree() TagsPrefixTreeNode {

	root := &tagsPrefixTreeNode{
		tag:    "",
		leaves: map[string]TagsPrefixTreeNode{},
	}
	for _, tag := range allTags {
		addTagsToPrefixTree(root, tag)
	}
	fmt.Println(root)
	return root
}

func addTagsToPrefixTree(root TagsPrefixTreeNode, tag string) {

	ptr := root
	for _, item := range strings.ToLower(tag) {
		ptr.addLeafIfNotExists(string(item))
		ptr, _ = ptr.getLeaf(string(item))
	}
	ptr.setTag(tag)
}

func SearchTags(prefix string) []string {
	var err error
	ptr := tagsPrefixTreeRoot
	prefix = strings.ToLower(prefix)
	for _, item := range prefix {
		ptr, err = ptr.getLeaf(string(item))
		if err != nil {
			return []string{}
		}
	}
	return ptr.scanTags()
}

func (n *tagsPrefixTreeNode) scanTags() []string {
	tags := []string{}
	if n.tag != "" {
		tags = append(tags, n.tag)
	}
	for _, value := range n.leaves {
		newTags := value.scanTags()
		tags = append(tags, newTags...)
	}
	return tags
}

func (n *tagsPrefixTreeNode) addLeafIfNotExists(key string) {
	if _, ok := n.leaves[key]; ok {
		return
	}
	n.leaves[key] = &tagsPrefixTreeNode{
		tag:    "",
		leaves: map[string]TagsPrefixTreeNode{},
	}
}

func (n *tagsPrefixTreeNode) getLeaf(key string) (TagsPrefixTreeNode, error) {
	if _, ok := n.leaves[key]; !ok {
		return nil, errors.New("key does not exists")
	}
	return n.leaves[key], nil
}

func (n *tagsPrefixTreeNode) setTag(tag string) {
	n.tag = tag
}

func (n *tagsPrefixTreeNode) getTag() string {
	return n.tag
}
