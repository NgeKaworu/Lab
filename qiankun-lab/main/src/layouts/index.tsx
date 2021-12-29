import React, { PropsWithChildren } from 'react';
import { QueryClient, QueryClientProvider } from 'react-query';

import { ConfigProvider, Space } from 'antd';

import { useHistory } from 'react-router';

import styled from 'styled-components';

import zhCN from 'antd/es/locale/zh_CN';

const queyClient = new QueryClient();

const Wrap = styled.div`
  background-color: red;
  padding: 24px;
`;

export default (props: PropsWithChildren<any>) => {
  const history = useHistory();
  return (
    <QueryClientProvider client={queyClient}>
      <ConfigProvider locale={zhCN}>
        <Space>
          <a
            onClick={() => {
              history.push('/sub1');
            }}
          >
            sub1
          </a>
          <a
            onClick={() => {
              history.push('/sub2');
            }}
          >
            sub2
          </a>
        </Space>
        <Wrap>{props.children}</Wrap>
      </ConfigProvider>
    </QueryClientProvider>
  );
};
