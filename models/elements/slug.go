package elements

type SlugElement struct {
	Slugs   map[string]string `bson:"controller_values"`
	Element `bson:"_,inline"`
}

func NewSlugElement() SlugElement {
	e := NewElement()
	cv := make(map[string]string)
	se := SlugElement{Element: e, Slugs: cv}
	return se
}

func LoadSlugElement(i string, w *wrapper.Wrapper) (SlugElement, error) {
	e := NewSlugElement()
	err := GetById(i, e, w)
	return e, err
}
