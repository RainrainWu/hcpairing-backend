package hcpairing

import (
	"errors"
	"sort"
	"strings"
)

const (

	// Tags
	Anxiety           = "Anxiety"
	Bloated           = "Bloated"
	BlurredVision     = "Blurred Vision"
	Burn              = "Burn"
	ChestFeelsTight   = "Chest Feels Tight"
	Constipation      = "Constipation"
	Cough             = "Cough"
	Diarrhea          = "Diarrhea"
	Drowsy            = "Drowsy"
	DrummingInTheEars = "Drumming In The Ears"
	Fatigue           = "Fatigue"
	Fever             = "Fever"
	Fracture          = "Fracture"
	Headache          = "Headache"
	Heatstroke        = "Heatstroke"
	Insomnia          = "Insomnia"
	ItchySkin         = "Itchy Skin"
	LossOfHearing     = "Loss Of Hearing"
	LostAppetite      = "Lost Appetite"
	NasalBleeding     = "Nasal Bleeding"
	NasalDischarge    = "Nasal Discharge"
	Pregnancy         = "Pregnancy"
	Rash              = "Rash"
	RunnyNose         = "Runny Nose"
	ShortOfBreath     = "Short Of Breath"
	Sneeze            = "Sneeze"
	SoreEyes          = "Sore Eyes"
	SoreMuscles       = "Sore Muscles"
	SoreThroat        = "Sore Throat"
	Stomachache       = "Stomachache"
	Stuffy            = "Stuffy"
	Toothache         = "Toothache"
	Vomit             = "Vomit"
)

var (
	allTags = []string{
		Anxiety,
		Bloated,
		BlurredVision,
		Burn,
		ChestFeelsTight,
		Constipation,
		Cough,
		Diarrhea,
		Drowsy,
		DrummingInTheEars,
		Fatigue,
		Fever,
		Fracture,
		Headache,
		Heatstroke,
		Insomnia,
		ItchySkin,
		LossOfHearing,
		LostAppetite,
		NasalBleeding,
		NasalDischarge,
		Pregnancy,
		Rash,
		RunnyNose,
		ShortOfBreath,
		Sneeze,
		SoreEyes,
		SoreMuscles,
		SoreThroat,
		Stomachache,
		Stuffy,
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
	tags := ptr.scanTags()
	sort.Strings(tags)
	return tags
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
