#

<p align="center">
  <img width="400" src="https://user-images.githubusercontent.com/8252317/83985147-9afb2800-a96a-11ea-9841-eec3a1f61d75.png">
  <h3 align="center">waka-box-go</h3>
  <p align="center">📊 将你的 WakaTime 每周统计更新在  pined gist / profile README  </p>
  <p align="center">  Golang 实现，查看原始版本  <a href="https://github.com/matchai/waka-box">waka-box</a>
  <p align="center">
    <a href="https://github.com/YouEclipse/waka-box-go/workflows/Update%20gist%20with%20WakaTime%20stats/badge.svg?branch=master"><img src="https://github.com/YouEclipse/waka-box-go/workflows/Update%20gist%20with%20WakaTime%20stats/badge.svg?branch=master" alt="Update gist with WakaTime stats"></a>
  </p>
</p>

---

[English](./README.md) | 简体中文

> 📌✨ 查看更多像这样的 Pinned Gist 项目,传送门: https://github.com/matchai/awesome-pinned-gists

## 💻 安装

### 🎒 前置工作

> 如果只想更新某个 markdown 文件，比如 profile README,可以跳过 1,2 两步

1. 创建一个公开的 GitHub Gist,文件名为`📊 Weekly development breakdown` (https://gist.github.com/)
1. 创建一个拥有 `gist` 权限的 token 并复制. (https://github.com/settings/tokens/new)
1. 创建一个 WakaTime 账号(如果已经有了可以跳过),配置好编辑器插件使用一段时间，建议 WakaTime 后台有数据了再进入下一步。 (https://wakatime.com/signup)
1. 在 WakaTime 的 profile settings (https://wakatime.com/settings/profile) 确保 `Display coding activity publicly` 和 `Display languages, editors, operating systems publicly` 被勾选了
1. 在你的 WakaTime 的 account settings, 复制 WakaTime API Key (https://wakatime.com/settings/api-key)
1. 如果需要更新到某个 markdown 文件，请在对应文件需要更新的地方添加以下注释

   ```markdown
    <!-- waka-box start -->
    <!-- waka-box end -->
   ```
### 🚀 开始安装

1. Fork 这个仓库


2. 编辑 `.github/workflows/schedule.yml` 中的[环境变量](https://github.com/YouEclipse/waka-box-go/actions/runs/126970182/workflow#L17-L19) :

   > 如果是需要更新 github profile README,可以在 profile README 的仓库中创建 Action，具体配置参考 我的 [YouEclipse](https://github.com/YouEclipse/YouEclipse) 中的 [waka-box.yml](https://github.com/YouEclipse/YouEclipse/blob/master/.github/workflows/waka-box.yml).因为使用 **`repo`** 权限的token 来通过 API 更新仓库，可能会不安全，所以我的示例中使用 git 命令来更新，这样更加安全。

   > 不要修改此文件中的 WAKATIME_API_KEY 和 GH_TOKEN VALUES, 使用下方设置的的 Secret.否则你的 WAKATIME API KEY 会变成公开的，导致泄露一些敏感信息。


   - **UPDATE_OPTION:** 默认是 `GIST`,如果只想更新到某个 markdown 文件，设置为`MARKDOWN`,并可以忽略以下以 **GIST\_** 开头的环境变量，如果想同时更新 gist 和 markdown,设置为`GIST_AND_MARKDOWN`
   - **MARKDOWN_FILE:** 如果是更新到某个 markdown 文件，填写 markdown 文件名(包含相对路径或者绝对路径)
   - **GIST_ID:** ID 是 gist url 的后缀 : `https://gist.github.com/YouEclipse/`**`9bc7025496e478f439b9cd43eba989a4`**.

     **以下为可选参数,感谢[@AarynSmith](https://github.com/AarynSmith)的 PR([#11](https://github.com/YouEclipse/waka-box-go/pull/11))**

   - **GIST_BARSTYLE:** 进度条的背景样式. 默认是 "SOLIDLT"，其他样式包括 "SOLIDMD", "SOLIDDK" (黑色), "EMPTY" (空白) 和 "UNDERSCORE"（下划线）.
   - **GIST_BARLENGTH:** 条形图的长度. 默认 21. 设置为 -1 可以自动适配.
   - **GIST_TIMESTYLE:** 时间文本的样式. 默认是 "LONG" ( "# hrs # mins" ). "SHORT" 则是 "#h#m".

3. 前往 fork 后的仓库的 **Settings > Secrets**
4. 添加以下环境变量:
   - **GH_TOKEN:** 前置工作中生成的 github token.
   - **WAKATIME_API_KEY:** WakaTime 的 API key.

## 🕵️ 工作原理

- 基于 WakaTime API 获取统计数据
- 基于 Github API 获取/更新 Gist
- 使用 Github Actions 定时更新 Gist

## 📄 开源协议

本项目使用 [Apache-2.0](./LICENSE) 协议
