import { defineConfig } from 'umi';

export default defineConfig({
  nodeModulesTransform: {
    type: 'none',
  },
  qiankun: {
    master: {
      // 注册子应用信息
      apps: [
        {
          name: 'sub1', // 唯一 id
          entry: 'http://localhost:9010', // html entry
        },
        {
          name: 'sub2', // 唯一 id
          entry: 'http://localhost:9020', // html entry
        },
      ],
    },
  },
  routes: [
    {
      path: '/',
      component: '@/layouts',
      routes: [
        { path: '/', component: '@/pages/index' },
        { path: '/sub1', microApp: 'sub1' },
        { path: '/sub2', microApp: 'sub2' },
      ],
    },
  ],
  fastRefresh: {},
});
