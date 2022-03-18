package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"shoping/tool"
	"time"
)

type GithubToken struct {
	AccessToken string `json:"access_token"`
	Scope []string `json:"scope"`
	TokenType string `json:"token_type"`
}

type GithubUser struct {
	Login             string    `json:"login"`
	ID                int       `json:"id"`
	NodeID            string    `json:"node_id"`
	AvatarURL         string    `json:"avatar_url"`
	GravatarID        string    `json:"gravatar_id"`
	URL               string    `json:"url"`
	HTMLURL           string    `json:"html_url"`
	FollowersURL      string    `json:"followers_url"`
	FollowingURL      string    `json:"following_url"`
	GistsURL          string    `json:"gists_url"`
	StarredURL        string    `json:"starred_url"`
	SubscriptionsURL  string    `json:"subscriptions_url"`
	OrganizationsURL  string    `json:"organizations_url"`
	ReposURL          string    `json:"repos_url"`
	EventsURL         string    `json:"events_url"`
	ReceivedEventsURL string    `json:"received_events_url"`
	Type              string    `json:"type"`
	SiteAdmin         bool      `json:"site_admin"`
	Name              string    `json:"name"`
	Blog              string    `json:"blog"`
	Location          string    `json:"location"`
	Email             *string   `json:"email"`
	Hireable          bool      `json:"hireable"`
	Bio               string    `json:"bio"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Token             string    `json:"-"`
}

// FetchGithubUser 获取github 用户信息
func FetchGithubUser(code string)(*GithubUser,error){
	//通过code，利用post方法和github提供的url(https://github.com/login/oauth/access_token)接口来获得access_token
	client:=http.Client{}
	cfg:=tool.GetCfg()
	params:=fmt.Sprintf(`{"client_id":"%s","client_secret":"%s","code":"%s"}`,cfg.GitHub.ClientId,cfg.GitHub.ClientSecret,code)
	req,err:=http.NewRequest("POST","https://github.com/login/oauth/access_token",bytes.NewBufferString(params))
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	req.Header.Add("Accept","application/json")
	req.Header.Add("Content-type","application/json")
	res,err:=client.Do(req)
	if err!=nil{
		log.Fatal(err.Error())
		return nil, err
	}
	defer res.Body.Close()
	bs,err:=ioutil.ReadAll(res.Body)

	gt:= GithubToken{}
	err=json.Unmarshal(bs,&gt)
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	//得到token struct
	//开始获取用户信息
	//通过token和get方法和github提供的url(https://api.github.com/user)来获取用户信息
	req,err=http.NewRequest("GET","https://api.github.com/user",nil)
	if err!=nil{
		log.Fatal(err.Error())
		return nil, err
	}
	res,err=client.Do(req)
	if err!=nil{
		log.Fatal(err.Error())
		return nil, err

	}
	bs,err=ioutil.ReadAll(res.Body)
	defer req.Body.Close()
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	user:=GithubUser{}
	err=json.Unmarshal(bs,&user)
	if err!=nil{
		log.Fatal(err.Error())
		return nil,err
	}
	if user.Email == nil {
		tEmail := fmt.Sprintf("%d@github.com", user.ID)
		user.Email = &tEmail
	}
	user.Token=gt.AccessToken
	return &user,nil









}