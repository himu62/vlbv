const path = require('path');

module.exports = {
    entry: './scripts/index.tsx',
    output: {
        path: path.resolve(__dirname, 'public_html'),
        filename: 'script.js'
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                use: 'ts-loader'
            },
            {
                test: /\.scss$/,
                use: [
                    'style-loader',
                    {loader: 'css-loader', options: {importLoaders: 1}},
                    'postcss-loader',
                    'sass-loader'
                ]
            }
        ]
    }
};
