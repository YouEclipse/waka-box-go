#

<p align="center">
  <img width="400" src="https://user-images.githubusercontent.com/8252317/83985147-9afb2800-a96a-11ea-9841-eec3a1f61d75.png">
  <h3 align="center">waka-box-go</h3>
  <p align="center">📊 Update  pinned gist / profile README  to contain your weekly WakaTime stats. </p>
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

> if only want's to update a markdown,like profile README,skip step 1 and step 2.

1. Create a new public GitHub Gist with name `📊 Weekly development breakdown` (https://gist.github.com/)
1. Create a token with the `gist` scope and copy it. (https://github.com/settings/tokens/new)
1. Create a WakaTime account (https://wakatime.com/signup)
1. In your WakaTime profile settings (https://wakatime.com/settings/profile) ensure `Display coding activity publicly` and `Display languages, editors, operating systems publicly` are checked.
1. In your account settings, copy the existing WakaTime API Key (https://wakatime.com/settings/api-key)
1. For updating a markdown file，add comments to the place where you want to update in the markdown file.
   ```markdown
    <!-- waka-box start -->
    <!-- waka-box end -->
   ```

### 🚀 Project setup

1. Fork this repo
   

2. Edit the [environment variable](https://github.com/YouEclipse/waka-box-go/actions/runs/126970182/workflow#L17-L19) in `.github/workflows/schedule.yml`:

   > For updating github profile README,you can follow [waka-box.yml](https://github.com/YouEclipse/YouEclipse/blob/master/.github/workflows/waka-box.yml) in [YouEclipse](https://github.com/YouEclipse/YouEclipse) to create a Action in your README repo.Remember it's unsafe to use token with **`repo`** scope for updating the repo, waka-box update the profile repo using git command in Github Action instead of using github API.
   
   > DO NOT CHANGE THE WAKATIME_API_KEY or GH_TOKEN VALUES IN THIS FILE, USE THE REPO SECRETS SET BELOW. FAILURE TO DO THIS WILL MAKE YOUR WAKATIME API KEY PUBLIC AND CAN POTENTIALLY EXPOSE SENSITIVE INFORMATION.

   - **UPDATE_OPTION:** Default is `GIST`.For only update a markdown file ,set to`MARKDOWN`,and ignore environment variables with prefix **GIST\_** below.Set to `GIST_AND_MARKDOWN` updates both the gist and the markdown file.
   - **MARKDOWN_FILE:** The filename for the markdown file.

   - **GIST_ID:** The ID portion from your gist url: `https://gist.github.com/YouEclipse/`**`9bc7025496e478f439b9cd43eba989a4`**.

     **the following are optional, thanks [@AarynSmith](https://github.com/AarynSmith) for PR([#11](https://github.com/YouEclipse/waka-box-go/pull/11))**

   - **GIST_BARSTYLE:** Background of the progress bar. Default is "SOLIDLT" other options include "SOLIDMD", "SOLIDDK" for medium and dark backgrounds, "EMPTY" for blank background, and "UNDERSCORE" for a line along the bottom.
   - **GIST_BARLENGTH:** Length of the progress bar. Default is 21. Set to -1 to auto size the bar.
   - **GIST_TIMESTYLE** Abbreviate the time text. Default is "LONG" ( "# hrs # mins" ). "SHORT" updates the text to "#h#m".

3. Go to the repo **Settings > Secrets**
4. Add the following environment variables:
   - **GH_TOKEN:** The GitHub token generated above.
   - **WAKATIME_API_KEY:** The API key for your WakaTime account.

## 🕵️ How it works

- Get stats from WakaTime API
- Update Gist with Github API
- Use Github Actions for updating Gist

## 📄 License

This project is licensed under [Apache-2.0](./LICENSE)
