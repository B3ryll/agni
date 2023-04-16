const {resolve} = require("path");

module.exports = (env) => {
    return {
        context: resolve(__dirname, "../"),
        output: {
            path: resolve(__dirname, "../dist"),
        },
    };
};
