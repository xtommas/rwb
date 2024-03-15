package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/xtommas/rwb/api/DTO"
)

const endpoint string = "https://wallhaven.cc/api/v1/search?q=anime%20girl&sorting=random&ratios=16x9&atleast=1920x1080"

func GetRandomWall() (string, error) {
	// request to get the wallpaper link
	res, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var response dto.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	wallURL := response.Wallpaper[0].Path

	return wallURL, nil

	// file, err := http.Get(wallURL)
	// if err != nil {
	// 	return nil, err
	// }
	// defer file.Body.Close()
	//
	// if res.StatusCode != 200 {
	// 	return nil, err
	// }
	//
	// img, _, err := image.Decode(file.Body)
	// if err != nil {
	// 	return nil, err
	// }

}
