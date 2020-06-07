# 


<p align="center">
  <img width="400" src="https://user-images.githubusercontent.com/4658208/60469862-2e40bf00-9c2c-11e9-87f7-afe164648de4.png">
  <h3 align="center">waka-box-go</h3>
  <p align="center">📊 Update a pinned gist to contain your weekly WakaTime stats. </p>
  <p align="center">  A Golang implementation, see the original version  <a href="https://github.com/matchai/waka-box">waka-box</a>
   <p align="center">
    <a href="https://github.com/YouEclipse/waka-box-go/workflows/Update%20gist%20with%20WakaTime%20stats/badge.svg?branch=master"><img src="https://github.com/YouEclipse/waka-box-go/workflows/Update%20gist%20with%20WakaTime%20stats/badge.svg?branch=master" alt="Update gist with WakaTime stats"></a>
  </p>
</p>


---
English | [简体中文](./README_zh.md)

> 📌✨ For more pinned-gist projects like this one, check out: https://github.com/matchai/awesome-pinned-gists





## 💻 Setup

### 🎒 Prep work
1. Create a new public GitHub Gist (https://gist.github.com/)
1. Create a token with the `gist` scope and copy it. (https://github.com/settings/tokens/new)
1. Create a WakaTime account (https://wakatime.com/signup)
1. In your WakaTime profile settings (https://wakatime.com/settings/profile) ensure `Display coding activity publicly` and `Display languages, editors, operating systems publicly` are checked.
1. In your account settings, copy the existing WakaTime API Key (https://wakatime.com/settings/api-key)

### 🚀 Project setup
1. Fork this repo
1. Edit the [environment variable](https://github.com/YouEclipse/waka-box-go/actions/runs/126970182/workflow#L17-L19) in `.github/workflows/schedule.yml`:

   - **GIST_ID:** The ID portion from your gist url: `https://gist.github.com/YouEclipse/`**`9bc7025496e478f439b9cd43eba989a4`**.

1. Go to the repo **Settings > Secrets**
1. Add the following environment variables:
   - **GH_TOKEN:** The GitHub token generated above.
   - **WAKATIME_API_KEY:** The API key for your WakaTime account. 

## 🕵️ How it works
- Get stats from  WakaTime API 
- Update Gist with Github API 
- Use Github Actions for updating Gist  

## 📄 License
This project is licensed under [Apache-2.0](./LICENSE)