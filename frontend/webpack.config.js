const path = require('path');

module.exports = {
    entry: "./src/js/entry.js",
    output: {
        path: path.resolve(__dirname, 'build'),
        filename: "bundle.js"
    },
    module: {
        rules: [
            {
                test: /\.css$/,
                exclude: "/node_modules",
                use: [
                    {loader: "style-loader"},
                    {loader: "css-loader"}
                ]
            }]
    },
    devServer: {
        inline: true,
        port: 8000,
        proxy: {
            "/api": "http://localhost:8080"
        }
    }
};
