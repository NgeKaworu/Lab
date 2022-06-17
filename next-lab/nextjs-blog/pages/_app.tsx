import "../styles/globals.less";
import type { AppProps } from "next/app";
import { useEffect } from "react";
import { ConfigProvider } from "antd";

function MyApp({ Component, pageProps }: AppProps) {
  useEffect(() => {
    ConfigProvider.config({
      theme: {
        primaryColor: "green",
      },
    });
  }, []);
  return (
    <ConfigProvider>
      <Component {...pageProps} />
    </ConfigProvider>
  );
}

export default MyApp;
