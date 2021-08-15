module.exports = {
  base: '/wiki/',
  title: 'TIL wiki',
  themeConfig: {
    repo: 'pravusid/TIL',
    sidebar: 'auto',
    searchMaxSuggestions: 10,
  },
  markdown: {
    lineNumbers: true,
  },
  plugins: [
    '@vuepress/back-to-top',
    [
      '@vuepress/google-analytics',
      {
        ga: 'UA-105311426-1',
      },
    ],
    '@vuepress/nprogress',
    [
      'vuepress-plugin-clean-urls',
      {
        normalSuffix: '',
        indexSuffix: '/',
        notFoundPath: '/404.html',
      },
    ],
  ],
};
