
package app

import (
	"math/rand"
	"time"
	"encoding/json"
	"os"
)

// ランダム変数生成器の初期化
func init() {
	rand.Seed(time.Now().UnixNano())
}

// JSONファイルからゲームの状態を読み込む
func LoadGameState(filePath string) (GameState, error) {
	var state GameState
	file, err := os.Open(filePath)
	if err != nil {
		return state, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&state)
	if err != nil {
		return state, err
	}
	
	return state, nil
}

// 現在のゲームの状態をJSONファイルに保存する
func SaveGameState(state GameState, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(state)
	if err != nil {
		return err
	}
	
	return nil
}

// ゲームの状態を更新し、結果をJSONファイルに保存する
func UpdateGameState(numExpectedFinger int, useSkill bool, isMyTurn bool, numMyUpFinger int) (map[string]interface{}, error) {
	// local fileから現在のゲームの状態を読み込む
	currentState, err := LoadGameState("game_state.json")
	if err != nil {
		return nil, err
	}
	
	// resultにレスポンスとして返す値を格納する
	result := make(map[string]interface{})
	
	// NPCのターン: 残りの指の数から、ランダムに指の数を決定する。skillを使う場合は、skillの情報をcurrentStateから取得する
	numNpcUpFinger := rand.Intn(currentState.numNpcFinger) + 1
	if mySkill == "さいみんじゅつ" {
		numNpcUpFinger = 1
	}

	// skillを使う場合は、skillの情報をcurrentStateから取得する
	mySkill := ""
	if useSkill && !currentState.usedMySkill {
		mySkill = currentState.mySkill
		if mySkill == "ほうこう" && currentState.numNpcFinger < currentState.numNpcMaxFinger {
			currentState.numNpcFinger += 1
		}
		else if mySkill == "すみをはく" {
			currentState.activeInk = true
		}
		currentState.usedMySkill = true
	}

	npcSkill := "" 
	useNpcSkill := rand.Intn(2) == 0
	if useNpcSkill && !currentState.usedNpcSkill {
		npcSkill = currentState.npcSkill
		if npcSkill == "ほうこう" && currentState.numMyFinger < currentState.numMyMaxFinger {
			currentState.numMyFinger += 1
		}
		else if npcSkill == "さいみんじゅつ" {
			numMyUpFinger = 1
		}
		else if npcSkill == "すみをはく" {
			currentState.activeInk = true
		}
		currentState.usedNpcSkill = true
	}
	
	totalFingers := numMyUpFinger + numNpcUpFinger

	// ゲームが終了したかどうかを判定する
	isFinished := false
	if numExpectedFinger == totalFingers {
		// 指の数を減らし、ゲームが終了したかどうかを判定する
		if currentState.numMyFinger > 0 && isMyTurn {
			currentState.numMyFinger -= 1
		}
		if currentState.numNpcFinger > 0 && !isMyTurn{
			currentState.numNpcFinger -= 1
		}
		if currentState.numMyFinger == 0 || currentState.numNpcFinger == 0 {
			isFinished = true
		}
	}
	
	// local fileに更新されたゲームの状態を保存する
	err = SaveGameState(currentState, "game_state.json")
	if err != nil {
		return nil, err
	}
	
	// resultにレスポンスとして返す値を格納する
	result["isFinished"] = isFinished
	result["isMyTurn"] = isMyTurn
	result["numMyFinger"] = currentState.numMyFinger
	result["numNpcFinger"] = currentState.numNpcFinger
	result["numMyUpFinger"] = numMyUpFinger
	result["numNpcUpFinger"] = numNpcUpFinger
	result["numExpectedFinger"] = numExpectedFinger
	result["usedMySkill"] = currentState.usedMySkill
	result["activeInk"] = currentState.activeInk
	
	return result, nil
}
