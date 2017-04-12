package optional

import (
	"encoding/json"
	"encoding/xml"
)

func (o Bool) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

func (o *Bool) UnmarshalJSON(data []byte) error {
	var v bool
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfBool(v)
	return nil
}

func (o Bool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

func (o *Bool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v bool
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfBool(v)
	return nil
}
