module.exports = {
    transform: {
            "^.+\\.svelte$": ["svelte-jester", {preprocess: true}],
            "^.+\\.js$":     "babel-jest",
            "^.+\\.ts$":     "ts-jest",
          },
    moduleFileExtensions: [
        "svelte", "js", "ts",
    ],
    testPathIgnorePatterns: ["node_modules"],
    bail: false,
    verbose: true,
    transformIgnorePatterns: ["node_modules"],
    setupFilesAfterEnv: [
        "@testing-library/jest-dom/extend-expect",
    ],
};
