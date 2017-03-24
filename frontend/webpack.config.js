const path = require('path');

module.exports = {
    entry: "./src/js/entry.js",
    output: {
        path: path.resolve(__dirname, 'dist'),
        filename: "bundle.js"
    },
    module: {
        rules: [
            {
                test: /\.css$/, use: [
                {loader: "style-loader"},
                {loader: "css-loader"}
            ]
            }]
    },
    proxy: {
        "/api": "http://localhost:8080"
    },
    devServer: {
        inline: true,
        port: 8000
    }
};
