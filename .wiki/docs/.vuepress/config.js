module.exports = {
  base: "/wiki/",
  title: "TIL wiki",
  themeConfig: {
    repo: "pravusid/TIL",
    sidebar: "auto",
    searchMaxSuggestions: 10
  },
  markdown: {
    lineNumbers: true
  },
  plugins: [
    [
      "@vuepress/google-analytics",
      {
        ga: "UA-105311426-1"
      }
    ],
    "@vuepress/nprogress"
  ]
};
