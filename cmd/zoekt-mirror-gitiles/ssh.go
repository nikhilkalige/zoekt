// Copyright 2016 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	//"bytes"
	//"fmt"
	//"io/ioutil"
	//"log"
	//"net/http"
	"net/url"
	//"regexp"
	"strings"
)


func getSshRepos(u *url.URL, filter func(string) bool) (map[string]*crawlTarget, error) {
	pages := map[string]*crawlTarget{}

	key := strings.TrimPrefix(u.Path, "/")
	key = strings.TrimSuffix(key, ".git")
	web_url := u.String()
	u.Scheme = ""
	pages[key] = &crawlTarget{
		webURL:     web_url,
		webURLType: "ssh",
		cloneURL: strings.TrimPrefix(u.String(), "//"),
	}
	return pages, nil
}

