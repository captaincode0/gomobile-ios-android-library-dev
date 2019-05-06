package transformer

type Transformer interface {
	// Gives a map representation of one object
	ToMap() (outputMap map[string]string)
}
