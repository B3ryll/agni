//
// webpack main configuration file
//

const {merge} = require("webpack-merge");

const devConfig  = require("./webpack/dev.config.js");
const baseConfig = require("./webpack/base.config.js");
const prodConfig = require("./webpack/production.config.js");

module.exports = (env) => {
    const mode =
        env.mode == "production" ? "production" : "development";

    const modConfig =
        mode == "development" ? devConfig(env) : prodConfig(env);

    return merge(
        baseConfig(env),
        modConfig,
    )
};
