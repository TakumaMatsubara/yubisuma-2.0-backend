package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HandleAnimalsGet handles the GET request for all animals
func HandleAnimalsGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		animals, err := GetAnimals()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to fetch animals",
			})
		}

		if len(animals) == 0 {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "No animals found",
			})
		}

		// Convert database animals to response format
		resp := animalsGetResponse{}
		for _, animal := range animals {
			resp.Animals = append(resp.Animals, animalResponse{
				ID:            animal.ID,
				Name:          animal.Name,
				NumFinger:     animal.NumFinger,
				Skill:         animal.Skill,
				SkillDesc:     animal.SkillDesc,
				HandUrl:       animal.HandUrl,
				UpFingerUrl:   animal.UpFingerUrl,
				DownFingerUrl: animal.DownFingerUrl,
			})
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func HandleAnimalPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req AnimalPostRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
		}
		myAnimal := animalSelectedResponse{
			IsPlayer:      true,
			ID:            req.ID,
			Name:          "Lion",
			NumFinger:     5,
			Skill:         "Roar",
			SkillDesc:     "Powerful Roar",
			HandUrl:       "http://example.com/lion/hand",
			UpFingerUrl:   "http://example.com/lion/upfinger",
			DownFingerUrl: "http://example.com/lion/downfinger",
		}
		npcAnimal := animalSelectedResponse{
			IsPlayer:      false,
			ID:            "npc_id_123",
			Name:          "Tiger",
			NumFinger:     5,
			Skill:         "Leap",
			SkillDesc:     "High Leap",
			HandUrl:       "http://example.com/tiger/hand",
			UpFingerUrl:   "http://example.com/tiger/upfinger",
			DownFingerUrl: "http://example.com/tiger/downfinger",
		}
		animalPostResp := AnimalPostResponse{
			MyAnimal:  myAnimal,
			NpcAnimal: npcAnimal,
		}
		return c.JSON(http.StatusOK, animalPostResp)
	}
}

func HandleAction() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req actionRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request format"})
		}
		// resp := actionResponse{IsFinished: false, // 実際のロジックに基づいて更新
		// 	IsMyTurn:          req.IsMyTurn,          // ターンを切り替える
		// 	NumMyFinger:       req.NumMyUpFinger,     // 仮の値
		// 	NumNpcFinger:      5,                     // 仮の値
		// 	NumMyUpFinger:     req.NumMyUpFinger,     // リクエストからの値
		// 	NumNpcUpFinger:    3,                     // 仮の値
		// 	NumExpectedFinger: req.NumExpectedFinger, // リクエストからの値
		// }

		// 必要なパラメータをリクエストボディから抽出する
		// numExpectedFinger, ok1 := requestBody["numExpectedFinger"].(int)
		// useSkill, ok2 := requestBody["useSkill"].(bool)
		// isMyTurn, ok3 := requestBody["isMyTurn"].(bool)
		// numUpMyFinger, ok4 := requestBody["numUpMyFinger"].(int)
		// if !ok1 || !ok2 || !ok3 || !ok4 {
		// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing or invalid parameters"})
		// }
		numExpectedFinger := req.NumExpectedFinger
		useSkill := req.UseSkill
		isMyTurn := req.IsMyTurn
		numUpMyFinger := req.NumMyUpFinger

		// logic.goのUpdateGameState関数を呼び出す
		result, err := UpdateGameState(*numExpectedFinger, useSkill, isMyTurn, numUpMyFinger)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// レスポンスをJSONとして返す
		// return c.JSON(http.StatusOK, resp)
		return c.JSON(http.StatusOK, result)
	}
}
