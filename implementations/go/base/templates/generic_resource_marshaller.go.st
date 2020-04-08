

// MarshalJSON is a Custom marshaller to add the resourceType property, as required by the specification
func (resource *<Name>) MarshalJSON() ([]byte, error) {
	resource.ResourceType = "<Name>"
	// Dereferencing the pointer to avoid infinite recursion.
	return json.Marshal(*resource)
}

