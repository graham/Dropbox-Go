package dropbox

import (
	"encoding/json"
	"net/url"
)

func Copy(s Session, uri Uri, p *Parameters) (c Contents, err error) {
	params := map[string]string{
		"root":    uri.Root,
		"to_path": url.QueryEscape(uri.Path),
	}

	if p != nil {
		if p.FromPath != "" {
			params["from_path"] = p.FromPath
		}

		if p.Locale != "" {
			params["locale"] = p.Locale
		}

		if p.FromCopyRef != "" {
			params["from_copy_ref"] = p.FromCopyRef
		}
	}

	body, _, err := s.MakeApiRequest("fileops/copy", params, POST)

	if err != nil {
		return
	}

	var fe FileError
	err = json.Unmarshal(body, &fe)

	if fe.ErrorText != "" {
		err = fe
		return
	}

	err = json.Unmarshal(body, &c)

	return
}

func CreateFolder(s Session, uri Uri, p *Parameters) (m Metadata, err error) {
	params := map[string]string{
		"root": uri.Root,
		"path": url.QueryEscape(uri.Path),
	}

	if p != nil && p.Locale != "" {
		params["locale"] = p.Locale
	}

	body, _, err := s.MakeApiRequest("fileops/create_folder", params, POST)

	if err != nil {
		return
	}

	var fe FileError
	err = json.Unmarshal(body, &fe)

	if fe.ErrorText != "" {
		err = fe
		return
	}

	err = json.Unmarshal(body, &m)

	return
}

func Delete(s Session, uri Uri, p *Parameters) (m Metadata, err error) {
	params := map[string]string{
		"root": uri.Root,
		"path": url.QueryEscape(uri.Path),
	}

	if p != nil && p.Locale != "" {
		params["locale"] = p.Locale
	}

	body, _, err := s.MakeApiRequest("fileops/delete", params, POST)

	if err != nil {
		return
	}

	var fe FileError
	err = json.Unmarshal(body, &fe)

	if fe.ErrorText != "" {
		err = fe
		return
	}

	err = json.Unmarshal(body, &m)

	return
}

func Move(s Session, uri Uri, to_path string, p *Parameters) (m Metadata, err error) {
	params := map[string]string{
		"root":      uri.Root,
		"from_path": url.QueryEscape(uri.Path),
		"to_path":   url.QueryEscape(to_path),
	}

	if p != nil && p.Locale != "" {
		params["locale"] = p.Locale
	}

	body, _, err := s.MakeApiRequest("fileops/move", params, POST)

	if err != nil {
		return
	}

	var fe FileError
	err = json.Unmarshal(body, &fe)

	if fe.ErrorText != "" {
		err = fe
		return
	}

	err = json.Unmarshal(body, &m)

	return
}
