package ktrans

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

func (this *Ktrans) ToJSON() (error,[]byte){
	b, err := json.Marshal(this)
	if err != nil {
		return err, []byte{}
	}
	return nil, b
}

func FromJSON(json_str []byte) (error, *Ktrans) {
	var ktrans Ktrans
	err := json.Unmarshal(json_str, &ktrans)
	if err != nil {
		return err, nil
	}
	return nil, &ktrans
}

func (this *Ktrans) ToYAML() (error,[]byte){
	y, err := yaml.Marshal(this)
	if err != nil {
		return err, []byte{}
	}
	return nil, y
}

func FromYAML(yml_str []byte) (error,*Ktrans) {
	var ktrans Ktrans
	err := yaml.Unmarshal(yml_str, &ktrans)
	if err != nil {
		return err, nil
	}
	return nil, &ktrans
}
