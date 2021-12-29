import { defineConfig } from 'umi';

export default defineConfig({
  nodeModulesTransform: {
    type: 'none',
  },
  qiankun: { slave: {} },
  routes: [
    {
      path: '/',
      component: '@/layouts',
      routes: [{ path: '/', component: '@/pages/index' }],
    },
  ],
  fastRefresh: {},
});
