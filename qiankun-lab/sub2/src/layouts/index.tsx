import React, { PropsWithChildren } from 'react';
import { QueryClient, QueryClientProvider } from 'react-query';

import { ConfigProvider } from 'antd';

import styled from 'styled-components';

import zhCN from 'antd/es/locale/zh_CN';

const queyClient = new QueryClient();

const Wrap = styled.div`
  background-color: blue;
  padding: 24px;
`;

export default (props: PropsWithChildren<any>) => {
  return (
    <QueryClientProvider client={queyClient}>
      <ConfigProvider locale={zhCN}>
        <Wrap>{props.children}</Wrap>
      </ConfigProvider>
    </QueryClientProvider>
  );
};
