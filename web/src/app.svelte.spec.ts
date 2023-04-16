/**
 * @jest-environment jsdom
 **/

import "@testing-library/jest-dom/extend-expect"

import {
    screen,
    render,
    fireEvent,
} from "@testing-library/svelte"

import App from "./app.svelte"

describe("app.svelte", () => {
    test("render properly", () => {
        render(App, {})

        const greet = screen.queryByText("hello, world!")
        expect(greet).toBeInTheDocument()
    })
})
