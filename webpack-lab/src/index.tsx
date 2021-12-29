import React from "react";
import { render } from "react-dom";
import graphql from "./test.graphql";

interface Iprops {
  children: string;
}
function App(props: Iprops) {
  console.log(graphql);
  return (
    <>
      <h1>{props.children}</h1>
    </>
  );
}
render(<App>hello TypeScript! 123 </App>, document.getElementById("root"));
