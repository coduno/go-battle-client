package bc

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/coduno/go-battle-client/model"
)

var hc = http.Client{}

type BattleHelper struct {
	Token     string
	ServerURL string
}

func (bh BattleHelper) Join(nick, t string) interface{} {
	b, err := json.Marshal(struct {
		Type,
		Nick string
	}{
		t,
		nick,
	})
	if err != nil {
		return err
	}
	if err := bh.post("/join", bytes.NewBuffer(b), nil); err != nil {
		return err
	}
	return nil
}

func (bh BattleHelper) Move(d string) (*model.GameObject, interface{}) {
	b, err := json.Marshal(struct {
		Direction string
	}{
		d,
	})
	if err != nil {
		return nil, err
	}
	var me model.GameObject
	if err := bh.post("/move", bytes.NewBuffer(b), &me); err != nil {
		return nil, err
	}
	return &me, nil
}

func (bh BattleHelper) Attack(d string) interface{} {
	b, err := json.Marshal(struct {
		Direction string
	}{
		d,
	})
	if err != nil {
		return err
	}
	if err := bh.post("/attack", bytes.NewBuffer(b), nil); err != nil {
		return err
	}
	return nil
}

func (bh BattleHelper) Map() (model.GameMap, interface{}) {
	var m model.GameMap
	if err := bh.get("/map", &m); err != nil {
		return nil, err
	}
	return m, nil
}

func (bh BattleHelper) Me() (*model.GameObject, interface{}) {
	var me model.GameObject
	if err := bh.get("/me", &me); err != nil {
		return nil, err
	}
	return &me, nil
}

func (bh BattleHelper) get(path string, dst interface{}) interface{} {
	req, err := http.NewRequest("GET", bh.ServerURL+path, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Token "+bh.Token)
	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	if terr := getError(resp); terr != nil {
		return terr
	}
	err = json.NewDecoder(resp.Body).Decode(dst)
	if err != nil {
		return err
	}
	return nil
}

func (bh BattleHelper) post(path string, body *bytes.Buffer, dst interface{}) interface{} {
	req, err := http.NewRequest("POST", bh.ServerURL+path, body)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Token "+bh.Token)
	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	if terr := getError(resp); terr != nil {
		return terr
	}
	if dst != nil {
		err = json.NewDecoder(resp.Body).Decode(dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func getError(resp *http.Response) interface{} {
	if resp.StatusCode == 200 {
		return nil
	}
	var terr model.TypedBattleError
	if err := json.NewDecoder(resp.Body).Decode(&terr); err != nil {
		// TODO add more info
		return model.TypedBattleError{Type: "ClientError"}
	}
	return terr
}
