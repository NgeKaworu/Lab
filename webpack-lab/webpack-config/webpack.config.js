const HtmlWebpackPlugin = require("html-webpack-plugin");
const path = require('path')

module.exports = {
  entry: "./src/index.tsx",
  output: {
    filename: "[name][chunkhash:8].js"
  },
  resolve: {
    extensions: [".js", ".jsx", ".ts", ".tsx"]
  },
  module: {
    rules: [
      {
        test: /\.graphql?$/i,
        loader: path.resolve('loader/loader-lab.js'),
      },
      {
        test: /tsx?$/i,
        loader: "ts-loader",
        exclude: /node_modules/
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: "./public/index.html"
    })
  ]
}