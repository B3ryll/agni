const {resolve}            = require("path");
const HtmlWebpackPlugin    = require("html-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const SveltePreprocess     = require("svelte-preprocess");

module.exports = (env) => {
    return {
        mode: "development",
        entry: {
            main: "./src/main.ts"
        },
        devtool: "source-map",

        module: {
            rules: [
                {
                    test: /\.(html|svelte)$/,
                    use: [
                        {
                            loader: 'svelte-loader',
                            options: {
                                compilerOptions: {
                                    dev: true,
                                },

                                preprocess: SveltePreprocess(),
                            }, 
                        },
                    ]
                },
                {
                    test: /\.css$/,
                    use: [
                        MiniCssExtractPlugin.loader,
                        'css-loader',
                    ]

                },
                {
                    test: /\.ts$/,
                    use: [
                        "ts-loader",
                    ],
                    exclude: [
                        /node_modules/,
                        /.*\.spec\.ts$/,
                    ],
                },
                {
                    test: /node_modules\/svelte\/.*\.mjs$/,
                    resolve: {
                        fullySpecified: false,
                    },
                },
            ],
        },
        plugins: [
            new MiniCssExtractPlugin({
                filename: "[name].css",
            }),
            new HtmlWebpackPlugin(),
        ],
        resolve: {
            alias: {
                svelte: resolve('node_modules', 'svelte')
            },
            extensions:     ['.mjs', '.js', ".ts", '.svelte'],
            mainFields:     ['svelte', 'browser', 'module', 'main'],
            conditionNames: ["svelte"],
        },
    };
};
