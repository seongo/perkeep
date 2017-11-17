// Code generated by reactGen. DO NOT EDIT.

package react

// ButtonProps defines the properties for the <button> element
type ButtonProps struct {
	ClassName               string
	DangerouslySetInnerHTML *DangerousInnerHTML
	ID                      string
	Key                     string

	OnChange
	OnClick

	Role  string
	Style *CSS
	Type  string
}

func (b *ButtonProps) assign(v *_ButtonProps) {

	v.ClassName = b.ClassName

	v.DangerouslySetInnerHTML = b.DangerouslySetInnerHTML

	if b.ID != "" {
		v.ID = b.ID
	}

	if b.Key != "" {
		v.Key = b.Key
	}

	if b.OnChange != nil {
		v.o.Set("onChange", b.OnChange.OnChange)
	}

	if b.OnClick != nil {
		v.o.Set("onClick", b.OnClick.OnClick)
	}

	v.Role = b.Role

	// TODO: until we have a resolution on
	// https://github.com/gopherjs/gopherjs/issues/236
	v.Style = b.Style.hack()

	v.Type = b.Type

}