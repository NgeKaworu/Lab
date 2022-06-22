import type { NextPage } from "next";
import Head from "next/head";
import Image from "next/image";

import { Button, ConfigProvider, Form, Input } from "antd";
import styles from "../styles/Home.module.less";
import "antd/dist/antd.variable.min.css";
import axios from "axios";

const Home: NextPage = () => {
  async function onFinish(data: any) {
    const res = await axios.post("api/login", { data });
    console.log("login client side", res);
  }

  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Form onFinish={onFinish}>
        <Form.Item name="acc">
          <Input />
        </Form.Item>
        <Form.Item name="pwd">
          <Input />
        </Form.Item>
        <Form.Item>
          <Button type="primary" htmlType="submit">
            Login
          </Button>
        </Form.Item>
      </Form>

      <footer className={styles.footer}>
        <a
          href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by{" "}
          <span className={styles.logo}>
            <Image src="/vercel.svg" alt="Vercel Logo" width={72} height={16} />
          </span>
        </a>
      </footer>
    </div>
  );
};

export default Home;
